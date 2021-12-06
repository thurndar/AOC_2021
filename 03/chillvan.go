package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strconv"
)

var (
	//go:embed test_input.txt
	example []byte
	//go:embed input.txt
	input []byte
)

func main() {
	// consumption(scan(example))
	consumption(scan(input))
	oxi := bytesToInt(oxigenGeneratorRating(scan(input), 0)[0])
	scrub := bytesToInt(scrubberGeneratorRating(scan(input), 0)[0])
	fmt.Printf("oxi: %d | scrub: %d | res: %d", oxi, scrub, oxi*scrub)
}

func bytesToInt(bytes string) int64 {
	i, err := strconv.ParseInt(bytes, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func oxigenGeneratorRating(bytes []string, idx int) []string {
	if len(bytes) == 1 {
		i, err := strconv.ParseInt(bytes[0], 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("oxi: bytes: %v | num: %d\n", bytes[0], i)
		return bytes
	}
	var (
		ones  = make([]int, 0, len(bytes))
		zeros = make([]int, 0, len(bytes))
	)
	for i, b := range bytes {
		if b[idx] == '1' {
			ones = append(ones, i)
		} else {
			zeros = append(zeros, i)
		}
	}
	if len(ones) >= len(zeros) {
		bytes = cleanSlice(bytes, zeros)
	} else {
		bytes = cleanSlice(bytes, ones)
	}
	return oxigenGeneratorRating(bytes, idx+1)
}

func scrubberGeneratorRating(bytes []string, idx int) []string {
	if len(bytes) == 1 {
		i, err := strconv.ParseInt(bytes[0], 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("scr: bytes: %v | num: %d\n", bytes[0], i)
		return bytes
	}
	var (
		ones  = make([]int, 0, len(bytes))
		zeros = make([]int, 0, len(bytes))
	)
	for i, b := range bytes {
		if b[idx] == '1' {
			ones = append(ones, i)
		} else {
			zeros = append(zeros, i)
		}
	}
	if len(zeros) <= len(ones) {
		bytes = cleanSlice(bytes, ones)
	} else {
		bytes = cleanSlice(bytes, zeros)
	}
	return scrubberGeneratorRating(bytes, idx+1)
}

func cleanSlice(bytes []string, idxs []int) []string {
	for i := len(idxs) - 1; i >= 0; i-- {
		bytes = append(bytes[:idxs[i]], bytes[idxs[i]+1:]...)
	}
	return bytes
}

func consumption(in []string) {
	ones := make([]int, len(in[0]))
	for _, i := range in {
		for idx := range i {
			countOne(i[idx], &(ones[idx]))
		}
	}
	g := gamma(ones, len(in))
	e := epsilon(ones, len(in))
	fmt.Printf("e: %d | g: %d | calc: %d\n", e, g, e*g)
}

func countOne(in byte, i *int) {
	if in == '1' {
		*i++
	}
}

func gamma(in []int, total int) int64 {
	var binary string
	for _, i := range in {
		if i > total/2 {
			binary += "1"
		} else {
			binary += "0"
		}
	}
	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func epsilon(in []int, total int) int64 {
	var binary string
	for _, i := range in {
		if i < total/2 {
			binary += "1"
		} else {
			binary += "0"
		}
	}
	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func scan(in []byte) (inn []string) {
	s := bufio.NewScanner(bytes.NewReader(in))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		inn = append(inn, s.Text())
	}
	return inn
}
