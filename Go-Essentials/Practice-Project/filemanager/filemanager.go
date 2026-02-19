package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct{
	InputFilePath string
	OutputFilePath string
} 


func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("Faild to open file.")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Faild to read Lines in file.")
	}
	file.Close()
	return lines, nil

}

func (fm FileManager) WriteResult(data interface{}) error{
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Faild to Create file.")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Faild to Convert data to JSON.")
	}
	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager{
	return FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}
