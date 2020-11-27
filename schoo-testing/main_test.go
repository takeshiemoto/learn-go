package main

import "testing"

func Echo(name string) string {
	return name
}

func TestEcho(t *testing.T) {
	want := "gorilla"
	got := Echo(want)
	if want != got {
		t.Fatalf("unexpected result. want=%q, got=%q", want, got)
	}
}
