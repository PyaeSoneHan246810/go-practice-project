package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func CreateFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		err = errors.New("failed to create file")
		return nil, err
	}
	return file, nil
}

func OpenFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		err = errors.New("failed to open file")
		return nil, err
	}
	return file, nil
}

func ReadLines(file *os.File) ([]string, error) {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		err = errors.New("failed to read file")
		return nil, err
	}
	return lines, nil
}

func WriteJson(file *os.File, data any) error {
	encoder := json.NewEncoder(file)
	err := encoder.Encode(data)
	if err != nil {
		err = errors.New("failed to encode data")
		return err
	}
	return nil
}
