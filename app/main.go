package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {

		allowed_commands := map[string]func(args []string){
			"exit": func(args []string) { os.Exit(0) },
			"echo": func(args []string) {
				fmt.Println(strings.Join(args, " "))
			},
		}

		fmt.Print("$ ")

		// Wait for user inputs
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		input := strings.TrimSpace(command)
		parts := strings.Fields(input)

		input_command := parts[0]
		input_arguments := parts[1:]

		if command_func, ok := allowed_commands[input_command]; ok {
			command_func(input_arguments)
		} else {
			fmt.Println(input_command + ": command not found")
		}
	}
}
