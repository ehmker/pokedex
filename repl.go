package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex")
	
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		
		inputCMD, inputLocation := processInput(reader.Text())

		if cmd, ok := commands[inputCMD]; ok {
			cmd.Callback(cmd.Config, inputLocation)
		} else {
			fmt.Printf("'%v' command not found\n", inputCMD)
		}
	}

}

func processInput(user_input string) (command, location string) {
	user_input = strings.ToLower(user_input)
	split_input := strings.Split(user_input, " ")

	if len(split_input) == 2{
		return split_input[0], split_input[1]
	}
	return split_input[0], ""

	
}