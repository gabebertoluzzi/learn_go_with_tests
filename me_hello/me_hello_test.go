package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say 'Hello World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Gabe", "")
		want := "Hello, Gabe"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Gabe", "Spanish")
		want := "Hola, Gabe"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T) {
		got := Hello("Esmé", "French")
		want := "Bonjour, Esmé"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in dutch", func(t *testing.T) {
		got := Hello("Gabe", "Dutch")
		want := "Hallo, Gabe"
		assertCorrectMessage(t, got, want)
	})
}
