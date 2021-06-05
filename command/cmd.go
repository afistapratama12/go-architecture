package command

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func CMD(path string, command string) {

	fmt.Println("command", command, " ...in progress")

	if runtime.GOOS == "windows" {
		output := exec.Command("cmd", "/C", command)
		output.Dir = path

		err := output.Run()
		if err != nil {
			log.Println(err.Error())
		}

	} else {
		output := exec.Command("bash", "-C", command)
		output.Dir = path

		err := output.Run()
		if err != nil {
			log.Println(err.Error())
		}

	}

	fmt.Println("success running", command)
	fmt.Println("")
}
