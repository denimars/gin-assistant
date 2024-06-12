package command

import (
	"fmt"
	"os"
)

func CreateFile(dir string, fileName string, text string) {
	fmt.Println(dir)
	file, err := os.Create(fmt.Sprintf("%v/%v", dir, fileName))
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
