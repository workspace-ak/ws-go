package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		params := Parameters{"Chris", ""}
		got := Hello(params)
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("When empty string, say Hello, World!", func(t *testing.T) {
		params := Parameters{"", ""}
		got := Hello(params)
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Hindi", func(t *testing.T) {
		params := Parameters{"Chris", "hindi"}
		got := Hello(params)
		want := "Namaskar, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		params := Parameters{"Chris", "french"}
		got := Hello(params)
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
