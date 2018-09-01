package race

import (
	"github.com/RyanCarrier/dijkstra"
	"io/ioutil"
	"strconv"
	"strings"
)

type FasterCar struct {
	citiesIdx  map[string]int32
	citiesName []string

	graph *dijkstra.Graph
}

const (
	delim    = ` - `
	delimLen = len(delim)
)

func NewFasterCar(m string) *FasterCar {
	car := &FasterCar{
		citiesIdx: map[string]int32{},
		graph:     dijkstra.NewGraph(),
	}

	data, _ := ioutil.ReadFile(m)
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
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

			car.graph.AddArc(int(idxLeft), int(idxRight), int64(n))
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

		c.graph.AddVertex(int(idx))
	}
	return idx
}

func (c *FasterCar) Go(start, finish string) []string {
	startIdx, ok := c.citiesIdx[start]
	if !ok {
		return nil
	}
	finishIdx, ok := c.citiesIdx[finish]
	if !ok {
		return nil
	}

	best, err := c.graph.Shortest(int(startIdx), int(finishIdx))
	if err != nil {
		return nil
	}

	cities := make([]string, len(best.Path))
	for pos, idx := range best.Path {
		cities[pos] = c.citiesName[idx]
	}
	return cities
}
