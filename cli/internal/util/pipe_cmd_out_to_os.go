package util

import (
	"os"
	"os/exec"
)

func PipeCMDOutToOS(cmds ...*exec.Cmd) {
	for _, cmd := range cmds {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
}
