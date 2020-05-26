package main

import "testing"

func TestSwar(t *testing.T) {
	testdata := map[uint32]uint8{
		1:  1,
		11: 3,
		53: 4,
	}

	for i, num := range testdata {
		if r := Swar(i); r != num {
			t.Errorf("fail %d has %d, but result is %d", i, num, r)
		}
	}
}

func BenchmarkSwar(b *testing.B) {
	for i := b.N; i > 0; i-- {
		Swar(uint32(i))
	}
}
