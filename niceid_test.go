package niceid

import (
	"fmt"
	"testing"
)

func TestPrintID(t *testing.T) {
	fmt.Println(ID())
}

func TestNumberOfIDs(t *testing.T) {
	fmt.Println(len(Adjectives) * len(Adjectives) * len(Nouns) * len(Verbs) * len(Adverbs))
}
