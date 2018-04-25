package utilities

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// GetListFiles - функция для получение списка всех файлов в папке
func GetListFiles(root string, extentions []string, onlyFolder bool) []string {

	var files []string

	error := filepath.Walk(root, func(path string, info os.FileInfo, error error) error {

		for _, extention := range extentions {

			if filepath.Ext(info.Name()) == extention {
				if onlyFolder {
					if info.IsDir() {
						files = append(files, path)
					}
				} else {
					if !info.IsDir() {
						files = append(files, path)
					}
				}

			}

		}

		return nil

	})

	if error != nil {
		log.Fatal("Произошла ошибка.")
	}

	return files
}

// ReadFile - функция для чтения файла
func ReadFile(path string) []byte {

	// открытие файла
	file, error := os.Open(path)
	defer file.Close()
	if error != nil {
		fmt.Println("Ошибка открытия файла.")
	}

	// чтение файла
	data, error := ioutil.ReadAll(file)
	if error != nil {
		log.Fatal("Ошибка чтения файла.")
	}

	return data
}
