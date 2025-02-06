package main

import (
	"flag"
	"fmt"
	"gin-assistant/command"
	"os"
	"strings"
)

/**
*-appv
*|--- dbv
*|------connectionv
*|---service
*|-----folder_service
*|-----basev
*|-----validatorv
*|--run
*|main
**/

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		command_ := args[0]
		dir, _ := os.Getwd()
		splitDir := strings.Split(dir, "/")
		projectName := splitDir[len(splitDir)-1]
		switch command_ {
		case "init":
			command.CreateInit(dir)
		case "service":
			if len(args) >= 2 {
				command.Service(dir, args[1])
			} else {
				fmt.Println("./gin-assistant service [nameService]")
			}
		case "auth":
			command.Auth(dir, projectName)
		default:
			fmt.Println("command not found")
		}
	} else {
		fmt.Println("uuups")
	}

}
