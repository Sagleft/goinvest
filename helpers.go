package goinvest

import (
	"strconv"
	"strings"
)

func getStringsArrayURLEncoded(array []string) string {
	return strings.Join(array, "%2C")
}

func floatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', 4, 32)
}
