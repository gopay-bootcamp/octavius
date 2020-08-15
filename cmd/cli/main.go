package main

import (
	"octavius/internal/cli/command"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/printer"
)

func main() {
	printer.InitPrinter()
	daemon.StartClient()
	command.Execute()
}
