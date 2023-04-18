package data_source

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const rootDirectory = "./resource/data/"

func GetFileData(fileName string) string {

	concatFileName := rootDirectory + fileName

	err := filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if string(concatFileName[2:]) == path {
			return nil
		}

		return err
	})

	if err == nil {
		return readingFileContent(concatFileName)
	}

	return ""
}

func readingFileContent(fileName string) string {

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return string(fileContent)
}
