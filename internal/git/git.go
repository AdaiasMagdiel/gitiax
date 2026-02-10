package git

import (
	"os"
	"os/exec"
)

/**
 * Add stages files to the git index.
 * @param files ...string
 * @return error
 */
func Add(files ...string) error {
	args := append([]string{"add"}, files...)
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

/**
 * GetStagedDiff retrieves the current staged changes.
 * @return string, error
 */
func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

/**
 * GetUnstagedDiff retrieves changes that are not yet staged.
 * @param files ...string
 * @return string, error
 */
func GetUnstagedDiff(files ...string) (string, error) {
	args := append([]string{"diff"}, files...)
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

/**
 * Commit creates a new commit with the given message.
 * @param message string
 * @return error
 */
func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
