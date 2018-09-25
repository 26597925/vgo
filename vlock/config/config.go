package config

import (
	"os"
	"github.com/spf13/pflag"
)

var (
	vlockRoot  = os.Getenv("VLOCK_ROOT")
)

const (
	defaultConfigFile = "/config/vlock.json"
	defaultPidfile = "/data/vlock.pid"
)

type Config struct {
	Root                 string                    `json:"root,omitempty"`
	ConfigFile			 string					   `json:"configfile,omitempty"`
	PidFile              string                    `json:"pidfile,omitempty"`
}

func New() *Config {
	config := Config{}
	config.Root = vlockRoot
	config.ConfigFile = vlockRoot + defaultConfigFile
	config.PidFile = vlockRoot + defaultPidfile
	return &config
}

func MergeDaemonConfigurations(flagsConfig *Config, flags *pflag.FlagSet, configFile string) (*Config, error) {
	fileConfig, err := getConflictFreeConfiguration(configFile, flags)
	if err != nil {
		return nil, err
	}
	return fileConfig, err
}

func Reload(configFile string, flags *pflag.FlagSet, reload func(*Config)) error {
	newConfig, err := getConflictFreeConfiguration(configFile, flags)
	reload(newConfig)
	return err
}

func getConflictFreeConfiguration(configFile string, flags *pflag.FlagSet) (*Config, error) {
	var config Config
	return &config, nil
}