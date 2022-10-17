package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("What is your letter? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		guess = strings.TrimSpace(guess)

		if len(guess) != 1 {
			fmt.Printf("Invalid letter size. letter=%v, len=%v", guess, len(guess))
			continue
		}
		valid = true
	}
	return guess, nil
}
