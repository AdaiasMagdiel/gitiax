# Gitiax

Gitiax is a professional command-line interface (CLI) tool designed to automate and enhance the Git workflow using Artificial Intelligence. It analyzes code changes and generates standardized Semantic Commit messages or technical explanations.

## Features

- **Automated Semantic Commits**: Generates commit messages following the `<type>(<scope>): <subject>` format.
- **AI-Powered Explanations**: Provides technical breakdowns of changes for code reviews or documentation purposes.
- **Direct Integration**: Supports adding and committing files in a single command.
- **Pipe Support**: Capable of processing diffs directly from standard input (stdin).
- **Customizable**: Supports custom prompts, multiple languages, and optional emoji integration.

---

## Installation

### Prerequisites

- Go 1.21 or higher.
- Git installed and configured in your system.

### Build from source

To install the binary globally, run:

```bash
go install github.com/adaiasmagdiel/gitiax@latest

```

---

## Configuration

Gitiax requires three environment variables to communicate with an OpenAI-compatible API (e.g., Groq, OpenAI, or LocalLLM).

Add these to your profile configuration file (`.bashrc`, `.zshrc`, or Windows Environment Variables):

```bash
export GITIAX_API_KEY="your_api_key"
export GITIAX_BASE_URL="https://api.openai.com/v1"
export GITIAX_MODEL="gpt-oss-120b"

```

---

## Usage

### Basic Usage

Generate a commit message suggestion for files already in the staging area:

```bash
gitiax

```

### Staging and Committing

Add files and commit them automatically with the AI-generated message:

```bash
gitiax .
# or for specific files
gitiax main.go internal/git/git.go

```

### Technical Explanation

Explain the current changes without creating a commit:

```bash
gitiax --explain

```

### Pipe Support

Process a specific diff output:

```bash
git diff HEAD~1 | gitiax

```

---

## CLI Options

| Flag          | Description                                                       | Default |
| ------------- | ----------------------------------------------------------------- | ------- |
| `--explain`   | Explains changes instead of generating a commit message.          | `false` |
| `--commit`    | Forces a commit even if no arguments were provided.               | `false` |
| `--no-commit` | Prevents automatic committing even when arguments are passed.     | `false` |
| `--lang`      | Specifies the language for the AI response (e.g., "pt-br", "es"). | `"en"`  |
| `--emoji`     | Includes relevant Gitmojis in the commit message.                 | `false` |
| `--prompt`    | Overrides the default system prompt with custom instructions.     | `""`    |

---

## Technical Architecture

The application is structured into three internal packages:

- `internal/ai`: Handles HTTP communication and JSON orchestration with the AI provider.
- `internal/git`: Interfaces with the local Git binary for index manipulation and history.
- `internal/prompt`: Manages the engineering of system instructions and context.

---

## License

Licensed under the **GNU General Public License v3.0 (GPLv3)**.
See the [LICENSE](LICENSE) and [COPYRIGHT](COPYRIGHT) files for details.
