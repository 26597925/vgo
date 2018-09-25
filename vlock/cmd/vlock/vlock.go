package main

import (
	"flag"
	"fmt"
	"vlock/config"
	"github.com/spf13/pflag"
	"github.com/kardianos/service"
)

const (
	version = "v1.0"
	build = 122
)

func showVersion() {
	fmt.Printf("Vlock version %s, build %d\n", version, build)
}

func main()  {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	svcConfig := &service.Config{
		Name:        "Vlock Master",
		DisplayName: "Vlock is master proccess",
		Description: "This is an Vlock as Daemon.",
	}
	vlock := NewDaemon()
	s, err := service.New(vlock, svcConfig)
	if err != nil {
		return
	}

	s.Run()

	opts := newDaemonOptions(config.New())

	opts.InstallFlags(pflag.CommandLine)

	pflag.Parse()

}
