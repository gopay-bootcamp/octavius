package main

import (
	"octavius/internal/cli/command"
	"octavius/internal/cli/config"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/fileUtil"
	"octavius/internal/cli/printer"
)

func main() {
	newPrinter := printer.GetPrinter()
	fileUtil := fileUtil.NewFileUtil()
	clientConfigLoader := config.NewLoader()
	octaviusDaemon := daemon.NewClient(clientConfigLoader)
	command.Execute(octaviusDaemon, fileUtil, newPrinter)
}
