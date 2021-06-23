package lexer

import (
	"fmt"
	//"one/token"
)

type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

func New(input string) *Lexer{
	fmt.Println("Lexer New, input ", input)
	l := &Lexer{input :input}

	return l
}

func (l *Lexer)Test(){
	fmt.Println("Lexer test")
}