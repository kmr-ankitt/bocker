package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateFS(){
	fmt.Println("Creating File System...")

	err := os.Mkdir("./rootfs", 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Failed to create rootfs directory: %v\n", err)
	}
	
	cmd := exec.Command("debootstrap", "--variant=minbase", "stable", "./rootfs", "http://deb.debian.org/debian/")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to run debootstrap: %v\n", err)
	}
}