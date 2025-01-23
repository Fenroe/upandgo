package main

import (
	"fmt"
	"runtime"

	"github.com/Fenroe/upandgo/internal/commands"
	"github.com/Fenroe/upandgo/internal/config"
)

func main() {
	fmt.Println("Hi banana!")
	fmt.Println(runtime.GOOS)
	fmt.Println("Attempting to update shell...")
	appConfig := config.New()
	commands.CommandUpdate(appConfig)
}
