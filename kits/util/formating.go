package util

import (
	"strconv"
)

func properUnitHelper(bytes uint64, pow uint8, unit string) string {
	quotient := bytes >> pow
	temp := bytes & ((1 << pow) - 1)
	temp = ((temp * 10) + ((1 << pow) >> 1)) >> pow
	if temp == 10 {
		temp = 0
		quotient += 1
	}
	return strconv.FormatUint(quotient, 10) +
		"." + strconv.FormatUint(temp, 10) + " " + unit
}

func ProperUnit(byteNum uint64) (formatted string) {
	if byteNum >= 1<<40 { // TiB
		return properUnitHelper(byteNum, 40, "TiB")
	} else if byteNum >= 1<<30 { // GiB
		return properUnitHelper(byteNum, 30, "GiB")
	} else if byteNum >= 1<<20 { // MiB
		return properUnitHelper(byteNum, 20, "MiB")
	} else if byteNum >= 1<<10 { // KiB
		return properUnitHelper(byteNum, 10, "KiB")
	}
	return strconv.FormatUint(byteNum, 10) + " B"
}
