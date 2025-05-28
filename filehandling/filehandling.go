package filehandling

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
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

//using Json file to store the datas

func WriteJSON(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("Failed to create a file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON")
	}
	file.Close()
	return nil

}
