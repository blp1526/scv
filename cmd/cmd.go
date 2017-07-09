package cmd

import (
	"os/exec"

	"github.com/blp1526/scv/api"
	"github.com/blp1526/scv/logger"
)

func Run(zoneName string, serverName string) error {
	body := &api.Body{}
	err := api.Request(body, zoneName, serverName)
	if err != nil {
		return err
	}

	vncPath := vncPath(*body)
	logger.Debug("vncPath is " + vncPath)
	err = exec.Command("open", vncPath).Run()
	return err
}

func vncPath(body api.Body) string {
	return "vnc://:" + body.Password + "@" + body.Host + ":" + body.Port
}
