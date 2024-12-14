package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Vec struct {
	x int
	y int
}

type Robot struct {
	pos      Vec
	velocity Vec
}

type Map struct {
	size Vec
}

func (m Map) GetPosition(robot Robot, sec int) Vec {
	x := robot.pos.x + sec*robot.velocity.x
	y := robot.pos.y + sec*robot.velocity.y
	return Vec{
		x: (x%m.size.x + m.size.x) % m.size.x,
		y: (y%m.size.y + m.size.y) % m.size.y,
	}
}

func (m Map) IndexRobots(robots []Robot, sec int) map[Vec][]Robot {
	robots_pos := make(map[Vec][]Robot)
	for _, robot := range robots {
		pos := m.GetPosition(robot, sec)
		robots_pos[pos] = append(robots_pos[pos], robot)
	}

	return robots_pos
}

func (m Map) Print(robots []Robot, sec int) {
	robots_pos := m.IndexRobots(robots, sec)

	for y := 0; y < m.size.y; y++ {
		for x := 0; x < m.size.x; x++ {
			rs := len(robots_pos[Vec{x: x, y: y}])
			if rs == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", rs)
			}
		}
		fmt.Println()
	}
}

func (m Map) GetSafetyFactor(robots []Robot, sec int) int {
	robots_pos := m.IndexRobots(robots, sec)
	quadrants := make(map[int]int)

	qy := m.size.y / 2
	qx := m.size.x / 2

	for y := 0; y < m.size.y; y++ {
		for x := 0; x < m.size.x; x++ {
			q := x/(qx+1) + 2*(y/(qy+1))
			rs := len(robots_pos[Vec{x: x, y: y}])
			if x != qx && y != qy {
				quadrants[q] += rs
			}
		}
	}

	safety_factor := 1
	for _, count := range quadrants {
		safety_factor *= count
	}

	return safety_factor
}

func (m Map) AreRobotsGrouped(robots []Robot, sec int) bool {
	robots_pos := m.IndexRobots(robots, sec)

	for y := 0; y < m.size.y; y++ {
		n_c := 0
		for x := 0; x < m.size.x; x++ {
			if len(robots_pos[Vec{x: x, y: y}]) == 0 {
				n_c = 0
			} else {
				n_c++
				if n_c == 10 {
					return true
				}
			}
		}
	}

	return false
}

func DeserializeRobot(line string) Robot {
	r, _ := regexp.Compile("p=(\\d+),(\\d+) v=(-?\\d+),(-?\\d+)")

	res := r.FindStringSubmatch(line)

	var robot Robot

	robot.pos.x, _ = strconv.Atoi(res[1])
	robot.pos.y, _ = strconv.Atoi(res[2])
	robot.velocity.x, _ = strconv.Atoi(res[3])
	robot.velocity.y, _ = strconv.Atoi(res[4])

	return robot
}

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)

	m := Map{size: Vec{x: 101, y: 103}}
	var robots []Robot

	for scanner.Scan() {
		robot := DeserializeRobot(scanner.Text())
		robots = append(robots, robot)
	}

	fmt.Println("Puzzle 1:", m.GetSafetyFactor(robots, 100))

	fmt.Println("Puzzle 2:")
	for i := 0; i < 10000; i++ {
		if m.AreRobotsGrouped(robots, i) {
			fmt.Printf("  > %ds\n", i)
			m.Print(robots, i)
		}
	}

	f.Close()
}
