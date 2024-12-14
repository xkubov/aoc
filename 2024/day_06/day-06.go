package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	TOP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func (d *Direction) TurnRight() Direction {
	return (*d + 1) % 4
}

type Position struct {
	x int
	y int
}

func (p *Position) Hash() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

func parse_position(hash string) Position {
	pos := strings.Split(hash, ",")
	pos_x, _ := strconv.Atoi(pos[0])
	pos_y, _ := strconv.Atoi(pos[1])

	return Position{
		pos_x,
		pos_y,
	}
}

func (p *Position) Move(direction Direction) *Position {
	switch direction {
	case TOP:
		return &Position{p.x - 1, p.y}
	case RIGHT:
		return &Position{p.x, p.y + 1}
	case DOWN:
		return &Position{p.x + 1, p.y}
	case LEFT:
		return &Position{p.x, p.y - 1}
	}
	return nil
}

type FloorMap []string

func (floor_map *FloorMap) GetStartingPosition() *Position {
	for posx, line := range *floor_map {
		for posy, v := range line {
			if v == rune('^') {
				return &Position{posx, posy}
			}
		}
	}
	return nil
}

func (floor_map *FloorMap) IsOutside(p Position) bool {
	return p.x < 0 || p.x >= len(*floor_map) || p.y < 0 || p.y >= len((*floor_map)[0])
}

func (floor_map *FloorMap) IsObstacle(p Position) bool {
	if floor_map.IsOutside(p) {
		return false
	}

	return (*floor_map)[p.x][p.y] == '#'
}

func (floor_map *FloorMap) GetGuardMoves() (*map[string]Direction, bool) {
	direction := TOP
	visited_positions := make(map[string]Direction)

	for position := floor_map.GetStartingPosition(); !floor_map.IsOutside(*position); position = position.Move(direction) {
		if prev_direction, ok := visited_positions[position.Hash()]; ok && direction == prev_direction {
			return &visited_positions, true
		}

		visited_positions[position.Hash()] = direction
		for floor_map.IsObstacle(*position.Move(direction)) {
			direction = direction.TurnRight()
		}
	}

	return &visited_positions, false
}

func (floor_map *FloorMap) CountAddedObstacleOptions(checked_positions []Position) (sum int) {
	for _, pos := range *&checked_positions {
		if *floor_map.GetStartingPosition() == pos {
			continue
		}
		prev := (*floor_map)[pos.x]
		(*floor_map)[pos.x] = (*floor_map)[pos.x][:pos.y] + "#" + (*floor_map)[pos.x][pos.y+1:]
		if _, loop := floor_map.GetGuardMoves(); loop {
			sum++
		}
		(*floor_map)[pos.x] = prev
	}

	return
}

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	var floor_map FloorMap

	for scanner.Scan() {
		floor_map = append(floor_map, strings.TrimSpace(scanner.Text()))
	}

	visited_positions, _ := floor_map.GetGuardMoves()

	fmt.Println("Puzzle 1:", len(*visited_positions))

	var checked_positions []Position

	for pos_hash := range *visited_positions {
		checked_positions = append(checked_positions, parse_position(pos_hash))

	}

	fmt.Println("Puzzle 2:", floor_map.CountAddedObstacleOptions(checked_positions))

	file.Close()
}
