package command

import (
	"fmt"
	"gin-assistant/code"

	"os"
	"os/exec"
	"strings"
)

func getProjectName(dir string) string {
	splitFolder := strings.Split(dir, "/")
	return splitFolder[len(splitFolder)-1]
}

func runCommand(command string) {
	// splitCommand := strings.Split(command, " ")
	// fmt.Println(splitCommand)
	// cmd := exec.Command("go", splitCommand...)
	cmd := exec.Command("go", "get", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		panic("uups....")
	}
	fmt.Println(string(output))
}

func installMinLibrary() {
	package_ := []string{
		"github.com/joho/godotenv",
		"gorm.io/driver/mysql",
		"gorm.io/gorm",
		"gorm.io/gorm/logger",
		"github.com/google/uuid",
		"github.com/go-playground/validator/v10",
		"github.com/gin-contrib/cors",
		"github.com/gin-gonic/gin",
	}

	for _, p := range package_ {
		runCommand(p)
	}
}

func CreateInit(dir string) {
	if _, err := os.Stat(dir + "/" + "app"); os.IsNotExist(err) {
		status, name := createDirectory(dir, "app")
		installMinLibrary()
		CreateFile(dir, "main.go", code.MainCode(getProjectName(dir)))
		if status {
			dir = fmt.Sprintf("%v/%v", dir, name)
			_, name = createDirectory(dir, "db")
			CreateFile(fmt.Sprintf("%v/%v", dir, name), "connection.go", code.FileConnection())
			_, name = createDirectory(dir, "service")
			CreateFile(fmt.Sprintf("%v/%v", dir, name), "base.go", code.Base())
			CreateFile(fmt.Sprintf("%v/%v", dir, name), "validator.go", code.Validation())
			CreateFile(fmt.Sprintf("%v/", dir), "run.go", code.Run())
		}
	} else {
		fmt.Println("app exist...")
	}

}

func Service(dir string, serviceName string) {
	if _, err := os.Stat(dir + "/" + "app"); os.IsNotExist(err) {
		fmt.Println("init before service")
	} else {
		if _, err = os.Stat(dir + "/app/service/" + serviceName); os.IsNotExist(err) {
			status, name := createDirectory(dir+"/app/service/", serviceName)
			if status {
				dirService := dir + "/app/service/" + name
				CreateFile(dirService, "repository.go", code.Repository(serviceName))
				CreateFile(dirService, "service.go", code.Service(serviceName))
				CreateFile(dirService, "handler.go", code.Handler(serviceName))
				CreateFile(dirService, "router.go", code.Router(serviceName))
			}

		}
	}
}
