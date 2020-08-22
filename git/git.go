package git

import (
	"os/exec"
)

func CloneRepo(webURL string, targetDir string) error {
	cmd := exec.Command("git", "clone", webURL, targetDir)
	return cmd.Run()
}

func InitializeRepo(targetDir string) error {
	cmd := exec.Command("git", "init", targetDir)
	return cmd.Run()
}

func StageAllChanges() error {
	cmd := exec.Command("git", "add", "-A")
	return cmd.Run()
}

func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	return cmd.Run()
}
