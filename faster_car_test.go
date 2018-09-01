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

func BenchmarkFasterCar_Go(b *testing.B) {
	car := NewFasterCar("./routes")
	for i := 0; i < b.N; i++ {
		car.Go("A", "D")
	}
}
