//go:build windows

package gpu

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func Info() GPUInfo {
	return ps()
}

// Use cmd, inaccurate, not recommended
func cmd() GPUInfo {
	var thisGPU GPUInfo

	// Get Name and VRAM from wmic
	// VRAM read is not accurate, it's 4GB max
	output, err := exec.Command("wmic", "path", "win32_VideoController", "get", "Name", ",", "AdapterRAM").Output()
	if err != nil {
		thisGPU.Name = "N/A"
	}

	gpus := strings.Split(string(output), "\n")[1:]
	for _, each := range gpus {
		temp := strings.Fields(each)
		vram, _ := strconv.Atoi(temp[0])
		thisGPU.Name = strings.Join(temp[1:], " ")
		thisGPU.VRAM = uint64(vram)
		break
	}

	return thisGPU
}

// Use powershell
func ps() GPUInfo {
	var thisGPU GPUInfo

	//Get Names
	output, _ := exec.Command("powershell", "-c", "gwmi win32_VideoController | FL Name").Output()
	re := regexp.MustCompile(`(:\s)(.*)(\s)`)
	match := re.FindString(string(output))
	thisGPU.Name = match[2 : len(match)-2]

	// Get VRAM
	output, _ = exec.Command("powershell", "-c", `Get-ItemProperty -Path "HKLM:\SYSTEM\ControlSet001\Control\Class\{4d36e968-e325-11ce-bfc1-08002be10318}\0*" -Name "HardwareInformation.AdapterString", "HardwareInformation.qwMemorySize" | Select-Object "HardwareInformation.AdapterString", "HardwareInformation.qwMemorySize"`).Output()
	gpus := strings.Split(string(output), "\n")[3:]
	for _, each := range gpus {
		if len(each) <= 1 {
			continue
		}
		temp := strings.Fields(each)
		vram, _ := strconv.Atoi(temp[len(temp)-1])
		name := strings.Join(temp[:len(temp)-1], " ")
		if name == thisGPU.Name {
			thisGPU.VRAM = uint64(vram)
			break
		}
	}
	return thisGPU
}
