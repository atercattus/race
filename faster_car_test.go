package race

import (
	"testing"
)

func TestFasterCar_Go(t *testing.T) {
	car := NewFasterCar("./routes")
	path := car.Go("A", "D")

	expected := []string{"A", "C", "D"}

	if len(expected) != len(path) {
		t.FailNow()
	}

	for i, city := range expected {
		if city != path[i] {
			t.Failed()
		}
	}

	t.Log(path)
}

func TestFasterCar_Go2(t *testing.T) {
	car := NewFasterCar("./routes.big.txt")
	path := car.Go("city1", "city4")

	expected := []string{"city1", "city2", "city4"}

	if len(expected) != len(path) {
		t.FailNow()
	}

	for i, city := range expected {
		if city != path[i] {
			t.Failed()
		}
	}

	t.Log(path)
}

func TestFasterCar_Go3(t *testing.T) {
	car := NewFasterCar("./routes.big.txt")
	path := car.Go("city1", "city57")

	expected := []string{"city1", "city2", "city4", "city13", "city14", "city15", "city16", "city17", "city25", "city26", "city27", "city28", "city29", "city30", "city34", "city36", "city37", "city38", "city44", "city45", "city46", "city54", "city56", "city57"}

	if len(expected) != len(path) {
		t.FailNow()
	}

	for i, city := range expected {
		if city != path[i] {
			t.Failed()
		}
	}

	t.Log(path)
}

func BenchmarkFasterCar_Go(b *testing.B) {
	car := NewFasterCar("./routes")
	for i := 0; i < b.N; i++ {
		car.Go("A", "D")
	}
}

func BenchmarkFasterCar_GoBigMap(b *testing.B) {
	car := NewFasterCar("./routes.big.txt")
	for i := 0; i < b.N; i++ {
		car.Go("city1", "city57")
	}
}
