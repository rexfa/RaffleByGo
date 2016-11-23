package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "settings.inf"
	fin, err := os.Open(userFile)
	defer fin.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

}
