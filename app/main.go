package main

import (
	"os"

	"github.com/kmr-ankitt/bocker/app/controls"
)

func main() {
	if len(os.Args) < 2 {
		controls.Help()
		return
	}

	switch os.Args[1] {
	case "run":
		controls.Run(os.Args[2:]...)
	case "child":
		controls.Child(os.Args[2:]...)
	case "help":
		controls.Help()
	}
}
