package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test_input.txt
var input []byte

type bingoCard struct {
	grid    [5]*[5]*num
	current int
}

type num struct {
	n     int
	drawn bool
}

func main() {
	measurments := scan()
	fmt.Print("Solution: ")
	playBingo(measurments)
}

func scan() (measurments []string) {
	s := bufio.NewScanner(bytes.NewReader(input))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		measurments = append(measurments, s.Text())
	}
	return measurments
}

func playBingo(measurments []string) (solution int64) {
	// bingoCalls := measurments[0]
	createBingoCards(measurments)

	return solution
}

func createBingoCards(measurments []string) (bingoCards []bingoCard) {
	// isCard := true
	// amountBingoCards := (len(measurments) - 3) / 6
	// bingoCards := make([]bingoCard, amountBingoCards)

	for i := 2; i < len(measurments); i = i + 6 {
		for j := 0; j < 4; j++ {
			row := strings.Split(measurments[i+j])

		}
		firstRow := strings.Split(measurments[i], " ")
		fmt.Print(firstRow)
		fmt.Print("\n")
	}

	return nil
}
