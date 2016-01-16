package newmath

import "testing"

func TestSqrt(t *testing.T) {
	const in, out = 4, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}

func TestAdd(t *testing.T) {
	const ina, inb = 1, 2
	cons out = 3
	if x := Add(ina, inb); x != out {
		t.Errorf("Add(%v,%v)=%v, want %v", ina, inb, x, out)
	}
}
