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
	//printTab("RAM:\t")
	var total string = kits.ProperUnit(memStat.Total)
	var used string = kits.ProperUnit(memStat.Used)
	var swap string = kits.ProperUnit(memStat.SwapTotal)
	var usage string = kits.Float2string(memStat.UsedPercent) + "%"
	//fmt.Print(used, " / ", total, " (", usage, ")", "\n")
	//printTab("Swap:\t")
	//fmt.Print(swap, "\n")
	/*
		memOut := kits.Tab{
			Title: "RAM",
			Content: used+" / "+total+" ("+usage+")",
		}
		swapOut := kits.Tab{
			Title: "Swap",
			Content: swap,
		}
		printMaterial=append(printMaterial, &memOut, &swapOut)
	*/
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
	//printTab("CPU:\t")
	//fmt.Print(name, " ("+cores+")", " "+count, "\n")
	printMaterial = append(
		printMaterial,
		&kits.Tab{Title: "CPU", Content: name + " (" + cores + ")" + " " + count},
	)
}

func diskFmt() {
	/*partitions, _ := disk.Partitions(true)
	var partiStats []string
	for _, data := range partitions {
		temp, _ := disk.Usage(data.Mountpoint)
		total := kits.ProperUnit(temp.Total)
		used := kits.ProperUnit(temp.Used)
		usage := kits.Float2string()(temp.UsedPercent)
		partiStats = append(partiStats, data.Mountpoint+" "+used+" / "+total+" "+usage+"%")
	}*/
	diskStat, _ := disk.Usage("/")
	total := kits.ProperUnit(diskStat.Total)
	used := kits.ProperUnit(diskStat.Used)
	usage := kits.Float2string(diskStat.UsedPercent)
	//printTab("Disk:\t")
	//fmt.Print(used, " / ", total, " ("+usage+"%)", "\n")
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

	//printTab("OS:\t")
	//fmt.Print(os, " ", osVer, " ", arch, "\n")
	printMaterial = append(
		printMaterial,
		&kits.Tab{Title: "OS", Content: os + " " + osVer + " " + arch},
	)

	//printTab("Kernel:\t")
	//fmt.Print(kernel, " ", kernalVer, "\n")
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

	//fmt.Print(kits.Logos("linux"))
	//for _, dt := range kits.Logos("apple"){
	//	fmt.Println(dt)
	//}
	//var a [] []string
	//a=append(a, []string{"a","b"})
	//a=append(a, []string{"a","2"}, []string{"1","b"})
	//for _, val := range printMaterial {
	//	fmt.Println(*val)
	//}
	kits.PrintMaterial(printMaterial)
}
