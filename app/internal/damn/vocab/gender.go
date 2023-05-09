package vocab

import (
	"encoding/json"
	"fmt"
)

type Gender byte

const (
	GenderMasculine Gender = iota
	GenderFeminine
)

var (
	strToGenderMap = map[string]Gender{
		"male":   GenderMasculine,
		"female": GenderFeminine,
	}
	genderToStrMap = map[Gender]string{
		GenderMasculine: "male",
		GenderFeminine:  "female",
	}
)

func ParseGender(s string) (Gender, error) {
	var (
		gender Gender
		err    error
	)
	gender, ok := strToGenderMap[s]
	if !ok {
		err = fmt.Errorf("unknown gender: %s", s)
	}

	return gender, err
}

func (g Gender) String() string {
	return genderToStrMap[g]
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

func (g *Gender) UnmarshalJSON(data []byte) (err error) {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if *g, err = ParseGender(s); err != nil {
		return err
	}

	return nil
}
