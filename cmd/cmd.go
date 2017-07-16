package cmd

import (
	"fmt"
	"os/exec"

	"github.com/blp1526/scv/api"
	"github.com/blp1526/scv/logger"
)

const version = "0.0.1"
const expectedArgsSize = 2

func Run(args ...string) (msg string, err error) {
	argsSize := len(args) - 1

	if argsSize == 0 {
		msg = fmt.Sprintf("scv version %s", version)
		return msg, nil
	} else if argsSize != expectedArgsSize {
		return msg, fmt.Errorf("Expected arguments size is %d", expectedArgsSize)
	}

	body := &api.Body{}
	zoneName := args[1]
	serverName := args[2]
	err = api.Request(body, zoneName, serverName)
	if err != nil {
		return msg, err
	}

	vncPath := vncPath(body)
	logger.Debug(fmt.Sprintf("VNC Path: %s", vncPath))
	command := "open"
	msg = fmt.Sprintf("%s %s", command, vncPath)
	err = exec.Command(command, vncPath).Run()
	return msg, err
}

func vncPath(body *api.Body) string {
	return "vnc://:" + body.Password + "@" + body.Host + ":" + body.Port
}
