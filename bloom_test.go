package bloom

import (
	"testing"
)

func TestBloomFilterBasic(t *testing.T) {
	bf := NewBloomFilter(2, 100)

	e1 := []byte("hi")

	if bf.Check(e1) {
		t.Fatalf("empty filter reported element present")
	}

	bf.Add(e1)

	if !bf.Check(e1) {
		t.Fatalf("added element reported as absent")
	}
}
