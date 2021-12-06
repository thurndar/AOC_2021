package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed test_input.txt
var input []byte

func main() {
	measurments := scan()
	fmt.Print("Solution: ")
	drawField(measurments)
}

func scan() (measurments []string) {
	s := bufio.NewScanner(bytes.NewReader(input))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		measurments = append(measurments, s.Text())
	}
	return measurments
}

func drawField(measurments []string) (points int) {

}
