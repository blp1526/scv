package cmd

import "os/exec"

func Run(arg string) error {
	argument := "foo"
	err := exec.Command("open", argument).Run()
	return err
}
