package filehandling

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileHandling struct {
	InputFilePath  string
	OutputFilePath string
}

func (fh FileHandling) ReadLines() ([]string, error) {
	file, err := os.Open(fh.InputFilePath)
	if err != nil {
		return nil, errors.New("Failed to open file ")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read line in the file ")
	}
	file.Close()
	return lines, nil
}
func (fh FileHandling) WriteResult(data any) error {
	file, err := os.Create(fh.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create a file")
	}
	time.Sleep(3 * time.Second) // to slow down the process of writing result in data

	
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON")
	}
	file.Close()
	return nil

}
func New(inputpath, outputPath string) FileHandling {
	return FileHandling{
		InputFilePath:  inputpath,
		OutputFilePath: outputPath,
	}

}
