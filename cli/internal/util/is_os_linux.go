package util

func IsOSLinux() bool {
	return getOS() == "linux"
}
