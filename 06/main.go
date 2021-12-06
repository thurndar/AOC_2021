package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	measurments := scan()
	fmt.Print("Solution : \n")
	lanternList := measurments[0]
	// fmt.Print("\nInitial: " + lanternList + "\n")
	for days := 1; days <= 256; days++ {
		// fmt.Print("\nDay " + strconv.Itoa(days) + " : ")
		lanternList = calculateLanterns(lanternList)
		// fmt.Print(lanternList + "\n")
	}
	fmt.Print(len(strings.Split(lanternList, ",")))
}

func scan() (measurments []string) {
	s := bufio.NewScanner(bytes.NewReader(input))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		measurments = append(measurments, s.Text())
	}
	return measurments
}

func calculateLanterns(state string) (list string) {
	lanternList := strings.Split(state, ",")
	list = ""
	amountOfNewLanterns := 0

	for i := 0; i < len(lanternList); i++ {
		actualDays, err := strconv.Atoi(lanternList[i])
		newDays := actualDays - 1
		if actualDays == 0 {
			newDays = 6
			amountOfNewLanterns++
		}
		if err != nil {
			log.Fatalf("unable to scan line: %v", err)
		}
		list = list + strconv.Itoa(newDays)
		if i != (len(lanternList) - 1) {
			list = list + ","
		}
	}

	for j := 0; j < amountOfNewLanterns; j++ {
		list = list + ",8"
	}

	return list
}
