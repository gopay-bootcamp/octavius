package main

import (
	"octavius/internal/cli/command"
	"octavius/internal/cli/printer"
	"octavius/internal/cli/daemon"
)



func main() {


	printer.InitPrinter()
	daemon.StartClient()
	command.Execute()
}