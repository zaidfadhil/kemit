package git

import (
	"os/exec"
	"strings"
)

type GitFileInfo struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Diff   string `json:"diff"`
}

func DiffFiles() ([]GitFileInfo, error) {
	stagedFiles, err := gitStagedFiles()
	if err != nil {
		return nil, err
	}

	var files []GitFileInfo
	for _, file := range stagedFiles {
		status := strings.TrimSpace(file[:2])
		name := file[3:]

		diff, err := gitFileDiff(name)
		if err != nil {
			return nil, err
		}

		files = append(files, GitFileInfo{
			Name:   name,
			Status: status,
			Diff:   diff,
		})
	}

	return files, nil
}

func gitStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "status", "--short")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(out)), "\n"), nil
}

func gitFileDiff(file string) (string, error) {
	cmd := exec.Command("git", "diff", "--staged", "--", file)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
