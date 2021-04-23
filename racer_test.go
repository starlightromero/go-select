package main

import (
	"testing"
)

func TestRacer(t *testing.T) {
	slowURL := "https://facebook.com"
	fastURL := "https://www.quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
