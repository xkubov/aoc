package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	TOP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func (d Direction) RotateRight() Direction {
	return Direction((d + 1) % 4)
}

type Point struct {
	x int
	y int
}

func (p Point) MoveAll() map[Direction]Point {
	return map[Direction]Point{
		TOP:   {p.x - 1, p.y},
		RIGHT: {p.x, p.y + 1},
		DOWN:  {p.x + 1, p.y},
		LEFT:  {p.x, p.y - 1},
	}
}

type Region struct {
	plant     rune
	pos       Point
	neighbors map[Direction]*Region
}

type Farm [][]Region

func (farm *Farm) IsOutside(point Point) bool {
	return point.x < 0 || point.y < 0 || point.x >= len(*farm) || point.y >= len((*farm)[0])
}

func (farm *Farm) GetNeighbors(region *Region) map[Direction]*Region {
	neighbors := make(map[Direction]*Region)

	for direction, pos := range region.pos.MoveAll() {
		if !farm.IsOutside(pos) && (*farm)[pos.x][pos.y].plant == region.plant {
			neighbors[direction] = &(*farm)[pos.x][pos.y]
		}
	}

	return neighbors
}

func FilterBorder(region []*Region, direction Direction) []*Region {
	var filtered []*Region

	for _, br := range region {
		if br.neighbors[direction] == nil {
			filtered = append(filtered, br)
		}
	}

	return filtered
}

func CountDistinctEdges(edges []*Region) int {
	edge2id := make(map[*Region]int)
	var edges_len map[int]int

	for range edges {
		edges_len = make(map[int]int)
		for i, reg := range edges {
			min_id := i
			for _, nei := range reg.neighbors {
				if v, ok := edge2id[nei]; ok && v < min_id {
					min_id = v
				}
			}
			edge2id[reg] = min_id
			edges_len[min_id]++
		}
	}

	return len(edges_len)
}

func CountEdges(region []*Region) int {
	edges := 0

	for direction := TOP; direction <= LEFT; direction += 1 {
		edges += CountDistinctEdges(FilterBorder(region, direction))
	}

	return edges
}

func GetBorderSize(region []*Region) int {
	sum := 0
	for _, r := range region {
		sum += 4 - len(r.neighbors)
	}
	return sum
}

func (farm *Farm) JoinRegions() {
	for i := 0; i < len(*farm); i++ {
		for j := 0; j < len((*farm)[0]); j++ {
			reg := &(*farm)[i][j]
			reg.neighbors = farm.GetNeighbors(reg)
		}
	}
}

func (farm *Farm) GetConnectedRegion(region *Region, seen_points *map[Point]bool) []*Region {
	if _, seen := (*seen_points)[region.pos]; seen {
		return []*Region{}
	}

	nbs := []*Region{region}

	(*seen_points)[region.pos] = true

	for _, nbr := range region.neighbors {
		nbs = append(nbs, farm.GetConnectedRegion(nbr, seen_points)...)
	}

	return nbs
}

func (farm Farm) GetDistinctRegions() [][]*Region {
	farm.JoinRegions()

	var regions [][]*Region
	seen := make(map[Point]bool)

	for i := 0; i < len(farm); i++ {
		for j := 0; j < len((farm)[0]); j++ {
			region := &(farm)[i][j]
			if _, ok := seen[region.pos]; ok {
				continue
			}
			connected := farm.GetConnectedRegion(region, &seen)
			regions = append(regions, connected)
		}
	}

	return regions
}

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)

	var farm Farm

	x := 0
	for scanner.Scan() {
		var line []Region
		for y, plant := range strings.TrimSpace(scanner.Text()) {
			line = append(line, Region{plant: plant, pos: Point{x, y}})
		}
		farm = append(farm, line)
		x++
	}

	sum_puzzle_1 := 0
	sum_puzzle_2 := 0
	for _, region := range farm.GetDistinctRegions() {
		price_p1 := len(region) * GetBorderSize(region)
		price_p2 := len(region) * CountEdges(region)
		sum_puzzle_1 += price_p1
		sum_puzzle_2 += price_p2
	}

	fmt.Println("Puzzle 1:", sum_puzzle_1)
	fmt.Println("Puzzle 2:", sum_puzzle_2)
}
