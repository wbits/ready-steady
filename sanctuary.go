package main

type Animal struct {
	name string
	likes string
	moves string
}

func (a Animal) String() string {
	return a.name
}

func (a Animal) move() string {
	return a.moves
}

func (a Animal) eat() string {
	return a.likes
}

func main() {
	animals := []Animal{
		Animal{"Cow", "Grass", "Slowly through the grassy fields"},
		Animal{"Dog", "Bones", "Running on the beach"},
		Animal{"Pig", "Anything", "Rolling through the mud"},
	}
}