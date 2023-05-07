package damn

import (
	"math/rand"
	"strings"

	"ktolstikhin/damn/internal/damn/vocab"
	"ktolstikhin/damn/internal/util"
)

type Damner struct {
	vocab *vocab.Vocabulary
}

func NewDamner(lang vocab.Language) *Damner {
	return &Damner{
		vocab: vocab.New(lang),
	}
}

func (d *Damner) DamnYou(level int, opts ...vocab.Option) string {
	if level < 1 {
		level = 1
	}
	var (
		output  []string
		adjSeen = make(map[string]bool)
		corpus  = d.vocab.Corpus(opts...)
	)

	// First, compose God damn adjectives
	for i := 0; i < level; i++ {
		adj := chooseRandom(corpus.Adjectives)
		if adjSeen[adj] {
			continue
		}
		adjSeen[adj] = true
		output = append(output, adj)

		if flipCoin() && len(output) < level {
			conjSeen := make(map[string]bool)
			for j := 0; j < rand.Intn(3); j++ {
				conj := chooseRandom(corpus.Conjunctions)
				if conjSeen[conj] {
					continue
				}
				conjSeen[conj] = true
				output = append(output, conj)
			}
		}
	}

	// Then, add a single noun
	output = append(output, chooseRandom(corpus.Nouns))

	// After that, add randomly one more adjective to the end, if not added yet
	if flipCoin() {
		adj := chooseRandom(corpus.Adjectives)
		if _, ok := adjSeen[adj]; !ok {
			output = append(output, adj)
		}
	}

	// Finally, append a random addition if the level is high enough
	if level > 3 && flipCoin() {
		k, v := util.RandomKeyValueFromMap(corpus.Additions)
		output = append(output, k, chooseRandom(v))
	}

	return strings.Join(output, " ")
}

func chooseRandom(items []string) string {
	if len(items) > 0 {
		return items[rand.Intn(len(items))]
	}

	return ""
}

func flipCoin() bool {
	return rand.Float32() < 0.5
}
