package damn

import (
	"math/rand"
	"strings"
	"time"

	"ktolstikhin/damn/internal/damn/lang"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func DamnYou(gnd lang.Gender, lng lang.Language, obscene bool, level int) (string, error) {
	vocab, err := lang.NewVocabulary(gnd, lng, obscene)
	if err != nil {
		return "", err
	}

	if level < 1 || level > 3 {
		level = 1
	}

	var (
		output  []string
		adjNum  = 0
		adjSeen = make(map[string]bool)
	)
	for adjNum < level {
		adj := chooseRandom(vocab.Adjectives)
		if !adjSeen[adj] {
			adjNum++
			adjSeen[adj] = true
			output = append(output, adj)
		}
	}
	output = append(output, chooseRandom(vocab.Nouns))

	return strings.Join(output, " "), nil
}

func chooseRandom(items []string) string {
	return items[rand.Intn(len(items))]
}
