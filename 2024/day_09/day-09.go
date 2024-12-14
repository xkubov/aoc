package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type File struct {
	id   int
	size int
}

type Space struct {
	file     *File
	capacity int
	next     *Space
	prev     *Space
}

func (d *Disk) AllocateFile(s *Space, id int, size int, fragment bool) int {
	allocated := size
	if s.capacity < allocated {
		allocated = s.capacity
		if !fragment {
			return 0
		}
	}
	s.capacity -= allocated

	file_space := s
	if s.capacity > 0 {
		file_space = &Space{prev: s.prev, next: s}
		if s.prev == nil {
			d.first_space = file_space
		} else {
			s.prev.next = file_space
		}
		s.prev = file_space
	}

	file_space.file = &File{id: id, size: allocated}
	return allocated
}

func (s *Space) Deallocate(capacity int) {
	s.file.size -= capacity
	s.capacity += capacity
	if s.file.size == 0 {
		s.file = nil
	}
}

type Disk struct {
	first_space *Space
	last_space  *Space
}

func (d *Disk) Add(space *Space) {
	if d.first_space == nil {
		d.first_space = space
		d.last_space = space
	} else {
		d.last_space.next = space
		space.prev = d.last_space
		d.last_space = space
	}
}

func CreateDisk(coded_disk string) Disk {
	var disk Disk
	id := 0
	for i, v := range coded_disk {
		space := Space{}
		if i%2 == 0 {
			space.file = &File{id: id, size: int(v - '0')}
			space.capacity = 0
			id++
		} else {
			space.capacity = int(v - '0')
		}
		disk.Add(&space)
	}
	return disk
}

func (d *Disk) GetPrevFileSpace(space *Space) *Space {
	for space != nil && space.file == nil {
		space = space.prev
	}
	return space
}

func (d *Disk) GetNextFreeSpace(space *Space, stop *Space) *Space {
	for space != nil && space.capacity == 0 {
		if space == stop {
			return nil
		}
		space = space.next
	}
	return space
}

func (d *Disk) ReasignFreeSpace(fragment bool) {
	file_space := d.GetPrevFileSpace(d.last_space)
	for ; file_space != nil; file_space = d.GetPrevFileSpace(file_space.prev) {
		free_space := d.GetNextFreeSpace(d.first_space, file_space)
		for ; free_space != nil; free_space = d.GetNextFreeSpace(free_space.next, file_space) {
			allocated := d.AllocateFile(free_space, file_space.file.id, file_space.file.size, fragment)
			file_space.Deallocate(allocated)
			if file_space.file == nil {
				break
			}
		}
	}
}

func (d *Disk) Checksum() (sum int) {
	i := 0
	for space := d.first_space; space != nil; space = space.next {
		if space.file == nil {
			i += space.capacity
			continue
		}
		for j := 0; j < space.file.size; j++ {
			sum += space.file.id * i
			i += 1
		}
	}
	return
}

func (d *Disk) String() string {
	result := ""
	for space := d.first_space; space != nil; space = space.next {
		if space.file != nil {
			for i := 0; i < space.file.size; i++ {
				result += string('0' + space.file.id)
			}
		}
		for i := 0; i < space.capacity; i++ {
			result += "."
		}
	}
	return result
}

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)

	var line string

	for scanner.Scan() {
		line += strings.TrimSpace(scanner.Text())
	}

	disk := CreateDisk(line)
	disk_2 := CreateDisk(line)

	disk.ReasignFreeSpace(true)

	fmt.Println("Puzzle 1:", disk.Checksum())

	disk_2.ReasignFreeSpace(false)

	fmt.Println("Puzzle 2:", disk_2.Checksum())

	f.Close()
}
