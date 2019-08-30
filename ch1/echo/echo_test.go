package main

import "testing"

func TestEchoArgs(t *testing.T) {
	a := []string{"-arg1", "-arg2", "argument"}
	want := "-arg1 -arg2 argument"
	got := echoArgs(a)
	if want != got {
		t.Errorf("want: %q, got %q", want, got)
	}
}
