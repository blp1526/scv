package cmd

import (
	"fmt"

	"github.com/blp1526/scv/api"
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

	msg = vncPath(body)
	return msg, err
}

func vncPath(body *api.Body) string {
	return "vnc://:" + body.Password + "@" + body.Host + ":" + body.Port
}
