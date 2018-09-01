package race

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type FasterCar struct {
	file string
}

func NewFasterCar(m string) *FasterCar {
	return &FasterCar{
		file: m,
	}
}

func (c *FasterCar) Go(start, finish string) []string {
	re := regexp.MustCompile(`(?P<Left>[A-Z]) - (?P<Edge>[0-9]) - (?P<Right>[A-Z])`)

	data, _ := ioutil.ReadFile(c.file)

	l := len(strings.Split(string(data), "\n")) - 1

	all := []map[string]string{}
	for i := 0; i < l; i++ {
		data, err := ioutil.ReadFile(c.file)
		if err != nil {
			return []string{}
		}

		line := strings.Split(string(data), "\n")[i]

		match := re.FindStringSubmatch(line)
		params := map[string]string{}
		for i, name := range re.SubexpNames() {
			if i > 0 && i <= len(match) {
				params[name] = match[i]
			}
		}

		all = append(all, params)
	}

	points := []Point{}
	for _, p := range all {
		edge, _ := strconv.Atoi(p["Edge"])
		points = append(points, Point{
			Left:  p["Left"],
			Edge:  edge,
			Right: p["Right"],
		})
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
