package cmd

import (
	"fmt"
	"os/exec"

	"github.com/blp1526/scv/api"
)

func Run(zoneName string, serverName string) error {
	apiResponse, err := api.Request(zoneName, serverId)
	vncPath = cmd.vncPath(apiResponse)
	err := exec.Command("open", vncPath).Run()

	if err != nil {
		fmt.Println(err)
	}
}

func vncPath(apiResponse) string {
	password = ""
	host = ""
	port = ""

	return "// vnc://:%s@%s:%s" + password + host + port
}
