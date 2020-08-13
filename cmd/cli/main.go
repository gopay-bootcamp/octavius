package main

import (
	"octavius/internal/cli/printer"
	"octavius/internal/cli/command"
)



func main() {
	printer.InitPrinter()
	command.Execute()
}