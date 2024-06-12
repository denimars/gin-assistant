package command

import (
	"fmt"
	"os"
)

func createDirectory(path string, name string) (bool, string) {
	fmt.Println("assistant created folder " + name)
	err := os.Mkdir(fmt.Sprintf("%v/%v", path, name), 0755)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return true, name
}
