package printer

import (

	"github.com/fatih/color"
)

type Printer interface {
	Println(string, ...color.Attribute)
}

var PrinterInstance Printer

type printer struct{}

func InitPrinter() {
	if PrinterInstance == nil {
		PrinterInstance = &printer{}
	}
}

func (p *printer) Println(msg string, attr ...color.Attribute) {
	color.New(attr...).Println(msg)
}
