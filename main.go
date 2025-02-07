package main

import (
	"flag"
	"fmt"
	"gin-assistant/command"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
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

var cmd *exec.Cmd

func run() {
	if cmd != nil {
		_ = cmd.Process.Kill()
	}
	cmd = exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	fmt.Println("Server restarted at", time.Now())
}

func watchFiles(watcher *fsnotify.Watcher, dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fmt.Println("👀 Watching:", path)
			return watcher.Add(path)
		}
		return nil
	})
}

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
			} else {
				fmt.Println("./gin-assistant service [nameService]")
			}
		case "auth":
			command.Auth(dir)
		case "run":
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.Fatal(err)
			}
			defer watcher.Close()

			projectDir := "."
			err = watchFiles(watcher, projectDir)
			if err != nil {
				log.Fatal(err)
			}
			run()
			fmt.Println("🚀 Watching for file changes...")

			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Op == fsnotify.Write || event.Op == fsnotify.Create || event.Op == fsnotify.Remove {
						fmt.Println("🔄 File changed:", event.Name)
						run()
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					fmt.Println("❌ Watcher error:", err)
				}
			}
		default:
			fmt.Println("command not found")
		}
	} else {
		fmt.Println("uuups")
	}

}
