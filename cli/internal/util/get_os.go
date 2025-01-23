package util

import "runtime"

func getOS() string {
	return runtime.GOOS
}
