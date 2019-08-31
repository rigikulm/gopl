package main

import "testing"

func TestEchoArgs(t *testing.T) {

	t.Run("Multiple - arguments", func(t *testing.T) {
		a := []string{"-arg1", "-arg2", "-arg3"}
		want := "-arg1 -arg2 -arg3"
		got := echoArgs(a)
		if want != got {
			t.Errorf("want: %q, got %q", want, got)
		}
	})

	t.Run("Mixture of hyphen arguments and regular parameters", func(t *testing.T) {
		a := []string{"-arg1", "-arg2", "filename"}
		want := "-arg1 -arg2 filename"
		got := echoArgs(a)
		if want != got {
			t.Errorf("want: %q, got %q", want, got)
		}
	})

	t.Run("No arguments", func(t *testing.T) {
		a := []string{}
		want := ""
		got := echoArgs(a)
		if want != got {
			t.Errorf("want: %q, got %q", want, got)
		}
	})

	a := []string{"-arg1", "-arg2", "argument"}
	want := "-arg1 -arg2 argument"
	got := echoArgs(a)
	if want != got {
		t.Errorf("want: %q, got %q", want, got)
	}
}
