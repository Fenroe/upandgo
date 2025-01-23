package util

import "os/exec"

func IsCommandInstalled(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}
