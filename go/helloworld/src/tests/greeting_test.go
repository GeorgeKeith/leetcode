package tests

import (
	"testing"
)
import gr "greeting"

func TestGreeting(t *testing.T) {
	text := "Hello, world"
	g := gr.NewGreeting(text)
	gText := g.GetText()
	if gText != text {
		t.Errorf("g.text == %q, wanted %q", gText, text)
	}
}
