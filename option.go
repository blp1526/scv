package scv

import (
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

func Run() (result string, err error) {
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

	logger := &Logger{}
	if optVerbose || optLongVerbose {
		logger.Verbose = true
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

	serverID, err := config.GetServerID(zoneName, serverName)
	if err != nil {
		return result, err
	}
	logger.Debug("ServerID: " + serverID)

	api := Api{
		ZoneName:          zoneName,
		ServerID:          serverID,
		AccessToken:       config.AccessToken,
		AccessTokenSecret: config.AccessTokenSecret,
		Logger:            *logger,
	}
	result, err = api.GetServerAddress()
	if err != nil {
		return result, err
	}
	return result, err
}
