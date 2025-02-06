package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dir string

func readFile(projectName string) []string {
	var result []string
	file, err := os.Open(dir)
	if err != nil {
		fmt.Println("file not open", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "db.AutoMigrate(") {
			modifiedLine := strings.Replace(line, "db.AutoMigrate(", "db.AutoMigrate(\n  middleware.BlacklistToken{},", 1)
			result = append(result, modifiedLine)
		} else if strings.Contains(line, "import (") {
			modifiedLine := strings.Replace(line, "import (", "import (\n  \""+projectName+"/app/middleware\"", 1)
			result = append(result, modifiedLine)
		} else {
			result = append(result, line)
		}
	}
	return result
}

func reWriteConnection(content []string) {
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
	newFile := readFile(projectName)
	reWriteConnection(newFile)
}
