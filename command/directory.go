package command

import (
	"fmt"
	"os"
)

func createDirectory(path string, name string) bool {
	fmt.Println("assistant created folder " + name)
	err := os.Mkdir(path, 0755)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return true
}
