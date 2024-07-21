//go:build darwin

package gpu

import (
	"fmt"
	"os/exec"
	"regexp"
)

func Info() GPUInfo {
	var thisGPU GPUInfo

	output, err := exec.Command("system_profiler", "SPDisplaysDataType").Output()
	if err != nil {
		thisGPU.Name = "N/A"
	}
	re := regexp.MustCompile(`Graphics/Displays:\n\n(.*?):\n\n`)
	out := re.Find(output)
	output = nil
	name := fmt.Sprintf("%q", out)
	name = gpu[27 : len(gpu)-6]
	thisGPU.Name = name
	return thisGPU
}
