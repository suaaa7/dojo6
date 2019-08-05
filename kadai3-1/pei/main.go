package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo6/kadai3-1/pei/pkg/typing"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

func main() {
	os.Exit(execute())
}

func execute() int {
	fmt.Println("Start Typing Game!")

	words := []string{"a", "b", "c"}

	typingCh := typing.Question(words)
	timerCh := time.After(5 * time.Second)

	counter := 0
	correctCounter := 0

	for {
		counter++
		select {
		case isCorrect := <-typingCh:
			if isCorrect {
				correctCounter++
			}
		case <-timerCh:
			fmt.Printf("\nScore: %d/%d \n", correctCounter, counter)
			fmt.Println("End Typing Game!")
			return ExitCodeOK
		}
	}
}