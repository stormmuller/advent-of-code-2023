package utils

import (
	"log"
	"os"
)

func WriteToFile(path string, content string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := []byte(content)

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
