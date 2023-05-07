package vocab

type Corpus struct {
	Nouns        []string
	Adjectives   []string
	Conjunctions []string
	Additions    map[string][]string
}

type Vocabulary struct {
	Lang Language
}

func New(lang Language) *Vocabulary {
	return &Vocabulary{lang}
}

type Config struct {
	gender  Gender
	obscene bool
}

type Option func(c *Config)

func WithGender(gender Gender) Option {
	return func(c *Config) {
		c.gender = gender
	}
}

func WithObscene(obscene bool) Option {
	return func(c *Config) {
		c.obscene = obscene
	}
}

func (v *Vocabulary) Corpus(opts ...Option) Corpus {
	corpus := Corpus{}
	config := makeConfig(opts)

	switch v.Lang {
	case LanguageRU:
		if config.gender == GenderMasculine {
			if !config.obscene {
				corpus.Nouns = masculineAbusiveNounsRU
				corpus.Adjectives = masculineAbusiveAdjectivesRU
				corpus.Additions = masculineAbusiveAdditionsRU
			} else {
				corpus.Nouns = masculineObsceneNounsRU
				corpus.Adjectives = masculineObsceneAdjectivesRU
				corpus.Additions = masculineObsceneAdditionsRU
			}
		} else {
			if !config.obscene {
				corpus.Nouns = feminineAbusiveNounsRU
				corpus.Adjectives = feminineAbusiveAdjectivesRU
				corpus.Additions = feminineAbusiveAdditionsRU
			} else {
				corpus.Nouns = feminineObsceneNounsRU
				corpus.Adjectives = feminineObsceneAdjectivesRU
				corpus.Additions = feminineObsceneAdditionsRU
			}
		}
		if !config.obscene {
			corpus.Conjunctions = conjunctionsAbusiveRU
		} else {
			corpus.Conjunctions = conjunctionsObsceneRU
		}
	}

	return corpus
}

func makeConfig(opts []Option) *Config {
	c := &Config{
		gender:  GenderMasculine,
		obscene: false,
	}
	for _, opt := range opts {
		opt(c)
	}

	return c
}
