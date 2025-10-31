package controls

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/kmr-ankitt/bocker/app/helpers"
)

/**
Implements the child process logic
- sets hostname to "bocker"
- executes the command passed as arguments
**/
func Child(osArgs ...string) {
	syscall.Sethostname([]byte("bocker"))
	helpers.Chroot()
	syscall.Mount("proc", "/proc", "proc", 0, "")

	cmd := exec.Command(osArgs[0], osArgs[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	defer syscall.Unmount("/proc", 0)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
