package greeting

type Greeting struct {
	text string
}

func NewGreeting(t string) *Greeting {
	g := Greeting{t}
	return &g
}

func (g Greeting) GetText() string {
	return g.text
}