package main

import (
	"octavius/internal/cli/command"
	"octavius/internal/cli/printer"
)

func main() {
	printer.InitPrinter()

	command.Execute()
}
