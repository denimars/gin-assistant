package command

import (
	"fmt"
	"gin-assistant/code"
	"runtime"

	"os"
	"os/exec"
	"strings"
)

func getProjectName(dir string) string {
	var splitFolder []string
	switch runtime.GOOS {
	case "windows":
		splitFolder = strings.Split(dir, "\\")
	default:
		splitFolder = strings.Split(dir, "/")
	}

	return splitFolder[len(splitFolder)-1]
}

func initProject(projectName string) {
	cmd := exec.Command("go", "mod", "init", projectName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		panic("uups....")
	}
	fmt.Println(string(output))

}

func runCommand(command string) {
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
		"github.com/oklog/ulid/v2",
		"golang.org/x/exp/rand",
		"github.com/golang-jwt/jwt",
		"golang.org/x/crypto/bcrypt",
	}

	for _, p := range package_ {
		runCommand(p)
	}
}

func CreateInit(dir string) {
	projectName := getProjectName(dir)
	initProject(projectName)

	if _, err := os.Stat(dir + "/" + "app"); os.IsNotExist(err) {
		status, name := createDirectory(dir, "app")
		installMinLibrary()
		CreateFile(dir, "main.go", strings.TrimSpace(code.MainCode(projectName)))
		if status {
			dir = fmt.Sprintf("%v/%v", dir, name)
			_, name = createDirectory(dir, "db")
			CreateFile(fmt.Sprintf("%v/%v", dir, name), "connection.go", strings.TrimSpace(code.FileConnection()))
			_, name = createDirectory(dir, "service")
			CreateFile(fmt.Sprintf("%v/%v", dir, name), "base.go", strings.TrimSpace(code.Base()))
			CreateFile(fmt.Sprintf("%v/%v", dir, name), "validator.go", strings.TrimSpace(code.Validation()))
			CreateFile(fmt.Sprintf("%v/", dir), "run.go", strings.TrimSpace(code.Run()))
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
				CreateFile(dirService, "repository.go", strings.TrimSpace(code.Repository(serviceName)))
				CreateFile(dirService, "service.go", strings.TrimSpace(code.Service(serviceName)))
				CreateFile(dirService, "handler.go", strings.TrimSpace(code.Handler(serviceName)))
				CreateFile(dirService, "router.go", strings.TrimSpace(code.Router(serviceName)))
			}

		}
	}
}

func hashPassword(dir string) {
	if _, err := os.Stat(dir + "/" + "app"); os.IsNotExist(err) {
		fmt.Println("init before middleware")
	} else {
		if _, err = os.Stat(dir + "/app/helper"); os.IsNotExist(err) {
			status, name := createDirectory(dir+"/app", "helper")
			if status {
				dirHalper := dir + "/app/" + name
				CreateFile(dirHalper, "hashPassword.go", strings.TrimSpace(code.HashPassword()))
				CreateFile(dirHalper, "token.go", strings.TrimSpace(code.Token()))
			}
		}
	}
}

func middleware(dir string, projectName string) {
	if _, err := os.Stat(dir + "/" + "app"); os.IsNotExist(err) {
		fmt.Println("init before middleware")
	} else {
		if _, err = os.Stat(dir + "/app/middleware"); os.IsNotExist(err) {
			status, name := createDirectory(dir+"/app", "middleware")
			if status {
				dirHalper := dir + "/app/" + name
				CreateFile(dirHalper, "middleware.go", strings.TrimSpace(code.Middleware(projectName)))
				CreateFile(dirHalper, "model.go", strings.TrimSpace(code.ModelBlackListToken(projectName)))
				CreateFile(dirHalper, "repository.go", strings.TrimSpace(code.RepositoryBlackListToken()))
			}
		}
	}
}

func Auth(dir string) {
	projectName := getProjectName(dir)
	hashPassword(dir)
	middleware(dir, projectName)
	addBlackListToken(dir+"/app/db/connection.go", getProjectName(dir))
}
