package main

import (
	"flag"
	"fmt"
	"os"

	"ktolstikhin/damn/internal/damn"
	"ktolstikhin/damn/internal/damn/vocab"
)

func main() {
	var (
		genderStr string
		langStr   string
		level     int
		obscene   bool
	)
	flag.StringVar(&genderStr, "gender", "m", "God damn gender: m - male, f - female.")
	flag.StringVar(&langStr, "language", "ru", "God damn language: ru.")
	flag.IntVar(&level, "level", 1, "God damn level: from 1 to sky is the limit.")
	flag.BoolVar(&obscene, "obscene", false, "Usage of obscene vocabulary.")
	flag.Parse()

	gender, ok := genders[genderStr]
	if !ok {
		fmt.Printf("unknown gender: %s\n", genderStr)
		os.Exit(1)
	}

	lng, ok := languages[langStr]
	if !ok {
		fmt.Printf("unknown language: %s\n", langStr)
		os.Exit(1)
	}

	tokens := damn.NewDamner(lng).DamnYou(
		level,
		vocab.WithGender(gender),
		vocab.WithObscene(obscene),
	)

	fmt.Println(compose(tokens))
}

var (
	genders = map[string]vocab.Gender{
		"m": vocab.GenderMasculine,
		"f": vocab.GenderFeminine,
	}
	languages = map[string]vocab.Language{
		"ru": vocab.LanguageRU,
	}
)

func compose(tokens []string) string {
	var s string
	for i, t := range tokens {
		if i == 0 || t == "," {
			s += t
		} else {
			s += " " + t
		}
	}

	return s
}
