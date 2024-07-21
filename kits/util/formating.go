package util

import (
	"fmt"
)

func ProperUnit(byteNum uint64) (formatted string) {
	byteNum = byteNum
	if byteNum >= 1099511627776 {
		// If the size over 1TiB
		formatted = Float2string(float64(byteNum)/float64(1024*1024*1024*1024), 1) + "TiB"
	} else if byteNum >= 1073741824 {
		// If the size over 1GiB
		formatted = Float2string(float64(byteNum)/float64(1024*1024*1024), 1) + "GiB"
	} else if byteNum >= 1048576 {
		// If the size over 1MiB
		formatted = Float2string(float64(byteNum)/float64(1024*1024), 1) + "MiB"
	} else if byteNum >= 1024 {
		// If the size over 1KiB
		formatted = Float2string(float64(byteNum)/float64(1024), 1) + "KiB"
	} else {
		// If the size is less than 1KiB
		formatted = fmt.Sprintf("%v", byteNum) + "Bytes"
	}
	return
}
