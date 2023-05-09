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
	flag.StringVar(&genderStr, "gender", "male", "God damn gender: male, female.")
	flag.StringVar(&langStr, "language", "ru", "God damn language: ru.")
	flag.IntVar(&level, "level", 1, "God damn level: from 1 to sky is the limit.")
	flag.BoolVar(&obscene, "obscene", false, "Usage of obscene vocabulary.")
	flag.Parse()

	gender, err := vocab.ParseGender(genderStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lang, err := vocab.ParseLanguage(langStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tokens := damn.NewDamner(lang).DamnYou(
		level,
		vocab.WithGender(gender),
		vocab.WithObscene(obscene),
	)

	fmt.Println(compose(tokens))
}

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
