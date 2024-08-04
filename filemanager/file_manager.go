package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fileManager FileManager) CreateFile() (*os.File, error) {
	file, err := os.Create(fileManager.OutputFilePath)
	if err != nil {
		err = errors.New("failed to create file")
		return nil, err
	}
	return file, nil
}

func (fileManager FileManager) OpenFile() (*os.File, error) {
	file, err := os.Open(fileManager.InputFilePath)
	if err != nil {
		err = errors.New("failed to open file")
		return nil, err
	}
	return file, nil
}

func (fileManager FileManager) ReadLines() ([]string, error) {
	file, err := fileManager.OpenFile()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		err = errors.New("failed to read file")
		return nil, err
	}
	file.Close()
	return lines, nil
}

func (fileManager FileManager) WriteData(data any) error {
	file, err := fileManager.CreateFile()
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		err = errors.New("failed to encode data")
		return err
	}
	return nil
}
