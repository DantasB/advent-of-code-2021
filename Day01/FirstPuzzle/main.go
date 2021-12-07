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

func isFirstMeasurement(previousMeasurement *int, measure string) bool {
	if *previousMeasurement == 0 {
		*previousMeasurement = convertToInteger(measure)
		return true
	}

	return false
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

	file := readFile(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	previousMeasurement, currentMeasurement, result := 0, 0, 0
	for scanner.Scan() {
		measure := scanner.Text()
		if isFirstMeasurement(&previousMeasurement, measure) {
			continue
		}

		currentMeasurement = convertToInteger(measure)

		calculateDepthIncrease(previousMeasurement, currentMeasurement, &result)

		previousMeasurement = currentMeasurement
	}

	fmt.Printf("[INFO] The obtained result is: %v", result)
}

func calculateDepthIncrease(previousMeasurement, currentMeasurement int, result *int) {

	if pulseIncreased(previousMeasurement, currentMeasurement) {
		*result += 1
	}
}
