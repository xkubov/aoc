package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func (p Position) String() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

type TopMap [][]int

func (m TopMap) GetTrailheads() []Position {
	var trailheads []Position
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				trailheads = append(trailheads, Position{x: i, y: j})
			}
		}
	}

	return trailheads
}

func (m TopMap) Print() {
	for _, heights := range m {
		for _, h := range heights {
			fmt.Printf("%d ", h)
		}
		fmt.Println()
	}
}

func (m TopMap) IsOutside(pos Position) bool {
	return pos.x < 0 || pos.y < 0 || pos.x >= len(m) || pos.y >= len(m[0])
}

func (m TopMap) MovesFrom(pos Position) []Position {
	val := m[pos.x][pos.y]

	var positions []Position

	top_move := Position{x: pos.x - 1, y: pos.y}
	if !m.IsOutside(top_move) && m[top_move.x][top_move.y] == val+1 {
		positions = append(positions, top_move)
	}

	down_move := Position{x: pos.x + 1, y: pos.y}
	if !m.IsOutside(down_move) && m[down_move.x][down_move.y] == val+1 {
		positions = append(positions, down_move)
	}

	right_move := Position{x: pos.x, y: pos.y + 1}
	if !m.IsOutside(right_move) && m[right_move.x][right_move.y] == val+1 {
		positions = append(positions, right_move)
	}

	left_move := Position{x: pos.x, y: pos.y - 1}
	if !m.IsOutside(left_move) && m[left_move.x][left_move.y] == val+1 {
		positions = append(positions, left_move)
	}

	return positions
}

func (m TopMap) PosReached(start Position) []Position {
	var reaches []Position

	for _, pos := range m.MovesFrom(start) {
		if m[pos.x][pos.y] == 9 {
			reaches = append(reaches, pos)
		} else {
			s := m.PosReached(pos)
			reaches = append(reaches, s...)
		}
	}

	return reaches
}

func CountUnique(positions []Position) int {
	seen := make(map[string]bool)
	sum := 0
	for _, p := range positions {
		h := p.String()
		if _, ok := seen[h]; !ok {
			sum++
			seen[h] = true
		}
	}
	return sum
}

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)

	var topmap TopMap

	for scanner.Scan() {
		var heights []int
		line := scanner.Text()
		for _, height := range strings.TrimSpace(line) {
			heights = append(heights, int(height-'0'))
		}
		topmap = append(topmap, heights)
	}

	sum := 0
	for _, t := range topmap.GetTrailheads() {
		score := CountUnique(topmap.PosReached(t))
		sum += score
	}

	fmt.Println("Puzzle 1:", sum)

	sum = 0
	for _, t := range topmap.GetTrailheads() {
		sum += len(topmap.PosReached(t))
	}

	fmt.Println("Puzzle 2:", sum)

	f.Close()
}
