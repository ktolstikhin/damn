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
	flag.StringVar(&genderStr, "gender", "m", "Gender to damn: m - male, f - female.")
	flag.StringVar(&langStr, "language", "ru", "God damn language: ru.")
	flag.IntVar(&level, "level", 1, "God damn level: from 1 to 4.")
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

	s := damn.NewDamner(lng).DamnYou(
		level,
		vocab.WithGender(gender),
		vocab.WithObscene(obscene),
	)

	fmt.Println(s)
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
