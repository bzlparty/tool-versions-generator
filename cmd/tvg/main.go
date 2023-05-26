package main

import (
	"fmt"
)

func main() {
	if err := RunApp(); err != nil {
		fmt.Println(err)
	}
}
