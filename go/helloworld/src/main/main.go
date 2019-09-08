package main

import (
	"fmt"
)
import gr "../greeting"

func main() {
	text := "Hello, World"
	g := gr.NewGreeting(text)
	var _, _ = fmt.Println(g.GetText())
}
