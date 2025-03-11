package command

import (
	"fmt"
	"os"
)

func CreateFile(dir string, fileName string, text string) {
	path := fmt.Sprintf("%v/%v", dir, fileName)
	file, err := os.Create(PathNormalization(path))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write([]byte(text))
	if err != nil {
		panic(err)
	}
	fmt.Println("assistant created file " + fileName)
}
