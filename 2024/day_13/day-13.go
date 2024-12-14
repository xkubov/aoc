package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vec struct {
	x int
	y int
}

type ClawMachine struct {
	button_a Vec
	button_b Vec
	prize    Vec
}

func ParseButton(line string) Vec {
	xy := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), ", ")
	x, _ := strconv.Atoi(strings.Split(xy[0], "+")[1])
	y, _ := strconv.Atoi(strings.Split(xy[1], "+")[1])

	return Vec{x: x, y: y}
}

func ParsePrice(line string) Vec {
	xy := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), ", ")
	x, _ := strconv.Atoi(strings.Split(xy[0], "=")[1])
	y, _ := strconv.Atoi(strings.Split(xy[1], "=")[1])

	return Vec{x: x, y: y}
}

func (c *ClawMachine) Cost(press_a int, press_b int) int {
	return press_a*3 + press_b
}

func (c *ClawMachine) Move(press_a int, press_b int) Vec {
	return Vec{
		x: c.button_a.x*press_a + c.button_b.x*press_b,
		y: c.button_a.y*press_a + c.button_b.y*press_b,
	}
}

func (c *ClawMachine) ComputeMinCost() int {
	tokens := 0
	diff_x := -(float64(c.button_b.x) / float64(c.button_b.y))

	a_prod := (float64(c.button_a.y)*diff_x + float64(c.button_a.x))

	a := math.Round((float64(c.prize.y)*diff_x + float64(c.prize.x)) / a_prod)
	b := math.Round((float64(c.prize.x) - a*float64(c.button_a.x)) / float64(c.button_b.x))

	if c.Move(int(a), int(b)) == c.prize {
		tokens = c.Cost(int(a), int(b))
	}
	return tokens
}

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)

	var claw_machines []ClawMachine

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		button_a := ParseButton(line)
		scanner.Scan()
		button_b := ParseButton(scanner.Text())
		scanner.Scan()
		price := ParsePrice(scanner.Text())

		claw_machines = append(claw_machines, ClawMachine{button_a, button_b, price})
	}

	tokens_needed := 0
	for _, m := range claw_machines {
		tokens_needed += m.ComputeMinCost()
	}
	fmt.Println("Puzzle 1:", tokens_needed)

	tokens_needed = 0
	for _, m := range claw_machines {
		m.prize.x += 10000000000000
		m.prize.y += 10000000000000
		tokens_needed += m.ComputeMinCost()
	}
	fmt.Println("Puzzle 2:", tokens_needed)
}
