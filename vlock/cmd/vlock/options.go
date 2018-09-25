package main

import (
    "os"
    "fmt"
    "vlock/config"
    "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type daemonOptions struct {
	configFile   string
    daemonConfig *config.Config
    flags        *pflag.FlagSet
	Debug        bool
    LogLevel     string
}

func newDaemonOptions(config *config.Config) *daemonOptions {
	return &daemonOptions{
		daemonConfig: config,
	}
}

func (o *daemonOptions) InstallFlags(flags *pflag.FlagSet) {
    
}

func setLogLevel(logLevel string) {
	if logLevel != "" {
		lvl, err := logrus.ParseLevel(logLevel)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse logging level: %s\n", logLevel)
			os.Exit(1)
		}
		logrus.SetLevel(lvl)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}