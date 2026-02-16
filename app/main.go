package main

import (
	"fmt"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	var command string

	fmt.Print("$ ")
	fmt.Scanf("%v", &command)
	fmt.Print(command, ": command not found")
}
