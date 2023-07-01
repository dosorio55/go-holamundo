package interfaces

type Alien struct {
	// this is how struct inherit from other interfaces
	Human
	SuperSmart bool
}