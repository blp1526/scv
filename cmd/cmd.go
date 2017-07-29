package cmd

import (
	"flag"
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/blp1526/scv/api"
	"github.com/blp1526/scv/conf"
	"github.com/blp1526/scv/logger"
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
	l := &logger.Logger{}
	if optVerbose {
		l.Verbose = true
	}

	argsSize := len(flag.Args())
	if argsSize != expectedArgsSize {
		return result, fmt.Errorf("Expected arguments size is %d, but given %d",
			expectedArgsSize, argsSize)
	}

	zoneName := flag.Args()[0]
	serverName := flag.Args()[1]

	current, _ := user.Current()
	dir := filepath.Join(current.HomeDir, "scv.json")

	config := conf.Config{}
	err = config.LoadFile(dir)
	if err != nil {
		return result, err
	}

	serverID, err := config.GetServerID(zoneName, serverName)
	if err != nil {
		return result, err
	}
	l.Debug("ServerID is " + serverID)

	vnc := api.Vnc{
		ZoneName:          zoneName,
		ServerID:          serverID,
		AccessToken:       config.AccessToken,
		AccessTokenSecret: config.AccessTokenSecret,
	}
	result, err = vnc.GetServerAddress()
	if err != nil {
		return result, err
	}
	return result, err
}
