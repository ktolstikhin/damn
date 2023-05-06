package main

import (
	"flag"
	"fmt"
	"os"

	"ktolstikhin/damn/internal/damn"
	"ktolstikhin/damn/internal/damn/lang"
)

func main() {
	var (
		gnd     string
		lng     string
		level   int
		obscene bool
	)
	flag.StringVar(&gnd, "gender", "male", "Gender of the person to damn: male, female.")
	flag.StringVar(&lng, "language", "ru", "God damn language: ru.")
	flag.IntVar(&level, "level", 1, "God damn level: from 1 to 3.")
	flag.BoolVar(&obscene, "obscene", false, "Usage of obscene vocabulary.")
	flag.Parse()

	s, err := damn.DamnYou(lang.Gender(gnd), lang.Language(lng), obscene, level)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(s)
}
