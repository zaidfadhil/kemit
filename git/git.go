package git

import (
	"os/exec"
	"strings"
)

type GitFileInfo struct {
	File   string `json:"file"`
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
		if file == "" {
			continue
		}

		status := strings.TrimSpace(file[:2])
		filePath := file[2:]

		fileDiff, err := gitFileDiff(filePath)
		if err != nil {
			return nil, err
		}

		files = append(files, GitFileInfo{
			File:   filePath,
			Status: status,
			Diff:   fileDiff,
		})
	}

	return files, nil
}

func gitStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-status", "--cached")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\n"), nil
}

func gitFileDiff(file string) (string, error) {
	cmd := exec.Command("git", "diff", "--staged", "--", file)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
