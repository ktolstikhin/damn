package damn

import (
	"math/rand"
	"strings"
	"unicode"

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

func (d *Damner) DamnYou(level int, opts ...vocab.Option) string {
	if level < 1 || level > 4 {
		level = 1
	}
	// TODO: level 1, 2, 3, 4* - add ending additions
	var (
		output   []string
		conjUsed bool
		adjSeen  = make(map[string]bool)
		corpus   = d.vocab.Corpus(opts...)
	)
	for len(adjSeen) < level {
		adj := chooseRandom(corpus.Adjectives)
		if !adjSeen[adj] {
			adjSeen[adj] = true
			output = append(output, adj)
			if flipCoin() && len(adjSeen) < level && !conjUsed {
				conjSeen := make(map[string]bool)
				for i := 0; i < rand.Intn(maxRandomConjenctions); i++ {
					conj := chooseRandom(corpus.Conjunctions)
					if !conjSeen[conj] {
						conjSeen[conj] = true
						output = append(output, conj)
					}
				}
				conjUsed = true
			}
		}
	}
	output = append(output, chooseRandom(corpus.Nouns))
	if flipCoin() {
		adj := chooseRandom(corpus.Adjectives)
		if _, ok := adjSeen[adj]; !ok {
			output = append(output, adj)
		}
	}

	return strings.Join(output, " ")
}

func ToSentence(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes) + "."
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

const maxRandomConjenctions = 3
