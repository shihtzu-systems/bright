package bingo

import (
	"log"
	"strings"
)

type Board struct {
	B Boxes
	I Boxes
	N Boxes
	G Boxes
	O Boxes
}

func (b Board) Print() {
	b.B.Print("B")
	b.I.Print("I")
	b.N.Print("N")
	b.G.Print("G")
	b.O.Print("O")
}

type Boxes []Box

func (b Boxes) Print(letter string) {
	for i, box := range b {
		log.Printf("%s-%d [ %s ]", strings.ToUpper(letter), i, box)
	}
}

type Box string
