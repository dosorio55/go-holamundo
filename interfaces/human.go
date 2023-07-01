package interfaces

type Human interface {
	Breathe()
	Think()
	Eat()
	Speak()
	Alive() bool
}

func MyHuman (h Human) Human {
	h.Breathe()
	h.Think()
	h.Eat()
	h.Speak()
	h.Alive()

	return h
}