package services

import (
	"fmt"
	"io"
	"os"
)

func ReadJson(pathFile string) ([]byte, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(pathFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	return byteValue, err
}
