package race

import (
	"testing"
)

func TestCar_Go(t *testing.T) {
	car := NewCar("./routes")
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

func BenchmarkCar_Go(b *testing.B) {
	car := NewCar("./routes")
	for i := 0; i < b.N; i++ {
		car.Go("A", "D")
	}
}
