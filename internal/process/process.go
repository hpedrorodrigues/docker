package process

import (
	"os"
	"os/exec"
	"syscall"
)

func Exec(command string, args []string) (*exec.Cmd, error) {
	cmd := exec.Command(command, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID,
	}

	err := cmd.Run()
	return cmd, err
}
