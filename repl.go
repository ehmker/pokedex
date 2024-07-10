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
		line := reader.Text()

		if cmd, ok := commands[strings.ToLower(line)]; ok {
			cmd.Callback(cmd.Config)
		} else {
			fmt.Printf("'%v' command not found\n", line)
		}
	}

}