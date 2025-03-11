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
		"github.com/oklog/ulid/v2",
		"golang.org/x/exp/rand",
		"github.com/golang-jwt/jwt",
		"golang.org/x/crypto/bcrypt",
		"github.com/gin-gonic/gin",
	}

	for _, p := range package_ {
		runCommand(p)
	}
}

func CreateInit(dir string) {
	projectName := getProjectName(dir)
	initProject(projectName)
	path := fmt.Sprintf("%v/%v", dir, "app")
	if _, err := os.Stat(PathNormalization(path)); os.IsNotExist(err) {
		status := createDirectory(PathNormalization(path), "app")
		CreateFile(PathNormalization(dir), "main.go", strings.TrimSpace(code.MainCode(projectName)))
		if status {
			createDirectory(PathNormalization(fmt.Sprintf("%v/%v", path, "db")), "db")
			CreateFile(PathNormalization(fmt.Sprintf("%v/%v", path, "db")), "connection.go", strings.TrimSpace(code.FileConnection()))
			CreateFile(PathNormalization(path), "run.go", strings.TrimSpace(code.Run()))
			path = fmt.Sprintf("%v/%v", path, "service")
			createDirectory(PathNormalization(path), "service")
			CreateFile(PathNormalization(path), "base.go", strings.TrimSpace(code.Base()))
			CreateFile(PathNormalization(path), "validator.go", strings.TrimSpace(code.Validation()))
		}
		installMinLibrary()
	} else {
		fmt.Println("app exist...")
	}

}

func Service(dir string, serviceName string) {
	path := fmt.Sprintf("%v/app", dir)
	if _, err := os.Stat(PathNormalization(path)); os.IsNotExist(err) {
		fmt.Println("init before service")
	} else {
		serviceNameSplit := strings.Split(serviceName, "/")
		path = fmt.Sprintf("%v/%v", path, "service")
		for i := 0; i < len(serviceNameSplit); i++ {
			path = fmt.Sprintf("%v/%v", path, serviceNameSplit[i])
			if _, err = os.Stat(PathNormalization(path)); os.IsNotExist(err) {
				status := createDirectory(path, serviceNameSplit[i])
				if status && (len(serviceNameSplit)-1 == i) {
					CreateFile(path, "repository.go", strings.TrimSpace(code.Repository(serviceName)))
					CreateFile(path, "service.go", strings.TrimSpace(code.Service(serviceName)))
					CreateFile(path, "handler.go", strings.TrimSpace(code.Handler(serviceName)))
					CreateFile(path, "router.go", strings.TrimSpace(code.Router(serviceName)))
				}

			}
		}
	}
}

func hashPassword(dir string) {
	path := fmt.Sprintf("%v/%v", dir, "app")
	if _, err := os.Stat(PathNormalization(path)); os.IsNotExist(err) {
		fmt.Println("init before middleware")
	} else {
		path = fmt.Sprintf("%v/%v", path, "helper")
		status := createDirectory(PathNormalization(path), "helper")
		if status {
			CreateFile(path, "hashPassword.go", strings.TrimSpace(code.HashPassword()))
			CreateFile(path, "token.go", strings.TrimSpace(code.Token()))
		}

	}
}

func middleware(dir string, projectName string) {
	path := fmt.Sprintf("%v/%v", dir, "app")
	if _, err := os.Stat(PathNormalization(path)); os.IsNotExist(err) {
		fmt.Println("init before middleware")
	} else {
		path = fmt.Sprintf("%v/%v", path, "middleware")
		if _, err = os.Stat(PathNormalization(path)); os.IsNotExist(err) {
			status := createDirectory(PathNormalization(path), projectName)
			if status {
				CreateFile(path, "middleware.go", strings.TrimSpace(code.Middleware(projectName)))
				CreateFile(path, "model.go", strings.TrimSpace(code.ModelBlackListToken(projectName)))
				CreateFile(path, "repository.go", strings.TrimSpace(code.RepositoryBlackListToken()))
			}
		}
	}
}

func Auth(dir string) {
	path := fmt.Sprintf("%v", dir)
	projectName := getProjectName(path)
	hashPassword(path)
	middleware(path, projectName)
	addBlackListToken(dir+"/app/db/connection.go", projectName)
}
