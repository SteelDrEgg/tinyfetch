package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"tinyfetch/kits"
)

var printMaterial []*kits.Tab

func memFmt() {
	memStat, _ := mem.VirtualMemory()
	var total string = kits.ProperUnit(memStat.Total)
	var used string = kits.ProperUnit(memStat.Used)
	var swap string = kits.ProperUnit(memStat.SwapTotal)
	var usage string = kits.Float2string(memStat.UsedPercent) + "%"
	printMaterial = append(
		printMaterial,
		&kits.Tab{Title: "RAM", Content: used + " / " + total + " (" + usage + ")"},
		&kits.Tab{Title: "Swap", Content: swap},
	)
}

func cpuFmt() {
	cpuStat, _ := cpu.Info()
	var name string = cpuStat[0].ModelName
	var cores string = fmt.Sprintf("%v", cpuStat[0].Cores)
	var count string = ""
	if len(cpuStat) > 1 {
		count = fmt.Sprintf("x%v", len(cpuStat))
	}
	printMaterial = append(
		printMaterial,
		&kits.Tab{Title: "CPU", Content: name + " (" + cores + ")" + " " + count},
	)
}

func diskFmt() {
	diskStat, _ := disk.Usage("/")
	total := kits.ProperUnit(diskStat.Total)
	used := kits.ProperUnit(diskStat.Used)
	usage := kits.Float2string(diskStat.UsedPercent)
	printMaterial = append(
		printMaterial,
		&kits.Tab{Title: "Disk", Content: used + " / " + total + " (" + usage + "%)"},
	)
}

func hostFmt() {
	hostStat, _ := host.Info()
	os := hostStat.Platform
	if len(os) < 1 {
		os = "Windows"
	} else if os == "darwin" {
		os = "macOS"
	}
	osVer := hostStat.PlatformVersion

	arch := hostStat.KernelArch
	kernel := hostStat.OS
	kernalVer := hostStat.KernelVersion

	printMaterial = append(
		printMaterial,
		&kits.Tab{Title: "OS", Content: os + " " + osVer + " " + arch},
	)

	printMaterial = append(
		printMaterial,
		&kits.Tab{Title: "Kernel", Content: kernel + " " + kernalVer},
	)
}

func main() {
	hostFmt()
	cpuFmt()
	memFmt()
	diskFmt()

	kits.PrintMaterial(printMaterial)
}
