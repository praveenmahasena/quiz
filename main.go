package main

import (
	"log"

	"github.com/praveenmahasena/quiz/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		log.Fatalln(err)
	}
}
