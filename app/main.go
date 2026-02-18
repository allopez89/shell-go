package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var allowed_commands map[string]func([]string)

func main() {

	allowed_commands = map[string]func(args []string){
		"exit": command_exit,
		"echo": command_echo,
		"type": command_type,
	}

	for {
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

func command_exit(args []string) { os.Exit(0) }

func command_echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func command_type(args []string) {
	if _, ok := allowed_commands[args[0]]; ok {
		fmt.Println(args[0] + " is a shell builtin")
		return
	}
	full_path, err := exec.LookPath(args[0])
	if err != nil {
		fmt.Println(args[0] + ": not found")
	} else {
		fmt.Println(args[0] + " is " + full_path)
	}
}
