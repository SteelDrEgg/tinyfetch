package collectinfo

import (
	"fmt"
	"tinyfetch/kits"
	"tinyfetch/kits/gpu"
	"tinyfetch/kits/util"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

var PrintMaterial []*kits.Tab

func MemFmt() {
	memStat, _ := mem.VirtualMemory()
	var total string = util.ProperUnit(memStat.Total)
	var used string = util.ProperUnit(memStat.Used)
	var swap string = util.ProperUnit(memStat.SwapTotal)
	var usage string = util.Float2string(memStat.UsedPercent, 0) + "%"
	PrintMaterial = append(
		PrintMaterial,
		&kits.Tab{Title: "RAM", Content: used + " / " + total + " (" + usage + ")"},
		&kits.Tab{Title: "Swap", Content: swap},
	)
}

func CpuFmt() {
	cpuStat, _ := cpu.Info()
	var name string = cpuStat[0].ModelName
	var threads string = fmt.Sprintf("%v", cpuStat[0].Cores)
	var count string = ""
	if len(cpuStat) > 1 {
		count = fmt.Sprintf("x%v", len(cpuStat))
	}
	PrintMaterial = append(
		PrintMaterial,
		&kits.Tab{Title: "CPU", Content: name + " (" + threads + count + ")"},
	)
}

func DiskFmt() {
	diskStat, _ := disk.Usage("/")
	total := util.ProperUnit(diskStat.Total)
	used := util.ProperUnit(diskStat.Used)
	usage := util.Float2string(diskStat.UsedPercent, 0)
	PrintMaterial = append(
		PrintMaterial,
		&kits.Tab{Title: "Disk", Content: used + " / " + total + " (" + usage + "%)"},
	)
}

func HostFmt() {
	hostStat, _ := host.Info()
	os := hostStat.Platform
	users, _ := host.Users()
	hostName := hostStat.Hostname
	osVer := hostStat.PlatformVersion

	arch := hostStat.KernelArch
	kernel := hostStat.OS
	kernalVer := hostStat.KernelVersion

	if len(users) > 0 {
		PrintMaterial = append(
			PrintMaterial,
			&kits.Tab{Title: "User", Content: users[0].User + "@" + hostName},
		)
	} else {
		PrintMaterial = append(
			PrintMaterial,
			&kits.Tab{Title: "Host", Content: hostName},
		)
	}

	PrintMaterial = append(
		PrintMaterial,
		&kits.Tab{Title: "OS", Content: os + " " + osVer + " " + arch},
	)

	PrintMaterial = append(
		PrintMaterial,
		&kits.Tab{Title: "Kernel", Content: kernel + " " + kernalVer},
	)
}

func GpuFmt() {
	thisGPU := gpu.Info()
	if len(thisGPU.Name) > 0 {
		PrintMaterial = append(
			PrintMaterial,
			&kits.Tab{Title: "GPU", Content: thisGPU.Name},
		)
	}
	if thisGPU.VRAM != 0 {
		PrintMaterial = append(
			PrintMaterial,
			&kits.Tab{Title: "VRAM", Content: util.ProperUnit(thisGPU.VRAM)},
		)
	}
}
