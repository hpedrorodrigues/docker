package util

import (
	"os"
	"os/exec"
)

func Untargz(fileName, targetDir string) error {
	cmd := exec.Command("tar", "-xf", fileName, "-C", targetDir)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	} else {
		return nil
	}
}
