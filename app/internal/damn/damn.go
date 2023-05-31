package damn

import (
	"ktolstikhin/damn/internal/damn/vocab"
)

type Damner struct {
	vocab *vocab.Vocabulary
}

func NewDamner(lang vocab.Language) *Damner {
	return &Damner{
		vocab: vocab.New(lang),
	}
}

const MaxLevel = 5

func (d *Damner) DamnYou(level int, opts ...vocab.Option) []string {
	if level < 1 || level > MaxLevel {
		level = 1
	}

	var (
		tokens []string
		corpus = d.vocab.Corpus(opts...)
	)

	// First, maybe start with a noun
	if flipCoin() {
		tokens = append(tokens, randPick(corpus.Nouns))
	}

	// Then, compose God damn adjectives
	adjSeen := make(map[string]bool)
	addUsed := false
	for i := 0; i < level; i++ {
		adj := randPick(corpus.Adjectives)
		if adjSeen[adj] {
			continue
		}
		adjSeen[adj] = true
		tokens = append(tokens, adj)

		if flipCoin() && len(tokens) < level {
			conjSeen := make(map[string]bool)
			for j := 0; j < randIntMinMax(1, 3); j++ {
				conj := randPick(corpus.Conjunctions)
				if conjSeen[conj] {
					continue
				}
				conjSeen[conj] = true
				tokens = append(tokens, conj)
			}
		}

		// Append at random one addition if the level is high enough
		if flipCoin() && !addUsed && level > 3 {
			k := randMapKey(corpus.Additions)
			v := corpus.Additions[k]
			tokens = append(tokens, k, randPick(v))
			addUsed = true
		}
	}

	// After that, add a single noun
	tokens = append(tokens, randPick(corpus.Nouns))

	// Finally, add randomly one more adjective to the end, if not added yet
	if flipCoin() {
		adj := randPick(corpus.Adjectives)
		if _, ok := adjSeen[adj]; !ok {
			tokens = append(tokens, adj)
		}
	}

	return tokens
}
