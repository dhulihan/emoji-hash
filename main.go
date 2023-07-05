package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/cespare/xxhash"
	"github.com/enescakir/emoji"
	"golang.org/x/exp/maps"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatal("please provide a string")
	}

	emojiMap := emoji.Map()

	// emojis are unordered
	emojis := maps.Values(emojiMap)

	sort.Strings(emojis)

	hash, err := computeHash(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	i := hash % uint64(len(emojis))

	fmt.Println(emojis[i])
}

func computeHash(s string) (uint64, error) {
	hash := xxhash.New()

	if _, err := hash.Write([]byte(s)); err != nil {
		return 0, errors.New("failed writing hash contents")
	}

	return hash.Sum64(), nil
}
