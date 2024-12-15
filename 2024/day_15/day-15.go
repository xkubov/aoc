package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WarehouseObject rune

const (
	ROBOT       WarehouseObject = '@'
	BOX         WarehouseObject = 'O'
	WALL        WarehouseObject = '#'
	EMPTY_SPACE WarehouseObject = '.'
	WIDE_BOX_L  WarehouseObject = '['
	WIDE_BOX_R  WarehouseObject = ']'
)

type Direction rune

const (
	TOP    Direction = '^'
	LEFT   Direction = '<'
	RIGHT  Direction = '>'
	BOTTOM Direction = 'v'
)

type Position struct {
	x int
	y int
}

func (p Position) GetNeighbor(d Direction) Position {
	switch d {
	case TOP:
		return Position{x: p.x - 1, y: p.y}
	case BOTTOM:
		return Position{x: p.x + 1, y: p.y}
	case LEFT:
		return Position{x: p.x, y: p.y - 1}
	default: // RIGHT
		return Position{x: p.x, y: p.y + 1}
	}
}

type Tile struct {
	obj WarehouseObject
}

type Warehouse [][]Tile

func (w Warehouse) Print() {
	for _, tiles := range w {
		for _, tile := range tiles {
			fmt.Print(string(tile.obj))
		}
		fmt.Println()
	}
}

func (w *Warehouse) GetRobotPos() *Position {
	for x, tiles := range *w {
		for y, tile := range tiles {
			if tile.obj == ROBOT {
				return &Position{x: x, y: y}
			}
		}
	}
	return nil
}

func (w *Warehouse) GetTile(p Position) *Tile {
	return &(*w)[p.x][p.y]
}

func (w *Warehouse) MoveObject(p Position, d Direction, move bool) bool {
	dest := p.GetNeighbor(d)
	switch w.GetTile(dest).obj {
	case EMPTY_SPACE:
		if move {
			w.GetTile(dest).obj = w.GetTile(p).obj
			w.GetTile(p).obj = EMPTY_SPACE
		}
		return true
	case BOX:
		if w.MoveObject(dest, d, move) {
			if move {
				w.GetTile(dest).obj = w.GetTile(p).obj
				w.GetTile(p).obj = EMPTY_SPACE
			}
			return true
		}
	case WIDE_BOX_L:
		left_box := dest
		right_box := left_box.GetNeighbor(RIGHT)

		w.GetTile(left_box).obj = BOX
		w.GetTile(right_box).obj = BOX

		can_move := w.MoveObject(right_box, d, false) && w.MoveObject(left_box, d, false)

		w.GetTile(left_box).obj = WIDE_BOX_L
		w.GetTile(right_box).obj = WIDE_BOX_R

		if move && can_move {
			w.MoveObject(right_box, d, true)
			w.MoveObject(left_box, d, true)
			w.GetTile(left_box).obj = w.GetTile(p).obj
			w.GetTile(p).obj = EMPTY_SPACE
		}
		return can_move

	case WIDE_BOX_R:
		right_box := dest
		left_box := right_box.GetNeighbor(LEFT)

		w.GetTile(left_box).obj = BOX
		w.GetTile(right_box).obj = BOX

		can_move := w.MoveObject(right_box, d, false) && w.MoveObject(left_box, d, false)

		w.GetTile(left_box).obj = WIDE_BOX_L
		w.GetTile(right_box).obj = WIDE_BOX_R

		if move && can_move {
			w.MoveObject(left_box, d, true)
			w.MoveObject(right_box, d, true)
			w.GetTile(right_box).obj = w.GetTile(p).obj
			w.GetTile(p).obj = EMPTY_SPACE
		}
		return can_move
	}
	return false
}

func (old Warehouse) WideCopy() Warehouse {
	var w Warehouse
	for _, old_tiles := range old {
		var new_tiles []Tile
		for _, tile := range old_tiles {
			switch tile.obj {
			case ROBOT:
				new_tiles = append(new_tiles, Tile{obj: ROBOT}, Tile{obj: EMPTY_SPACE})
			case BOX:
				new_tiles = append(new_tiles, Tile{obj: WIDE_BOX_L}, Tile{obj: WIDE_BOX_R})
			case WALL:
				new_tiles = append(new_tiles, Tile{obj: WALL}, Tile{obj: WALL})
			case EMPTY_SPACE:
				new_tiles = append(new_tiles, Tile{obj: EMPTY_SPACE}, Tile{obj: EMPTY_SPACE})
			}
		}
		w = append(w, new_tiles)
	}
	return w
}

func (w *Warehouse) SumBoxGPS() int {
	sum := 0
	for x, tiles := range *w {
		for y, tile := range tiles {
			if tile.obj == BOX || tile.obj == WIDE_BOX_L {
				sum += 100*x + y
			}
		}
	}
	return sum
}

func (w *Warehouse) MoveRobot(d Direction) {
	r_pos := w.GetRobotPos()
	w.MoveObject(*r_pos, d, true)
}

func ParseTiles(line string) []Tile {
	var tiles []Tile
	for _, obj := range strings.TrimSpace(line) {
		tiles = append(tiles, Tile{obj: WarehouseObject(obj)})
	}
	return tiles
}

func ParseMovements(line string) []Direction {
	var movements []Direction
	for _, m := range strings.TrimSpace(line) {
		movements = append(movements, Direction(m))
	}
	return movements
}

func main() {
	f, _ := os.Open(os.Args[1])

	scanner := bufio.NewScanner(f)

	var w Warehouse

	var rm []Direction

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		w = append(w, ParseTiles(line))
	}

	w_wide := w.WideCopy()

	for scanner.Scan() {
		rm = append(rm, ParseMovements(scanner.Text())...)
	}

	for _, m := range rm {
		w.MoveRobot(m)
	}
	fmt.Println("Puzzle 1:", w.SumBoxGPS())

	for _, m := range rm {
		w_wide.MoveRobot(m)
	}

	fmt.Println("Puzzle 2:", w_wide.SumBoxGPS())

	f.Close()
}
