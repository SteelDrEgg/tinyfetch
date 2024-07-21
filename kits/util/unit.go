package util

import "fmt"

func Float2string(num float64, decimal int) string {
	strOut := fmt.Sprintf("%."+fmt.Sprintf("%d", decimal)+"f", num)
	return strOut
}
