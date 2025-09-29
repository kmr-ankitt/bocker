package controls

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

/**
Implements the child process logic
- sets hostname to "bocker"
- executes the command passed as arguments
**/
func Child(osArgs ...string) {
	syscall.Sethostname([]byte("bocker"))

	cmd := exec.Command(osArgs[0], osArgs[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}