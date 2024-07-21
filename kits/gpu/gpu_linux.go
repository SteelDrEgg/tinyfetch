//go:build linux

package gpu

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func Info() GPUInfo {
	var thisGPU GPUInfo

	output, _ := exec.Command("lspci", "-mm").Output()
	out := fmt.Sprintf("%q", output)
	reg := regexp.MustCompile(`"(VGA|Display|3D) compatible controller+(.*?)+[0-9]:`)
	regOut := reg.FindString(out)
	temp1 := strings.Split(regOut, `\" \"`)
	temp2 := temp1[len(temp1)-1:]
	temp2 = strings.Split(temp2[0], `\"`)

	thisGPU.Name = temp2[0]
	return thisGPU
}
