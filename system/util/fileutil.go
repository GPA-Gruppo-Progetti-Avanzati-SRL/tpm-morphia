package util

import (
	"bufio"
	"io/ioutil"
	"math"
	"os"
)

func CopyFile(sourceFile string, destinationFile string) (int, error) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		return 0, err
	}

	return len(input), nil
}

func ReadTextLines(fileName string, numberOfLines int) ([]string, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var initialCapacity = numberOfLines
	if numberOfLines < 0 {
		initialCapacity = 100
		numberOfLines = math.MaxInt32
	} else if numberOfLines == 0 {
		initialCapacity = 1
		numberOfLines = 1
	}

	lines := make([]string, 0, initialCapacity)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && len(lines) < numberOfLines {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}

	/*
		for i, s := range lines {
			fmt.Printf("[%d] - %s\n", i, s)
		}
	*/

	return lines, nil
}
