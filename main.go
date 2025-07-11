package main

import (
	"github.com/MRaihanZ/learn-japanese/quiz"
	"github.com/MRaihanZ/learn-japanese/randomize"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func menu(reader *bufio.Reader) string {
	var inputUser string

	fmt.Println("==============================")
	fmt.Println("1. quiz Hiragana to romaji")
	fmt.Println("2. quiz Romaji to Hiragana")
	fmt.Println("3. Randomize Hiragana")
	fmt.Println("4. Randomize Romaji")
	fmt.Println("5. Exit")
	fmt.Println("==============================")
	fmt.Print("Input: ")
	inputUser, _ = reader.ReadString('\n')
	inputUser = strings.TrimSpace(inputUser)
	return inputUser
}

func main() {
	var kana, romaji []string
	reader := bufio.NewReader(os.Stdin)
	kana, romaji = openJson("E:/GIU/Devel/go_app/Simple-Random/db/qna.json")
	shuffled := make([]string, len(romaji))
	for {
		inputUser := menu(reader)
		switch inputUser {
		case "1":
			quiz.HiraganaQuizRomaji(kana, romaji)
		case "2":
			fmt.Println("TBA")
		case "3":
			copy(shuffled, kana)
			fmt.Println()
			fmt.Println("Hiragana: ", randomize.SingleRandomize(shuffled))
			fmt.Println()
		case "4":
			copy(shuffled, romaji)
			fmt.Println()
			fmt.Println("Romaji: ", randomize.SingleRandomize(shuffled))
			fmt.Println()
		case "5":
			fmt.Println()
			color.Blue("MADE BY MRZ")
			fmt.Println()
			fmt.Println("Press Enter to exit, or wait 5 seconds...")

			done := make(chan struct{})

			// Reuse the same reader
			go func() {
				reader.ReadString('\n') // Wait for Enter
				done <- struct{}{}
			}()

			select {
			case <-done:
				fmt.Println("You pressed Enter. Exiting.")
				return
			case <-time.After(5 * time.Second):
				fmt.Println("Timeout. Exiting.")
				return
			}
		default:
			fmt.Println()
			color.Red("Wrong Input!!!")
			fmt.Println()
		}
	}
}
