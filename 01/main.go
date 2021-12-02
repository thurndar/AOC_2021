package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strconv"
)

//go:embed input.txt
var input []byte

func main() {
	measurments := scan()
	fmt.Print("normal method: ")
	fmt.Print(countIncreases(measurments))
	fmt.Print("\n")
	fmt.Print("modified method: ")
	fmt.Print(countModIncreases(measurments))
}

func scan() (measurments []int) {
	s := bufio.NewScanner(bytes.NewReader(input))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("unable to scan line: %v", err)
		}
		measurments = append(measurments, i)
	}
	return measurments
}

func countIncreases(measurments []int) (counter int) {

	counter = 0
	for i := 1; i < len(measurments); i++ {
		if measurments[i] > measurments[i-1] {
			counter++
		}
	}

	return counter
}

func countModIncreases(measurments []int) (counter int) {
	tempCounterA := 0
	tempCounterB := 0
	for i := 0; i < len(measurments); i++ {
		if (i + 3) >= len(measurments) {
			break
		}
		tempCounterA = measurments[i] + measurments[i+1] + measurments[i+2]
		tempCounterB = measurments[i+1] + measurments[i+2] + measurments[i+3]
		if tempCounterB > tempCounterA {
			counter++
		}
	}

	return counter
}
