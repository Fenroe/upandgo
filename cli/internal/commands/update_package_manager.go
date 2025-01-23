package commands

import (
	"os/exec"

	"github.com/Fenroe/upandgo/internal/util"
)

func updatePackageManager(packageManager string) error {
	updateCMD := exec.Command(packageManager, "update")
	upgradeCMD := exec.Command(packageManager, "upgrade", "-y")
	util.PipeCMDOutToOS(updateCMD, upgradeCMD)
	util.RunCommands(updateCMD, upgradeCMD)
	return nil
}
