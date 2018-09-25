package main

import (
	"github.com/kardianos/service"
)

type Daemon struct {
	
}

func NewDaemon() *Daemon {
	return &Daemon{}
}

func (v *Daemon) Start(s service.Service) error {
	return nil
}

func (v *Daemon) Stop(s service.Service) error {
	return nil
}