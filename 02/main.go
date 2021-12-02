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
	fmt.Print("Position: \n")
	fmt.Print(calculatePosition(measurments))

	fmt.Print("\n Position with Aim: \n")
	fmt.Print(calculatePositionWithAim(measurments))
}

func scan() (measurments []string) {
	s := bufio.NewScanner(bytes.NewReader(input))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		measurments = append(measurments, s.Text())
	}
	return measurments
}

func calculatePosition(measurments []string) (counter int) {

	xPosition := 0
	yPosition := 0

	for i := 0; i < len(measurments); i++ {
		splitted := strings.Split(measurments[i], " ")
		positionValue, err := strconv.Atoi(splitted[1])
		if err != nil {
			log.Fatalf("unable to scan line: %v", err)
		}
		if splitted[0] == "down" {
			yPosition = yPosition + positionValue
		} else if splitted[0] == "up" {
			yPosition = yPosition - positionValue
		} else {
			xPosition = xPosition + positionValue
		}

	}

	return xPosition * yPosition
}

func calculatePositionWithAim(measurments []string) (counter int) {

	xPosition := 0
	depth := 0
	aim := 0

	for i := 0; i < len(measurments); i++ {
		splitted := strings.Split(measurments[i], " ")
		positionValue, err := strconv.Atoi(splitted[1])
		if err != nil {
			log.Fatalf("unable to scan line: %v", err)
		}
		if splitted[0] == "down" {
			aim = aim + positionValue
		} else if splitted[0] == "up" {
			aim = aim - positionValue
		} else {
			xPosition = xPosition + positionValue
			depth = depth + (positionValue * aim)
		}

	}

	return xPosition * depth
}
