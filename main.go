package main

import (
	"fmt"
	"main/scr/hangman-classic/dictionary"
	"main/scr/hangman-classic/hangman"
	"os"
)

func main() {
	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionary: %v\n", err)
		os.Exit(1)
	}

	g, err := hangman.New(8, dictionary.PickWord())
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}

	hangman.DrawWelcome()
	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l

		g.MakeAGuess(guess)
	}
}
