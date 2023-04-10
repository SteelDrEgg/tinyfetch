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

func calcSpacing(word string, width int) string {
	spaceing := width - len(word)
	space := ""
	for i := 0; i < spaceing; i++ {
		space += " "
	}
	return space
}

func PrintMaterial(material []*Tab) {
	temp := *material[0]
	os := strings.Split(temp.Content, " ")[0]
	var logo [18]string
	if os == "macOS" {
		logo = Logos("apple")
	} else if os == "ubuntu" {
		logo = Logos("ubuntu")
	} else if os == "arch" {
		logo = Logos("arch")
	} else {
		logo = Logos("linux")
	}

	fmt.Println("")
	for index, line := range logo {
		if index < len(material) {
			fmt.Print(line)
			material := *material[index]
			printTab(material.Title + ":" + calcSpacing(material.Title, 8))
			fmt.Println(material.Content)
		} else {
			fmt.Println(line)
		}
	}
	fmt.Println("")
}
