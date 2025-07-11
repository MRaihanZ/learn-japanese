package randomize

import (
	"math/rand"
	"time"
)

func Randomize(hiragana []string, romaji []string) ([]string, []string) {

	// Shuffle in sync
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(hiragana), func(i, j int) {
		hiragana[i], hiragana[j] = hiragana[j], hiragana[i]
		romaji[i], romaji[j] = romaji[j], romaji[i]
	})
	return hiragana, romaji
}

func SingleRandomize(random []string) []string {

	// Shuffle in sync
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(random), func(i, j int) {
		random[i], random[j] = random[j], random[i]
	})
	return random
}
