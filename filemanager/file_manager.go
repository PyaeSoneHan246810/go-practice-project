package filemanager

import (
	"bufio"
	"errors"
	"os"
)

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
