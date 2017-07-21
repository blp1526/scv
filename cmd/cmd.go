package cmd

import (
	"flag"
	"fmt"

	"github.com/blp1526/scv/api"
)

const version = "0.0.3"
const expectedArgsSize = 2

func Run() (msg string, err error) {
	var optVersion bool
	flag.BoolVar(&optVersion, "version", false, "print version number")

	flag.Parse()
	if optVersion {
		return fmt.Sprintf("scv version %s", version), nil
	}

	argsSize := len(flag.Args())
	if argsSize != expectedArgsSize {
		return msg, fmt.Errorf("Expected arguments size is %d, but given %d", expectedArgsSize, argsSize)
	}

	body := &api.Body{}
	zoneName := flag.Args()[0]
	serverName := flag.Args()[1]
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
