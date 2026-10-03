package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "hello, gobootcamp"
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
