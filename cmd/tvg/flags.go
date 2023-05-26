package main

import (
	"errors"
	"fmt"
	"strings"
)

func isSupportedAlgo(a int) bool {
	for _, b := range supportedAlgos {
		if b == a {
			return true
		}
	}
	return false
}

func appendPlatformFlag(a *[]string) func(string) error {
	return func(s string) error {
		*a = append(*a, strings.Split(s, ",")...)
		return nil
	}
}

func validateFlags() error {
	if repo == "" {
		return errors.New("No repo given, use --repo | -r")
	}

	if !isSupportedAlgo(algo) {
		return errors.New(fmt.Sprintf("'%d' is not supported, choose from: %s", algo, joinSupportedAlgos()))
	}

	if len(platforms) == 0 {
		return errors.New("No platform given, use --platform | -p")
	}

	return nil
}
