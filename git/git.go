package git

import (
	"os/exec"
)

func Diff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
