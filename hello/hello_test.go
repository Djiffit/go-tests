package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to a specific person", func(t *testing.T) {
		got := Hello("Pekka", "")
		want := "Hello, Pekka"
		assertMessage(t, got, want)
	})

	t.Run("saying hello to nobody is 'Hello world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Monde", "French")
		want := "Bonjour, Monde"
		assertMessage(t, got, want)
	})
}
