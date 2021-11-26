package goinvest

import "strings"

func getStringsArrayURLEncoded(array []string) string {
	return strings.Join(array, "%2C")
}
