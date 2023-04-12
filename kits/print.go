package kits

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var printRed func(a ...interface{})
var printHost func(a ...interface{})
var printTab func(a ...interface{})

func init() {
	printRed = color.New(color.FgRed, color.Bold).PrintFunc()
	printHost = color.New(color.FgGreen, color.Bold).PrintFunc()
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
	temp := *material[1]
	os := strings.Split(temp.Content, " ")[0]
	var logo [18]string
	if os == "macOS" {
		logo = Logos("apple")
	} else if os == "ubuntu" {
		logo = Logos("ubuntu")
	} else if os == "arch" {
		logo = Logos("arch")
	} else if os == "centos" {
		logo = Logos("centos")
	} else {
		logo = Logos("linux")
	}

	fmt.Println("")
	for index, line := range logo {
		if index < len(material) {
			fmt.Print(line)
			material := *material[index]
			if material.Title == "User" {
				printHost(calcSpacing(material.Title, 8) + material.Content + "\n")
			} else {
				printTab(calcSpacing(material.Title, 8) + material.Title + ": ")
				fmt.Println(material.Content)
			}
		} else {
			fmt.Println(line)
		}
	}
	fmt.Println("")
}
