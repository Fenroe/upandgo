package util

import "os/exec"

func RunCommands(cmds ...*exec.Cmd) {
	for _, cmd := range cmds {
		cmd.Run()
	}
}
