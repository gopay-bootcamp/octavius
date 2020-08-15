package main

import (
	"octavius/internal/cli/command"
	"octavius/internal/cli/config"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/printer"
)

func main() {
	printer.InitPrinter()
	clientConfigLoader := config.NewLoader()
	octaviusDaemon := daemon.NewClient(clientConfigLoader)
	command.Execute(octaviusDaemon)
}
