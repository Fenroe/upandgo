package util

func GetPackageManagers() ([]string, error) {
	opsys := getOS()
	linuxManagers := []string{
		"apt",
		"dnf",
		"pacman",
		"yum",
	}
	macOSManagers := []string{
		"brew",
	}
	windowsManagers := []string{
		"choco",
		"winget",
	}
	switch opsys {
	case "linux":
		return linuxManagers, nil
	case "darwin":
		return macOSManagers, nil
	case "windows":
		return windowsManagers, nil
	default:
		return nil, nil
	}
}
