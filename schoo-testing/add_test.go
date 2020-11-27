package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 3},
		{4, 5, 9},
		{3, 99, 102},
	}

	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("unexpected result. want=%d, got=%d", tt.want, got)
		}
	}
}
