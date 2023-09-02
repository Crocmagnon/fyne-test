package main

import (
	"fyne.io/fyne/v2/test"
	"testing"
	"time"
)

func TestUI(t *testing.T) {
	in, out, err := makeUI()
	got := out.Text
	want := "Hello, world!"
	if got != want {
		t.Errorf("Incorrect initial greeting, got %q want %q", got, want)
	}

	in.SetText("")
	test.Type(in, "Gabriel")
	got = out.Text
	want = "Hello, Gabriel!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if err.Text != "" {
		t.Errorf("err label should be empty, found %q", err.Text)
	}

	test.Type(in, "hello world")
	if err.Text == "" {
		t.Error("err label should not be empty")
	}
}

func TestSleep(t *testing.T) {
	sleep := time.Millisecond
	in, out, err := makeUI()

	time.Sleep(sleep) // added sleep

	got := out.Text
	want := "Hello, world!"
	if got != want {
		t.Errorf("Incorrect initial greeting, got %q want %q", got, want)
	}

	in.SetText("")
	test.Type(in, "Gabriel")

	time.Sleep(sleep) // added sleep

	got = out.Text
	want = "Hello, Gabriel!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if err.Text != "" {
		t.Errorf("err label should be empty, found %q", err.Text)
	}

	test.Type(in, "hello world")
	if err.Text == "" {
		t.Error("err label should not be empty")
	}
}
