package main

import (
	"flag"
	"fmt"
	"gin-assistant/command"
	"os"
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
		switch command_ {
		case "init":
			command.CreateInit(dir)
		case "service":
			if len(args) >= 2 {
				command.Service(dir, args[1])
			}
			fmt.Println("./gin-assistant service [nameService]")

		default:
			fmt.Println("command not found")
		}
	} else {
		fmt.Println("uuups")
	}

}
