package prompt

import "fmt"

/**
 * PromptEngine manages the system instructions for the AI.
 */
type PromptEngine struct {
	Language string
	UseEmoji bool
}

/**
 * GetSystemPrompt returns the professional rules for generating commits.
 * @return string
 */
func (e *PromptEngine) GetSystemPrompt() string {
	emojiInstruction := "Do NOT use emojis."
	if e.UseEmoji {
		emojiInstruction = "Use relevant Gitmoji at the start of the commit message."
	}

	return fmt.Sprintf(`You are an expert developer assistant. Your task is to generate a professional Semantic Commit message.
Rules:
1. Format: <type>(<scope>): <subject> (short and concise)
2. Body: Use bullet points starting with "-" if the change is complex.
3. Style: Professional, active voice, present tense.
4. Language: %s.
5. %s
6. Return ONLY the commit message text.`, e.Language, emojiInstruction)
}

/**
 * GetExplainPrompt returns instructions for the --explain flag.
 * @return string
 */
func (e *PromptEngine) GetExplainPrompt(diff string) string {
	return fmt.Sprintf("Explain the following code changes in a clear, technical way for a senior developer:\n\n%s", diff)
}
