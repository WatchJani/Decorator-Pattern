package main

import "fmt"

type optFunc func(*Setting)

type Setting struct {
	Moves int
	Time  string
}

func SetMoves(moves int) optFunc {
	return func(opts *Setting) {
		opts.Moves = moves
	}
}

func Default() Setting {
	return Setting{
		Moves: 0,
		Time:  "00:00",
	}
}

type Games struct {
	Setting
}

func New(opts ...optFunc) *Games {
	S := Default()

	for _, fn := range opts {
		fn(&S)
	}

	return &Games{
		Setting: S,
	}
}

func main() {
	myGame := New(SetMoves(2000))

	fmt.Println(myGame)
}
