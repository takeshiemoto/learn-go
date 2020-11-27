package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"1+2", 1, 2, 3},
		{"2+2", 2, 2, 4},
		{"100+100", 100, 100, 200},
	}

	for _, tt := range tests {
		// サブテスト
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("unexpected result. want=%d, got=%d", tt.want, got)
			}
		})
	}
}
