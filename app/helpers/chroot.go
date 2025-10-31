package helpers

import (
	"fmt"
	"syscall"
)

func Chroot(){
	CreateFS()

	err := syscall.Chroot("./rootfs")
	if err != nil {
		fmt.Printf("Chroot failed: %v\n", err)
	}
	err = syscall.Chdir("/")
	if err != nil {
		fmt.Printf("Chdir failed: %v\n", err)
	}
}