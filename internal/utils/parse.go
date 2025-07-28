package utils

import (
	"strconv"
	"time"
)

func ParseDate(input string) (time.Time, error) {
	return time.Parse("2006-01-02", input)
}

func ParsePrice(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}
