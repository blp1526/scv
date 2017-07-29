package scv

import (
	"flag"
	"fmt"
	"os/user"
	"path/filepath"
)

var version string

const expectedArgsSize = 2

func Run() (result string, err error) {
	var optVersion bool
	flag.BoolVar(&optVersion, "version", false, "print version number")

	var optVerbose bool
	flag.BoolVar(&optVerbose, "verbose", false, "print debug log")

	flag.Parse()
	if optVersion {
		return fmt.Sprintf("scv version %s", version), err
	}
	logger := &Logger{}
	if optVerbose {
		logger.Verbose = true
	}

	argsSize := len(flag.Args())
	if argsSize != expectedArgsSize {
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
