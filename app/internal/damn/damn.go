package damn

import (
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
	if util.FlipCoin() {
		tokens = append(tokens, util.RandPick(corpus.Nouns))
	}

	// Then, compose God damn adjectives
	adjSeen := make(map[string]bool)
	addUsed := false
	for i := 0; i < level; i++ {
		adj := util.RandPick(corpus.Adjectives)
		if adjSeen[adj] {
			continue
		}
		adjSeen[adj] = true
		tokens = append(tokens, adj)

		if util.FlipCoin() && len(tokens) < level {
			conjSeen := make(map[string]bool)
			for j := 0; j < util.RandIntMinMax(1, 3); j++ {
				conj := util.RandPick(corpus.Conjunctions)
				if conjSeen[conj] {
					continue
				}
				conjSeen[conj] = true
				tokens = append(tokens, conj)
			}
		}

		// Append at random one addition if the level is high enough
		if util.FlipCoin() && !addUsed && level > 3 {
			k := util.RandMapKey(corpus.Additions)
			v := corpus.Additions[k]
			tokens = append(tokens, k, util.RandPick(v))
			addUsed = true
		}
	}

	// After that, add a single noun
	tokens = append(tokens, util.RandPick(corpus.Nouns))

	// Finally, add randomly one more adjective to the end, if not added yet
	if util.FlipCoin() {
		adj := util.RandPick(corpus.Adjectives)
		if _, ok := adjSeen[adj]; !ok {
			tokens = append(tokens, adj)
		}
	}

	return tokens
}
