package main

import (
	"tinyfetch/kits"
	"tinyfetch/kits/collectinfo"
)

func main() {

	collectinfo.HostFmt()
	collectinfo.CpuFmt()
	collectinfo.MemFmt()
	collectinfo.GpuFmt()
	collectinfo.DiskFmt()

	kits.PrintMaterial(collectinfo.PrintMaterial)
}
