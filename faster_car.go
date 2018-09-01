package race

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type FasterCar struct {
	points     []PointInt
	citiesIdx  map[string]int32
	citiesName []string
}

const (
	delim    = ` - `
	delimLen = len(delim)
)

func NewFasterCar(m string) *FasterCar {
	car := &FasterCar{
		citiesIdx: map[string]int32{},
	}

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
			nameLeft := line[:pos1]
			nameRight := line[pos2+delimLen:]

			idxLeft := car.addCity(nameLeft)
			idxRight := car.addCity(nameRight)

			car.points = append(car.points, PointInt{
				Left:  idxLeft,
				Edge:  int32(n),
				Right: idxRight,
			})
		}
	}

	return car
}

func (c *FasterCar) addCity(name string) (idx int32) {
	idx, ok := c.citiesIdx[name]
	if !ok {
		idx = int32(len(c.citiesName))
		c.citiesIdx[name] = idx
		c.citiesName = append(c.citiesName, name)
	}
	return idx
}

func (c *FasterCar) Go(start, finish string) []string {
	startIdx := c.addCity(start)
	finishIdx := c.addCity(finish)

	cities := []string{}
	value := int32(0)
	for _, p := range c.points {
		if p.Left == startIdx {
			next := minInt(c.points, startIdx)
			cities = append(cities, c.citiesName[next.Left])
			startIdx = next.Right
			value += next.Edge
		}

		if p.Right == finishIdx {
			cities = append(cities, c.citiesName[p.Right])
			break
		}
	}

	return cities
}
