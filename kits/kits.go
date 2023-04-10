package kits

import (
	"fmt"
	"strings"
)

func ProperUnit(byteNum uint64) (formatted string) {
	byteNum = byteNum
	if byteNum >= 1099511627776 {
		// If the size over 1TiB
		formatted = Float2string(float64(byteNum)/float64(1024*1024*1024*1024)) + "TiB"
	} else if byteNum >= 1073741824 {
		// If the size over 1GiB
		formatted = Float2string(float64(byteNum)/float64(1024*1024*1024)) + "GiB"
	} else if byteNum >= 1048576 {
		// If the size over 1MiB
		formatted = Float2string(float64(byteNum)/float64(1024*1024)) + "MiB"
	} else if byteNum >= 1024 {
		// If the size over 1KiB
		formatted = Float2string(float64(byteNum)/float64(1024)) + "KiB"
	} else {
		// If the size is less than 1KiB
		formatted = fmt.Sprintf("%v", byteNum) + "Bytes"
	}
	return
}

func Float2string(num float64) string {
	strOut := fmt.Sprintf("%.1f", num)
	if strOut[len(strOut)-1:len(strOut)] == "0" /* && strOut[len(strOut)-2:len(strOut)-1] == "0" */ {
		return strings.Split(strOut, ".")[0]
	}
	return strOut
}
