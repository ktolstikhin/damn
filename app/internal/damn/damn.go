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
		output  []string
		adjSeen = make(map[string]bool)
		corpus  = d.vocab.Corpus(opts...)
	)

	// First, compose God damn adjectives
	for i := 0; i < level; i++ {
		adj := util.RandPick(corpus.Adjectives)
		if adjSeen[adj] {
			continue
		}
		adjSeen[adj] = true
		output = append(output, adj)

		if util.FlipCoin() && len(output) < level {
			conjSeen := make(map[string]bool)
			for j := 0; j < util.RandIntMinMax(1, 3); j++ {
				conj := util.RandPick(corpus.Conjunctions)
				if conjSeen[conj] {
					continue
				}
				conjSeen[conj] = true
				output = append(output, conj)
			}
		}
	}

	// Then, add a single noun
	output = append(output, util.RandPick(corpus.Nouns))

	// After that, add randomly one more adjective to the end, if not added yet
	if util.FlipCoin() {
		adj := util.RandPick(corpus.Adjectives)
		if _, ok := adjSeen[adj]; !ok {
			output = append(output, adj)
		}
	}

	// Finally, append at random one addition if the level is high enough
	if util.FlipCoin() && level > 3 {
		if d.vocab.Lang == vocab.LanguageRU {
			output = append(output, ",")
		}
		k := util.RandMapKey(corpus.Additions)
		v := corpus.Additions[k]
		output = append(output, k, util.RandPick(v))
	}

	return output
}
