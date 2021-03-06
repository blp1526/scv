package scv

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os/user"
	"path/filepath"
	"strings"
)

const Help = `
Usage: scv [options] [zone name] [server name]

Options:
  -h, --help        print help message and exit
  --init            create $HOME/.scv.json if it does not exist
  -v, --verbose     print debug log
  -V, --version     print version and exit
`

var Version string

type CLI struct {
	Logger Logger
}

func (cli *CLI) Run(args []string) (result string, err error) {
	f := flag.NewFlagSet("scv", flag.ContinueOnError)

	var optHelp bool
	var optLongHelp bool
	f.BoolVar(&optHelp, "h", false, "")
	f.BoolVar(&optLongHelp, "help", false, "")

	var optInit bool
	f.BoolVar(&optInit, "init", false, "")

	var optVersion bool
	var optLongVersion bool
	f.BoolVar(&optVersion, "V", false, "")
	f.BoolVar(&optLongVersion, "version", false, "")

	var optVerbose bool
	var optLongVerbose bool
	f.BoolVar(&optVerbose, "v", false, "")
	f.BoolVar(&optLongVerbose, "verbose", false, "")

	err = f.Parse(args)
	if err != nil {
		return result, err
	}

	if optVersion || optLongVersion {
		result, err = cli.Versionf(Version)
		return result, err
	}

	if optInit {
		current, _ := user.Current()
		dir := filepath.Join(current.HomeDir, ".scv.json")
		config := &Config{Servers: []Server{}}
		result, err = config.CreateFile(dir)
		return result, err
	}

	if optVerbose || optLongVerbose {
		cli.Logger.Verbose = true
	}

	optSize := len(f.Args())
	if optHelp || optLongHelp || optSize == 0 {
		result = Help
		return result, err
	}

	valid := cli.ValidateOptSize(optSize)
	if !valid {
		err = fmt.Errorf("Expected arguments size is 2, but given %d", optSize)
		return result, err
	}

	zoneName := f.Args()[0]
	serverName := f.Args()[1]

	current, _ := user.Current()
	dir := filepath.Join(current.HomeDir, ".scv.json")

	config := &Config{}
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

func (cli *CLI) Versionf(version string) (string, error) {
	a := strings.Split(version, "-")
	if len(a) > 2 {
		return fmt.Sprintf("scv version %s, build %s", a[0][1:], a[2]), nil
	}
	return "", fmt.Errorf("version: %s", version)
}
