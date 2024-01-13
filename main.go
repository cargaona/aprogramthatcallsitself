package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// A program that calls itself.
func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Println(err)
		return
	}
	exPath := filepath.Dir(ex)
	cmds := exec.Command(exPath + "/app")
	if err != nil {
		log.Println(err)
		return
	}
	err = cmds.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
