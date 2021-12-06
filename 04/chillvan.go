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

var (
	//go:embed test_input.txt
	example []byte
	//go:embed input.txt
	input []byte
)

func main() {
	numbers, grids := scan(input)
	winningNumber, g := play(numbers, grids)
	fmt.Printf("part 1: n: %d | unmarked: %d | res: %d\n", winningNumber, g.sum(), winningNumber*g.sum())
	numbers, grids = scan(input)
	lastNumber, g := lastBoard(numbers, grids)
	fmt.Printf("part 2: n: %d | unmarked: %d | res: %d\n", lastNumber, g.sum(), lastNumber*g.sum())
}

func play(numbers []int, grids []*grid) (int, *grid) {
	for _, g := range grids {
		if g.drawn(numbers[0]) {
			return numbers[0], g
		}
	}
	return play(numbers[1:], grids)
}

func lastBoard(numbers []int, grids []*grid) (int, *grid) {
	for i := len(grids) - 1; i >= 0; i-- {
		if grids[i].drawn(numbers[0]) {
			if len(grids) == 1 {
				return numbers[0], grids[0]
			}
			grids[i] = nil
			grids = append(grids[:i], grids[i+1:]...)
		}
	}
	return lastBoard(numbers[1:], grids)
}

type grid struct {
	grid    [5][5]*num
	current int
}

type num struct {
	n     int
	drawn bool
}

func (g *grid) drawn(num int) (won bool) {
	for x, line := range g.grid {
		for y, n := range line {
			if num == n.n {
				n.drawn = true
				return g.hasWon(x, y)
			}
		}
	}
	return false
}

func (g *grid) sum() (s int) {
	for _, line := range g.grid {
		for _, n := range line {
			if !n.drawn {
				s += n.n
			}
		}
	}
	return s
}

func (g *grid) hasWon(x, y int) bool {
	var wonX, wonY = true, true

	for _, x := range g.grid[x] {
		wonX = wonX && x.drawn
	}
	for _, line := range g.grid {
		wonY = wonY && line[y].drawn
	}

	return wonX || wonY
}

func (g *grid) append(nums [5]*num) {
	g.grid[g.current] = nums
	g.current++
}

func scan(input []byte) (numbers []int, grids []*grid) {
	s := bufio.NewScanner(bytes.NewReader(input))
	s.Split(bufio.ScanLines)

	s.Scan()
	numbers = scanNumbers(s.Text())

	for s.Scan() {
		if len(s.Text()) == 0 {
			grids = append(grids, new(grid))
			continue
		}
		grids[len(grids)-1].append(scanLine(s.Text()))
	}
	return numbers, grids
}

func scanNumbers(input string) []int {
	numbs := strings.Split(input, ",")
	numbers := make([]int, len(numbs))
	for i, n := range numbs {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("scan num: ", err)
		}
		numbers[i] = num
	}
	return numbers
}

func scanLine(input string) (numbs [5]*num) {
	line := strings.Fields(input)
	for i, number := range line {
		n, _ := strconv.Atoi(number)
		numbs[i] = &num{n: n}
	}
	return numbs
}
