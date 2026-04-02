package niceid

import (
	"math/rand/v2"
	"strings"
)

var Options struct{}

func ID() string {
	a1 := Adjectives[rand.IntN(len(Adjectives))]
	a2 := Adjectives[rand.IntN(len(Adjectives))]
	for a1 == a2 {
		a2 = Adjectives[rand.IntN(len(Adjectives))]
	}
	n := Nouns[rand.IntN(len(Nouns))]
	v := Verbs[rand.IntN(len(Verbs))]
	d := Adverbs[rand.IntN(len(Adverbs))]
	return strings.Join([]string{a1, a2, n, v, d}, "-")
}
