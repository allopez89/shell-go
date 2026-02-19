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
		"pwd":  command_pwd,
		"cd":   command_cd,
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
		parts := parseInput(input)

		input_command := parts[0]
		input_arguments := parts[1:]

		if is_builtin(input_command) {
			command_func := allowed_commands[input_command]
			command_func(input_arguments)
		} else if command_name, ok := is_executable(input_command); ok {
			cmd := exec.Command(command_name)
			cmd.Args = parts
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			fmt.Println(input_command + ": command not found")
		}
	}
}

// builtin commands
func command_exit(args []string) { os.Exit(0) }

func command_echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func command_type(args []string) {
	if is_builtin(args[0]) {
		fmt.Println(args[0] + " is a shell builtin")
		return
	}

	if command_path, ok := is_executable(args[0]); ok {
		fmt.Println(args[0] + " is " + command_path)
		return
	}
	fmt.Println(args[0] + ": not found")
}

func command_pwd(args []string) {
	working_dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error retrieving directory: ", err)
	} else {
		fmt.Println(working_dir)
	}
}

func command_cd(directory []string) {
	home_directory, _ := os.UserHomeDir()
	var target string

	if len(directory) == 0 || directory[0] == "~" {
		target = home_directory
	} else {
		target = directory[0]
	}
	err := os.Chdir(target)
	if err != nil {
		fmt.Println("cd: " + target + ": No such file or directory")
		return
	}
}

// utils
func is_builtin(command string) bool {
	if _, ok := allowed_commands[command]; ok {
		return true
	} else {
		return false
	}
}

func is_executable(path string) (string, bool) {
	full_path, err := exec.LookPath(path)
	if err != nil {
		return "", false
	} else {
		return full_path, true
	}
}

func parseInput(input string) []string {
	var args []string
	var current strings.Builder
	var quoteChar rune = 0

	for _, r := range input {
		switch r {
		case '\'', '"':
			if quoteChar == 0 {
				quoteChar = r
			} else if quoteChar == r {
				quoteChar = 0
			} else {
				current.WriteRune(r)
			}

		case ' ':
			if quoteChar != 0 {
				current.WriteRune(r)
			} else {
				if current.Len() > 0 {
					args = append(args, current.String())
					current.Reset()
				}
			}

		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}
