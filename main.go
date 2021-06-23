package main

import (
	"fmt"
	"os"
	"one/repl"
)

func main(){
	fmt.Println("Hello! This is One programming language!")
	fmt.Println("Type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
