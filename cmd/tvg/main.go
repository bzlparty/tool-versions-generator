package main

import (
	"fmt"
	"os"
)

func main() {
	if err := RunApp(); err != nil {
		fmt.Println("App Error:", err.Error())
		os.Exit(2)
	}
}
