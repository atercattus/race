package race

import (
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

type FasterCar struct {
	file string
}

func NewFasterCar(m string) *FasterCar {
	return &FasterCar{
		file: m,
	}
}

const (
	delim    = ` - `
	delimLen = len(delim)
)

var (
	pointsPool = sync.Pool{
		New: func() interface{} {
			return []Point{}
		},
	}
)

func (c *FasterCar) Go(start, finish string) []string {
	data, _ := ioutil.ReadFile(c.file)
	lines := strings.Split(string(data), "\n")

	points := pointsPool.Get().([]Point)[:0]
	defer func() { pointsPool.Put(points) }()

	for _, line := range lines {
		if len(line) == 0 {
			continue
		} else if pos1 := strings.Index(line, delim); pos1 == -1 {
			return nil
		} else if pos2 := strings.LastIndex(line, delim); pos2 == pos1 {
			return nil
		} else if n, err := strconv.Atoi(line[pos1+delimLen : pos2]); err != nil {
			return nil
		} else {
			points = append(points, Point{
				Left:  line[:pos1],
				Edge:  int(n),
				Right: line[pos2+delimLen:],
			})
		}
	}

	cities := []string{}
	value := 0
	for _, p := range points {
		if p.Left == start {
			next := min(points, start)
			cities = append(cities, next.Left)
			start = next.Right
			value += next.Edge
		}

		if p.Right == finish {
			cities = append(cities, p.Right)
			break
		}
	}

	return cities
}
