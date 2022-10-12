package main

import (
	_init "github.com/c/websshterminal.io/init"
	"github.com/c/websshterminal.io/router"
)

func main() {
	_init.CmdRun()
	router.RunSshTerminal()
}
