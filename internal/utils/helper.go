package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func GetEnv(varName string, isRequired bool) string {
	value := os.Getenv(varName)

	if value == "" && isRequired {
		log.Fatalf("%s env variable is required", varName)
	}

	return value
}

func StoreFile(folder string, fileName string, file io.Reader) error {
	os.MkdirAll(folder, os.ModePerm)

	path := fmt.Sprintf("%s/%s", folder, fileName)

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	io.Copy(dst, file)
	return nil
}
