package scv

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os/user"
	"path/filepath"
)

var version string

const expectedArgsSize = 2
const help = `
Usage: scv [options] [zone name] [server name]

Options:
  -h, --help        print help message and exit
  -v, --verbose     print debug log
  -V, --version     print version and exit
`

type CLI struct {
	Logger Logger
}

func (cli *CLI) Run() (result string, err error) {
	var optHelp bool
	var optLongHelp bool
	flag.BoolVar(&optHelp, "h", false, "")
	flag.BoolVar(&optLongHelp, "help", false, "")

	var optVersion bool
	var optLongVersion bool
	flag.BoolVar(&optVersion, "V", false, "")
	flag.BoolVar(&optLongVersion, "version", false, "")

	var optVerbose bool
	var optLongVerbose bool
	flag.BoolVar(&optVerbose, "v", false, "")
	flag.BoolVar(&optLongVerbose, "verbose", false, "")

	flag.Parse()

	if optHelp || optLongHelp {
		return fmt.Sprintf("%s", help), err
	}

	if optVersion || optLongVersion {
		return fmt.Sprintf("scv version %s", version), err
	}

	if optVerbose || optLongVerbose {
		cli.Logger.Verbose = true
	}

	argsSize := len(flag.Args())
	if argsSize == 0 {
		return fmt.Sprintf("%s", help), err
	} else if argsSize != expectedArgsSize {
		return result, fmt.Errorf("Expected arguments size is %d, but given %d",
			expectedArgsSize, argsSize)
	}

	zoneName := flag.Args()[0]
	serverName := flag.Args()[1]

	current, _ := user.Current()
	dir := filepath.Join(current.HomeDir, ".scv.json")

	config := Config{}
	err = config.LoadFile(dir)
	if err != nil {
		return result, err
	}

	serverID, err := config.ServerID(zoneName, serverName)
	if err != nil {
		return result, err
	}
	cli.Logger.Debug("ServerID: " + serverID)

	api := &API{}
	url := api.URL(zoneName, serverID)
	body, err := api.GET(url, config.AccessToken, config.AccessTokenSecret)
	if err != nil {
		return result, err
	}
	cli.Logger.Debug("URL: " + url)

	vnc := &VNC{}
	json.NewDecoder(bytes.NewReader(body)).Decode(vnc)
	result = vnc.Path()
	return result, err
}