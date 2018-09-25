package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func InitSignal() chan os.Signal {
	c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	// In windows
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	return c
}

func HandleSignal(c chan os.Signal) {
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("stop===============")
			return
		case syscall.SIGHUP:
			// TODO reload
			//return
		default:
			return
		}
	}
}

func main() {
	HandleSignal(InitSignal())
}