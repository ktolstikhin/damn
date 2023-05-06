package lang

import "fmt"

type Vocabulary struct {
	Adjectives []string
	Nouns      []string
}

func NewVocabulary(gnd Gender, lng Language, obscene bool) (*Vocabulary, error) {
	switch lng {
	case LanguageRU:
		return vocabularyRU(gnd, obscene)
	default:
		return nil, fmt.Errorf("unknown language: %v", lng)
	}
}

func vocabularyRU(gnd Gender, obscene bool) (*Vocabulary, error) {
	var (
		nouns      []string
		adjectives []string
	)

	switch gnd {
	case GenderMasculine:
		nouns = masculineAbusiveNounsRU
		adjectives = masculineAbusiveAdjectivesRU
		if obscene {
			nouns = append(nouns, masculineObsceneNounsRU...)
			adjectives = append(adjectives, masculineObsceneAdjectivesRU...)
		}
	case GenderFeminine:
		nouns = feminineAbusiveNounsRU
		adjectives = feminineAbusiveAdjectivesRU
		if obscene {
			nouns = append(nouns, feminineObsceneNounsRU...)
			adjectives = append(adjectives, feminineObsceneAdjectivesRU...)
		}
	default:
		return nil, fmt.Errorf("unknown gender: %v", gnd)
	}

	return &Vocabulary{
		Adjectives: adjectives,
		Nouns:      nouns,
	}, nil
}
