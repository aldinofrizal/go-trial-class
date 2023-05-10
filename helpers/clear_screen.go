package helpers

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("clear") // Please comment this line if you are using windows
	// cmd := exec.Command("cmd", "/c", "cls") // Please uncomment this line if youre using windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
