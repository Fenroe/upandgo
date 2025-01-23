package commands

import (
	"github.com/Fenroe/upandgo/internal/config"
	"github.com/Fenroe/upandgo/internal/util"
)

func CommandUpdate(config *config.Config) {
	managers, err := util.GetPackageManagers()
	if err != nil {
		return
	}
	installedManagers := []string{}
	for _, manager := range managers {
		if util.IsCommandInstalled(manager) {
			installedManagers = append(installedManagers, manager)
		}
	}
	for _, manager := range installedManagers {
		updatePackageManager(manager)
	}
}
