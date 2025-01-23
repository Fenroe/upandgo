package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("Hi banana!")
	fmt.Println(runtime.GOOS)
	fmt.Println("Attempting to update shell...")
	cmd := exec.Command("sudo", "apt", "update")
	cmd2 := exec.Command("sudo", "apt", "upgrade", "-y")
	cmd.Stdout = os.Stdout
	cmd2.Stdout = os.Stderr
	cmd.Run()
	cmd2.Run()
}
