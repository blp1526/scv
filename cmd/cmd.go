package cmd

import (
	"os/exec"

	"github.com/blp1526/scv/api"
)

func Run(zoneName string, serverName string) error {
	body := &api.Body{}
	err := api.Request(body, zoneName, serverName)
	if err != nil {
		return err
	}

	vncPath := vncPath(*body)
	err = exec.Command("open", vncPath).Run()
	return err
}

func vncPath(body api.Body) string {
	return "// vnc://:" + body.Password + "@" + body.Host + ":" + body.Port
}
