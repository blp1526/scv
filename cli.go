package scv

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os/user"
	"path/filepath"
)

const Help = `
Usage: scv [options] [zone name] [server name]

Options:
  -h, --help        print help message and exit
  -v, --verbose     print debug log
  -V, --version     print version and exit
`

var Version string

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

	if optVersion || optLongVersion {
		result = cli.Versionf(Version)
		return result, err
	}

	if optVerbose || optLongVerbose {
		cli.Logger.Verbose = true
	}

	optSize := len(flag.Args())
	if optHelp || optLongHelp || optSize == 0 {
		result = Help
		return result, err
	}

	valid := cli.ValidateOptSize(optSize)
	if !valid {
		err = fmt.Errorf("Expected arguments size is 2, but given %d", optSize)
		return result, err
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
	cli.Logger.Debug("URL: " + url)
	body, err := api.GET(url, config.AccessToken, config.AccessTokenSecret)
	if err != nil {
		return result, err
	}

	vnc := &VNC{}
	json.NewDecoder(bytes.NewReader(body)).Decode(vnc)
	result = vnc.Path()
	return result, err
}

func (cli *CLI) ValidateOptSize(optSize int) bool {
	return optSize == 2
}

func (cli *CLI) Versionf(version string) string {
	return fmt.Sprintf("scv version %s", version)
}
