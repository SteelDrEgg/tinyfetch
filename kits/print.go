package kits

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var printRed func(a ...interface{})
var printTab func(a ...interface{})

func init() {
	printRed = color.New(color.FgRed).PrintFunc()
	printTab = printRed
}

type Tab struct {
	Title   string
	Content string
}

func PrintMaterial(material []*Tab) {
	temp := *material[0]
	os := strings.Split(temp.Content, " ")[0]
	var logo [18]string
	if os == "macOS" {
		logo = Logos("apple")
	} else if os == "ubuntu" {
		logo = Logos("ubuntu")
	} else {
		logo = Logos("linux")
	}

	fmt.Println("")
	for index, line := range logo {
		if index < len(material) {
			fmt.Print(line)
			material := *material[index]
			printTab(material.Title + ":   \t")
			fmt.Println(material.Content)
		} else {
			fmt.Println(line)
		}
	}
	fmt.Println("")
}
