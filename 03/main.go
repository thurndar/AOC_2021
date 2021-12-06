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
	fmt.Print("Power Consumption: \n")
	fmt.Print(calculatePowerConsumption(measurments, len(measurments[0])))
	fmt.Print("\nLife Support Rating: \n")
	fmt.Print(calculateLifeSupportRating(measurments, len(measurments[0])))
}

func scan() (measurments []string) {
	s := bufio.NewScanner(bytes.NewReader(input))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		measurments = append(measurments, s.Text())
	}
	return measurments
}

func getGammaRate(measurments []string, stringlen int) (gammaRate string) {
	gammaRateString := ""

	for position := 0; position < stringlen; position++ {
		zeros := 0
		ones := 0

		for i := 0; i < len(measurments); i++ {
			if string(measurments[i][position]) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if zeros > ones {
			gammaRateString = gammaRateString + "0"
		} else {
			gammaRateString = gammaRateString + "1"
		}
	}

	return gammaRateString
}

func getEpsilonRate(measurments []string, stringlen int) (epsilonRate string) {
	epsilonRateString := ""

	for position := 0; position < stringlen; position++ {
		zeros := 0
		ones := 0

		for i := 0; i < len(measurments); i++ {
			if string(measurments[i][position]) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if zeros > ones {
			epsilonRateString = epsilonRateString + "1"
		} else {
			epsilonRateString = epsilonRateString + "0"
		}
	}

	return epsilonRateString
}

func calculatePowerConsumption(measurments []string, stringlen int) (powerConsumption int64) {
	gammaRateString := getGammaRate(measurments, stringlen)
	epsilonRateString := getEpsilonRate(measurments, stringlen)

	gammaRate, error := strconv.ParseInt(gammaRateString, 2, 64)
	if error != nil {
		log.Fatalf("unable to scan line: %v", error)
	}
	epsilonRate, error := strconv.ParseInt(epsilonRateString, 2, 64)
	if error != nil {
		log.Fatalf("unable to scan line: %v", error)
	}

	return gammaRate * epsilonRate
}

func getRating(measurments []string, commonValueString string) (rating int64) {
	modifiedMeasurments := make([]string, len(measurments))
	copy(modifiedMeasurments, measurments)
	commonString := ""
	removedStrings := len(modifiedMeasurments)
	boolBreak := false
	// fmt.Print("\nString: ")
	// fmt.Print(commonValueString)
	// fmt.Print("\nRemoved Strings:")
	// fmt.Print(removedStrings)
	// fmt.Print("\n")
	position := 0
	length := len(modifiedMeasurments[0]) - 1
	for !boolBreak && position < length {
		for i := 0; i < len(modifiedMeasurments); i++ {
			if removedStrings == 1 {
				// fmt.Print("I have entered the break")
				boolBreak = true
				break
			}
			if modifiedMeasurments[i] != "None" {
				// fmt.Print("\nmod Meas: ")
				// fmt.Print(modifiedMeasurments[i][position])
				// fmt.Print("\n common: ")
				// fmt.Print(commonValueString[position])
				if modifiedMeasurments[i][position] != commonValueString[position] {
					modifiedMeasurments[i] = "None"
					removedStrings--
					// fmt.Print("\n Removed a string now: ")
					// fmt.Print(removedStrings)
					// fmt.Print("\n")
				}
			}
		}
		position++
	}
	if removedStrings > 1 {
		for i := 0; i < len(modifiedMeasurments); i++ {
			if modifiedMeasurments[i] != "None" {
				commonString = modifiedMeasurments[i]
			}
		}
	} else {
		for i := 0; i < len(modifiedMeasurments); i++ {
			if modifiedMeasurments[i] != "None" {
				commonString = modifiedMeasurments[i]
			}
		}
	}
	// fmt.Print("\n Common string: ")
	// fmt.Print(commonString)
	// fmt.Print(modifiedMeasurments)
	rating, error := strconv.ParseInt(commonString, 2, 64)
	if error != nil {
		log.Fatalf("unable to scan line: %v", error)
	}
	return rating
}

func calculateLifeSupportRating(measurments []string, stringlen int) (lifeSupportRating int64) {
	mostCommonValueString := getGammaRate(measurments, stringlen)
	leastCommonValueString := getEpsilonRate(measurments, stringlen)

	oxygenRating := getRating(measurments, mostCommonValueString)
	// fmt.Print("\nOxygenRating: ")
	// fmt.Print(oxygenRating)
	// fmt.Print("\n")
	coRating := getRating(measurments, leastCommonValueString)
	// fmt.Print("\ncoRating: ")
	// fmt.Print(coRating)
	// fmt.Print("\n")

	return oxygenRating * coRating
}
