package main

import "testing"

func TestHello(t *testing.T)  {
	got := Hello("Yaser")
	want := "Hello, Yaser"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	
}