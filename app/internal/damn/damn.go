package damn

import (
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
		adj := util.RandStr(corpus.Adjectives)
		if adjSeen[adj] {
			continue
		}
		adjSeen[adj] = true
		output = append(output, adj)

		if util.FlipCoin() && len(output) < level {
			conjSeen := make(map[string]bool)
			for j := 0; j < util.RandMinMaxInt(1, 3); j++ {
				conj := util.RandStr(corpus.Conjunctions)
				if conjSeen[conj] {
					continue
				}
				conjSeen[conj] = true
				output = append(output, conj)
			}
		}
	}

	// Then, add a single noun
	output = append(output, util.RandStr(corpus.Nouns))

	// After that, add randomly one more adjective to the end, if not added yet
	if util.FlipCoin() {
		adj := util.RandStr(corpus.Adjectives)
		if _, ok := adjSeen[adj]; !ok {
			output = append(output, adj)
		}
	}

	// Finally, append at random one addition if the level is high enough
	if util.FlipCoin() && level > 3 {
		k, v := util.RandomKeyValueFromMap(corpus.Additions)
		output = append(output, util.RandStr(corpus.Conjunctions), k, util.RandStr(v))
	}

	return strings.Join(output, " ")
}
