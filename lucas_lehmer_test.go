package main

import "testing"

func TestIsMersennePrimeExponent(t *testing.T) {
	tests := []struct {
		name     string
		exponent uint64
		want     bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"two", 2, true},
		{"three", 3, true},
		{"five", 5, true},
		{"seven", 7, true},
		{"eleven", 11, false},
		{"thirteen", 13, true},
		{"seventeen", 17, true},
		{"nineteen", 19, true},
		{"twenty three", 23, false},
		{"twenty nine", 29, false},
		{"thirty one", 31, true},
		{"composite exponent", 4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMersennePrimeExponent(tt.exponent); got != tt.want {
				t.Fatalf("isMersennePrimeExponent(%d) = %v, want %v", tt.exponent, got, tt.want)
			}
		})
	}
}

func TestMersenneNumber(t *testing.T) {
	got := mersenneNumber(13).String()
	if got != "8191" {
		t.Fatalf("mersenneNumber(13) = %s, want 8191", got)
	}
}
