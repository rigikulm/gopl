package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	want := "Hello Fred!"
	got := greeting("Fred")
	if want != got {
		t.Errorf("Want %q, Got %q", want, got)
	}
}
