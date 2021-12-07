package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("[ERROR] Can't read the file %v", fileName)
		panic(err)
	}

	return file
}

func convertToInteger(line string) int {
	trimmedLine := strings.TrimSuffix(line, "\n")
	converted, err := strconv.Atoi(trimmedLine)
	if err != nil {
		fmt.Printf("[ERROR] Can't convert the value %v to integer", trimmedLine)
		panic(err)
	}

	return converted
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR] Missing parameter. You must specify a file path.")
		return
	}

	horizontal, depth := readPosition(os.Args[1])

	result := horizontal * depth

	fmt.Printf("[INFO] The obtained result is: %v", result)
}

func readPosition(fileName string) (int, int) {
	horizontal, depth := 0, 0

	file := readFile(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		checkPosition(line, &horizontal, &depth)
	}

	return horizontal, depth
}

func checkPosition(line string, horizontal *int, depth *int) {
	splittedLine := strings.Split(line, " ")
	if len(splittedLine) != 2 {
		fmt.Println("[ERROR] The line is not valid.")
		return
	}

	units := convertToInteger(splittedLine[1])
	if splittedLine[0] == "up" {
		*depth -= units
	} else if splittedLine[0] == "down" {
		*depth += units
	} else if splittedLine[0] == "forward" {
		*horizontal += units
	} else {
		fmt.Println("[ERROR] Invalid movement.")
	}

}
