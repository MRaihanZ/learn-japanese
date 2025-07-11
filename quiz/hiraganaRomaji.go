package quiz

import (
	"github.com/MRaihanZ/learn-japanese/randomize"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func HiraganaQuizRomaji(kana []string, romaji []string) {
	// get total lenght of the question array
	totalIdx := len(kana) - 1
	var answer, cont string
	var correct float64
	var wrong float64
	var questions, answers []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("===============================")
	fmt.Println("Hiragana To Romaji Quiz")
	fmt.Println("===============================")
	for {
		// create an random value everytime it loop
		questions, answers = randomize.Randomize(kana, romaji)

		// loop every item in array
		for i, val := range questions {
			fmt.Println()
			fmt.Println("===============================")
			color.Yellow("No." + strconv.Itoa(i+1))
			fmt.Println("> Question: ", val)
			fmt.Print("> Answer: ")
			answer, _ = reader.ReadString('\n')
			answer = strings.TrimSpace(answer)
			if answers[i] == answer {
				color.Green("✅")
				correct += 1
			} else {
				fmt.Println("> Correct Answer: ", answers[i])
				color.Red("❌")
				wrong += 1
			}
			fmt.Println("===============================")

			// check if loop index already in the end
			if i == totalIdx {
				// calculate the total score
				score := (correct / (float64(totalIdx) + 1)) * 100
				color.Green("Correct = " + strconv.FormatFloat(correct, 'f', 2, 64))
				color.Red("Wrong = " + strconv.FormatFloat(wrong, 'f', 2, 64))
				color.Blue("Score = " + strconv.FormatFloat(score, 'f', 2, 64))
				fmt.Print("Do you want to continue? (y/n): ")
				cont, _ = reader.ReadString('\n')
				cont = strings.TrimSpace(cont)

				// if user press "n"
				if cont == "n" {
					fmt.Println()
					fmt.Println("Kana:", questions)
					fmt.Println("Romaji:", answers)
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
				}

				// if not then reset the value of variable
				correct = 0
				wrong = 0
				score = 0
			}
		}
	}
}
