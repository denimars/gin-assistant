package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dir string

func readFile(searchFile string, fixedFile string, isALl bool) []string {
	var result []string
	file, err := os.Open(dir)
	if err != nil {
		fmt.Println("file not open", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchFile) {
			if isALl {
				result = append(result, fixedFile)
			} else {
				modifiedLine := strings.Replace(line, searchFile, fixedFile, 1)
				result = append(result, modifiedLine)
			}

		} else {
			result = append(result, line)
		}

	}
	return result
}

func reWriteFile(content []string) {
	connection, err := os.Create(dir)
	if err != nil {
		fmt.Println("error open file")
		return
	}
	defer connection.Close()
	for _, line := range content {
		_, err := connection.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error write file:", err)
			return
		}
	}
}

func addBlackListToken(file string, projectName string) {
	dir = file
	addBlackListToken := readFile("db.AutoMigrate(", "db.AutoMigrate(\n  middleware.BlacklistToken{},", false)
	reWriteFile(addBlackListToken)
	addImportFile := readFile("import (", "import (\n  \""+projectName+"/app/middleware\"", false)
	reWriteFile(addImportFile)
}

func ReWritePort(port string, file string) {
	dir = PathNormalization(fmt.Sprintf("%v/app/run.go", file))
	newPort := fmt.Sprintf("port := \"%v\"", port)
	addNewPort := readFile("port := \"", newPort, true)
	reWriteFile(addNewPort)
}
