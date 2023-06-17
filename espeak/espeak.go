package espeak

import (
	"log"
	"os/exec"
)

func Speak() {
	s := "Make the computer speak"
	cmd := exec.Command("espeak", s)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
