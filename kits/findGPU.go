package kits

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func FindGPU(kernel string) (gpu string) {
	if kernel == "darwin" {
		gpu = osx()
	} else if kernel == "linux" {
		gpu = linux()
	}
	return
}

func osx() string {
	output, _ := exec.Command("system_profiler", "SPDisplaysDataType").Output()
	reg := regexp.MustCompile(`Graphics/Displays:\n\n(.*?):\n\n`)
	out := reg.Find(output)
	output = nil
	gpu := fmt.Sprintf("%q", out)
	gpu = gpu[27 : len(gpu)-6]
	return gpu
}

func linux() string {
	output, _ := exec.Command("lspci", "-mm").Output()
	out := fmt.Sprintf("%q", output)
	reg := regexp.MustCompile(`"(VGA|Display|3D) compatible controller+(.*?)+[0-9]:`)
	regOut := reg.FindString(out)
	temp1 := strings.Split(regOut, `\" \"`)
	temp2 := temp1[len(temp1)-1:]
	temp2 = strings.Split(temp2[0], `\"`)
	return temp2[0]
}
