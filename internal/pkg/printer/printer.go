package printer

import "github.com/fatih/color"

// Println function will print the message in given color
func Println(msg string, attr ...color.Attribute) {
	color.New(attr...).Println(msg)
}
