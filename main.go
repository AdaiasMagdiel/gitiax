package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/adaiasmagdiel/gitiax/internal/ai"
	"github.com/adaiasmagdiel/gitiax/internal/git"
	"github.com/adaiasmagdiel/gitiax/internal/prompt"
)

func main() {
	// Flags
	explain := flag.Bool("explain", false, "Explain the changes instead of generating a commit")
	commitFlag := flag.Bool("commit", false, "Force commit even if no files were passed as arguments")
	noCommit := flag.Bool("no-commit", false, "Do not commit even if files were passed as arguments")
	emoji := flag.Bool("emoji", false, "Add emojis to the commit message")
	lang := flag.String("lang", "en", "Language for the message")
	userPrompt := flag.String("prompt", "", "Override default prompt")
	flag.Parse()

	config := ai.Config{
		APIKey:  os.Getenv("GITIAX_API_KEY"),
		BaseURL: os.Getenv("GITIAX_BASE_URL"),
		Model:   os.Getenv("GITIAX_MODEL"),
	}

	if err := validateConfig(config); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Logic to handle Piped input (e.g., git diff | gitiax)
	var pipedDiff string
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		reader := bufio.NewReader(os.Stdin)
		var buf []byte
		for {
			line, err := reader.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			buf = append(buf, line...)
		}
		pipedDiff = string(buf)
	}

	// 1. Handle arguments: "gitiax .", "gitiax main.go", etc.
	args := flag.Args()
	autoCommit := false

	if len(args) > 0 && !*explain {
		fmt.Printf("Adding files to staging: %v\n", args)
		if err := git.Add(args...); err != nil {
			log.Fatalf("Failed to add files: %v", err)
		}

		if !*noCommit {
			autoCommit = true
		}
	}

	// 2. Determine which diff to use
	var diff string
	if pipedDiff != "" {
		diff = pipedDiff
	} else if *explain && len(args) > 0 {
		var err error
		diff, err = git.GetUnstagedDiff(args...)
		if err != nil {
			log.Fatalf("Failed to get diff: %v", err)
		}
	} else {
		var err error
		diff, err = git.GetStagedDiff()
		if err != nil {
			log.Fatalf("Failed to get git diff: %v", err)
		}
	}

	if diff == "" {
		fmt.Println("No changes detected. Stage some files or pass them as arguments.")
		return
	}

	// 3. AI Processing
	engine := &prompt.PromptEngine{Language: *lang, UseEmoji: *emoji}
	client := &ai.Client{Cfg: config}

	var system, user string
	if *explain {
		system = "You are a senior developer explaining code changes."
		user = engine.GetExplainPrompt(diff)
	} else {
		system = engine.GetSystemPrompt()
		if *userPrompt != "" {
			system = *userPrompt
		}
		user = diff
	}

	fmt.Println("Gitiax is thinking...")
	response, err := client.FetchCompletion(system, user)
	if err != nil {
		log.Fatalf("AI Error: %v", err)
	}

	fmt.Println("\n--- Gitiax Suggestion ---")
	fmt.Println(response)

	// 4. Execution Logic
	// We commit if: (autoCommit OR --commit flag is true)
	// AND we are NOT in --explain mode
	// AND we are NOT in --no-commit mode
	if (autoCommit || *commitFlag) && !*explain && !*noCommit {
		if err := git.Commit(response); err != nil {
			log.Fatalf("Failed to commit: %v", err)
		}
		fmt.Println("\nChanges committed successfully!")
	}
}

/**
 * validateConfig ensures all required environment variables are set.
 * @param cfg ai.Config
 * @return error
 */
func validateConfig(cfg ai.Config) error {
	if cfg.APIKey == "" {
		return fmt.Errorf("missing GITIAX_API_KEY environment variable")
	}
	if cfg.BaseURL == "" {
		return fmt.Errorf("missing GITIAX_BASE_URL environment variable")
	}
	if cfg.Model == "" {
		return fmt.Errorf("missing GITIAX_MODEL environment variable")
	}
	return nil
}
