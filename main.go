package main

import (
	"ascii-art/Modify"
	"log"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		log.Fatalf("Too many arguments")
	}
	Modify.Argcheckers(arg[0])
	Modify.HashCheck(arg[0])
	arr := Modify.FromTxt()
	Modify.PrintLetter(arg[0], arr)
}
