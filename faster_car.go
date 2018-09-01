package race

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type FasterCar struct {
	points []Point
}

const (
	delim    = ` - `
	delimLen = len(delim)
)

func NewFasterCar(m string) *FasterCar {
	car := &FasterCar{}

	data, _ := ioutil.ReadFile(m)
	lines := strings.Split(string(data), "\n")

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
			car.points = append(car.points, Point{
				Left:  line[:pos1],
				Edge:  int(n),
				Right: line[pos2+delimLen:],
			})
		}
	}

	return car
}

func (c *FasterCar) Go(start, finish string) []string {
	cities := []string{}
	value := 0
	for _, p := range c.points {
		if p.Left == start {
			next := min(c.points, start)
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
