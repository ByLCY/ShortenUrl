package util

import (
	"ByLCY/ShortenUrl/variable"
	"bytes"
	"math"
)

func GenShortPath(num uint64, length int) string {
	chartLength := uint64(len(variable.URIChartSet))
	var buffer bytes.Buffer
	var idx uint64

	for num != 0 {
		idx = num % chartLength
		buffer.WriteString(variable.URIChartSet[idx : idx+1])
		num = uint64(math.Floor(float64(num / chartLength)))
	}

	for i := len(buffer.String()); i < length; i++ {
		buffer.WriteString(variable.URIChartSet[0:0])
	}

	return buffer.String()
}
