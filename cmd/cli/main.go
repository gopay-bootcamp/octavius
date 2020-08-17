package main

import (
	"octavius/internal/cli/command"

	"octavius/internal/cli/printer"
	"octavius/internal/logger"

)



func main() {
	printer.InitPrinter()
	logger.Setup()
	command.Execute()
}