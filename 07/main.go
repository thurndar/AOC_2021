package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	measurments := scan()
	fmt.Print("Solution 1: \n")
	fmt.Print(calculateCrabLine(measurments[0]))
	fmt.Print("\nSolution 2: \n")
	fmt.Print(calculateCrabLineExpanded(measurments[0]))
}

func scan() (measurments []string) {
	s := bufio.NewScanner(bytes.NewReader(input))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		measurments = append(measurments, s.Text())
	}
	return measurments
}

func calculateCrabLineExpanded(actualPosition string) (leastAmountOfFuel int) {
	splitted := strings.Split(actualPosition, ",")
	leastAmountOfFuel = 999999999
	for position := 0; position <= getHighestValue(actualPosition); position++ {
		var floatPosition float64 = float64(position)
		actualAmountOfFuel := 0
		for j := 0; j < len(splitted); j++ {
			crabPosition, err := strconv.ParseFloat(splitted[j], 64)
			if err != nil {
				log.Fatalf("unable to scan line: %v", err)
			}
			var distance float64 = math.Abs(floatPosition - crabPosition)
			usedFuel := 0
			for i := 1; i <= int(distance); i++ {
				usedFuel = usedFuel + i
			}
			actualAmountOfFuel = actualAmountOfFuel + usedFuel
		}
		if actualAmountOfFuel < leastAmountOfFuel {
			leastAmountOfFuel = actualAmountOfFuel
		}
	}

	return leastAmountOfFuel
}

func calculateCrabLine(actualPosition string) (leastAmountOfFuel float64) {
	splitted := strings.Split(actualPosition, ",")
	leastAmountOfFuel = 999999
	for position := 0; position <= getHighestValue(actualPosition); position++ {
		var floatPosition float64 = float64(position)
		var actualAmountOfFuel float64 = 0
		for j := 0; j < len(splitted); j++ {
			crabPosition, err := strconv.ParseFloat(splitted[j], 64)
			if err != nil {
				log.Fatalf("unable to scan line: %v", err)
			}
			var usedFuel float64 = math.Abs(floatPosition - crabPosition)
			actualAmountOfFuel = actualAmountOfFuel + usedFuel
		}
		if actualAmountOfFuel < leastAmountOfFuel {
			leastAmountOfFuel = actualAmountOfFuel
		}
	}

	return leastAmountOfFuel
}

func getHighestValue(actualPosition string) (highestValue int) {
	splitted := strings.Split(actualPosition, ",")
	highestValue = 0
	for i := 0; i < len(splitted); i++ {
		crabPosition, err := strconv.Atoi(splitted[i])
		if err != nil {
			log.Fatalf("unable to scan line: %v", err)
		}
		if crabPosition > highestValue {
			highestValue = crabPosition
		}
	}

	return highestValue
}
