package controls

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Run(osArgs ...string) {
	/**
	if no arguments passed then show help menu for run
	**/
	if len(osArgs) < 1 {
		fmt.Println("Run command requires at least one argument")
		return
	}

	/**
	Using /proc/self/exe to re-execute the current binary
	with "child" as the first argument followed by the user command.
	**/
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}