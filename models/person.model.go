package models

type Person struct {
	Name      string
	Age       int
	IsAlive     bool
	Thinking  bool
	Eating    bool
	Speaking  bool
	Breathing bool
}

func (p *Person) Breathe() {
	p.Breathing = true
}

func (p *Person) Think() {
	p.Thinking = true
}

func (p *Person) Eat() {
	p.Eating = true
}

func (p *Person) Speak() {
	p.Speaking = true
}

func (p *Person) Alive() bool {
	return p.IsAlive
}
