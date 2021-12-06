package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
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
	part1(scan(input))
	part2(scan(input))
}

func part1(vents []*vent) {
	var points []point
	for _, v := range vents {
		points = append(points, v.horizontal()...)
		points = append(points, v.vertical()...)
	}
	fmt.Println("1: points: ", countDangerousPoints(points))
}

func part2(vents []*vent) {
	var points []point
	for _, v := range vents {
		points = append(points, v.horizontal()...)
		points = append(points, v.vertical()...)
		points = append(points, v.diagonal()...)
	}
	fmt.Println("2: points: ", countDangerousPoints(points))
}

func countDangerousPoints(points []point) int {
	dangerousPoints := make(map[point]int)
	for _, point := range points {
		counter := dangerousPoints[point]
		dangerousPoints[point] = counter + 1
	}
	var danger int
	for _, count := range dangerousPoints {
		if count >= 2 {
			danger++
		}
	}
	return danger
}

type vent struct {
	from, to point
}

func (v *vent) horizontal() (points []point) {
	if v.from.x != v.to.x {
		return nil
	}
	for i := v.from.y; i <= v.to.y; i++ {
		points = append(points, point{x: v.from.x, y: i})
	}
	for i := v.to.y; i <= v.from.y; i++ {
		points = append(points, point{x: v.from.x, y: i})
	}
	return points
}

func (v *vent) vertical() (points []point) {
	if v.from.y != v.to.y {
		return nil
	}
	for i := v.from.x; i <= v.to.x; i++ {
		points = append(points, point{y: v.from.y, x: i})
	}
	for i := v.to.x; i <= v.from.x; i++ {
		points = append(points, point{y: v.from.y, x: i})
	}
	return points
}

func (v *vent) diagonal() (points []point) {
	xDiff := v.from.x - v.to.x
	yDiff := v.from.y - v.to.y
	if math.Abs(float64(xDiff)) != math.Abs(float64(yDiff)) {
		return nil
	}

	for i := 0; float64(i) <= math.Abs(float64(xDiff)); i++ {
		var p point
		if xDiff > 0 {
			p.x = v.from.x - i
		} else {
			p.x = v.from.x + i
		}
		if yDiff > 0 {
			p.y = v.from.y - i
		} else {
			p.y = v.from.y + i
		}

		points = append(points, p)
	}

	return points
}

type point struct {
	x, y int
}

func scan(input []byte) (vents []*vent) {
	s := bufio.NewScanner(bytes.NewReader(input))
	s.Split(bufio.ScanLines)

	for s.Scan() {
		vents = append(vents, scanVents(s.Text()))
	}
	return vents
}

func scanVents(line string) *vent {
	points := strings.Split(line, " -> ")
	return &vent{
		from: scanPoint(points[0]),
		to:   scanPoint(points[1]),
	}
}

func scanPoint(text string) point {
	coords := strToNums(strings.Split(text, ","))
	return point{x: coords[0], y: coords[1]}
}

func strToNums(nums []string) []int {
	numbers := make([]int, len(nums))
	for i, n := range nums {
		numbers[i], _ = strconv.Atoi(n)
	}
	return numbers
}
