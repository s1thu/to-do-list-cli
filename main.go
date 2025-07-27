package main

import (
	"fmt"
	"to_do_cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
	}
}
