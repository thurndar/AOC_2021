package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var (
	//go:embed test_input.txt
	example []byte
	//go:embed input.txt
	input []byte

	cache = make(map[uint64]uint64, 256)
)

func main() {
	part1(scan(input))
	part2(scan(input))
}

type latternfishes []*latternfish

func (f *latternfishes) String() string {
	days := make([]string, len(*f))
	for i, fish := range *f {
		days[i] = strconv.Itoa(fish.timer)
	}
	return strings.Join(days, ",")
}

type latternfish struct {
	timer int
}

func (f *latternfish) nextDay() *latternfish {
	if f.timer == 0 {
		f.timer = 6
		return &latternfish{timer: 8}
	}
	f.timer--
	return nil
}

func (f *latternfish) simulate(day, remaining uint64) (sum uint64) {
	sum = 1
	for ; day < remaining; day++ {
		if fish := f.nextDay(); fish != nil {
			if c, ok := cache[remaining-(day+1)]; ok {
				sum += c
				continue
			}
			s := fish.simulate(day+1, remaining)
			cache[remaining-(day+1)] = s
			sum += s
		}
	}
	return sum
}

func part1(fishs *latternfishes) {
	// fmt.Printf("Initial state: %s\n", fishs)
	for i := 0; i < 80; i++ {
		for _, fish := range *fishs {
			if f := fish.nextDay(); f != nil {
				*fishs = append(*fishs, f)
			}
		}
		// fmt.Printf("After %2d days: %s\n", i, fishs)
	}
	fmt.Println("fish count: ", len(*fishs))
}

func part2(fishs *latternfishes) {
	sum := uint64(0)
	counter := map[latternfish]uint64{}
	for _, f := range *fishs {
		c := counter[*f]
		counter[*f] = c + 1
	}
	for f, c := range counter {
		s := (&f).simulate(0, 256)
		sum += s * c
	}
	fmt.Println("fish count: ", sum)
}

func scan(input []byte) (fishs *latternfishes) {
	s := bufio.NewScanner(bytes.NewReader(input))
	s.Split(bufio.ScanLines)
	fishs = new(latternfishes)

	for s.Scan() {
		days := strToNums(strings.Split(s.Text(), ","))
		for _, day := range days {
			*fishs = append(*fishs, &latternfish{timer: day})
		}
	}
	return fishs
}

func strToNums(nums []string) []int {
	numbers := make([]int, len(nums))
	for i, n := range nums {
		numbers[i], _ = strconv.Atoi(n)
	}
	return numbers
}
