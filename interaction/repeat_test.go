package interaction

import "testing"

func TestRepeat(t *testing.T) {
	repetitions := Repeat("a", 5)
	expect := "aaaaa"

	if repetitions != expect {
		t.Errorf("expect '%s' - result '%s'", expect, repetitions)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
