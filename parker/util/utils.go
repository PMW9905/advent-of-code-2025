package utils

import (
	"bufio"
	"os"
)

func ReadInput(path string) ([]string, error) {

	file, err := os.OpenFile(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
