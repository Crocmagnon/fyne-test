package main

import (
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestUI(t *testing.T) {
	in, out := makeUI()
	got := out.Text
	want := "Hello, world!"
	if got != want {
		t.Errorf("Incorrect initial greeting, got %q want %q", got, want)
	}

	test.Type(in, "Gabriel")
	got = out.Text
	want = "Hello, Gabriel!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
