package main

import (
	"strconv"
	"strings"
)

var supportedAlgos = [4]int{1, 256, 384, 512}

func joinSupportedAlgos() string {
	var result string

	for _, a := range supportedAlgos {
		result = result + ", " + strconv.Itoa(a)
	}

	return strings.TrimPrefix(result, ", ")
}
