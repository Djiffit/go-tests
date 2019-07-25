package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Pekka")

	want := "Hello, Pekka!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
