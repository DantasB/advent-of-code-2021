package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pulseIncreased(previous int, current int) bool {

	return previous < current
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

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("[ERROR] Can't read the file %v", fileName)
		panic(err)
	}

	return file
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR] Missing parameter. You must specify a file path.")
		return
	}

	numbers := readNumbers(os.Args[1])

	result := checkNumberOfIncreses(numbers)

	fmt.Printf("[INFO] The obtained result is: %v", result)
}

func checkNumberOfIncreses(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers)-3; i++ {
		if pulseIncreased(numbers[i], numbers[i+3]) {
			result++
		}
	}

	return result
}

func readNumbers(fileName string) []int {
	file := readFile(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, convertToInteger(scanner.Text()))
	}

	return numbers
}
