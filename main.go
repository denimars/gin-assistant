package main

import (
	"flag"
	"fmt"
	"gin-assistant/command"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
)

var cmd *exec.Cmd
var mutex sync.Mutex
var debounceTimer *time.Timer
var serverPort string

func stopServer() {
	if cmd != nil && cmd.Process != nil {
		fmt.Println("🛑 Stopping server...")
		if runtime.GOOS == "windows" {
			err := cmd.Process.Kill()
			if err != nil {
				fmt.Println("❌ Error killing process:", err)
			}
		} else {
			err := cmd.Process.Signal(syscall.SIGTERM)
			if err != nil {
				fmt.Println("⚠️ Error sending SIGTERM:", err)
			}
			_, _ = cmd.Process.Wait()
		}

		if processExists(cmd.Process.Pid) {
			fmt.Println("⚠️ Force killing process...")
			err := cmd.Process.Kill()
			if err != nil {
				fmt.Println("❌ Error force killing process:", err)
			}
		}
	}
	freePort(serverPort)
}

func processExists(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func freePort(port string) {
	switch runtime.GOOS {
	case "windows":
		freePortWindows(port)
	default:
		freePortUnix(port)
	}
}

func freePortWindows(port string) {
	cmd := exec.Command("cmd", "/C", "netstat -ano | findstr :"+port)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("⚠️ Error checking port:", err)
		return
	}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 4 {
			pid := fields[len(fields)-1]
			fmt.Println("🔄 Killing process using port:", port, "PID:", pid)
			killCmd := exec.Command("taskkill", "/F", "/PID", pid)
			killCmd.Run()
		}
	}
}

func freePortUnix(port string) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("lsof -t -i:%s | xargs kill -9", port))
	cmd.Run()
}

func run() {
	mutex.Lock()
	defer mutex.Unlock()
	stopServer()
	time.Sleep(1 * time.Second)
	cmd = exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
	fmt.Println("🚀 Server restarted at", time.Now())
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

func debounceRestart() {
	if debounceTimer != nil {
		debounceTimer.Stop()
	}
	debounceTimer = time.AfterFunc(1*time.Second, run)
}

func setPort(args []string) bool {
	godotenv.Load()
	if os.Getenv("PORT") == "" {
		if len(args) >= 2 {
			port := args[1]
			if num, err := strconv.Atoi(port); err == nil && num >= 7000 && num <= 9999 {
				serverPort = port
				return true
			}
			fmt.Println("Port must be a number between 7000 - 9999")
			return false
		}

		serverPort = "8080"
		return true
	}
	serverPort = os.Getenv("PORT")
	return true
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
			if setPort(args) {
				fmt.Println("****\n", serverPort, "\n", dir, "\n****")
				command.ReWritePort(serverPort, dir)
				watcher, err := fsnotify.NewWatcher()
				if err != nil {
					log.Fatal(err)
				}
				defer watcher.Close()
				projectDir := command.PathNormalization("./app")
				if err := watchFiles(watcher, projectDir); err != nil {
					log.Fatal(err)
				}
				fmt.Println("🚀 Starting server...")
				run()
				fmt.Println("🚀 Watching for file changes...")
				for {
					select {
					case event, ok := <-watcher.Events:
						if !ok {
							return
						}
						fmt.Println("🔄 File changed:", event.Name)
						if event.Op&fsnotify.Create != 0 {
							if fileInfo, err := os.Stat(event.Name); err == nil && fileInfo.IsDir() {
								fmt.Println("📂 New folder detected, adding to watcher:", event.Name)
								if err := watcher.Add(event.Name); err != nil {
									fmt.Println("❌ Error adding new folder to watcher:", err)
								}
							}
						}
						debounceRestart()
					case err, ok := <-watcher.Errors:
						if !ok {
							return
						}
						fmt.Println("❌ read file error:", err)
					}
				}
			}
		default:
			fmt.Println("command not found")
		}
	} else {
		fmt.Println("uuups")
	}
}
