package vocab

import (
	"encoding/json"
	"fmt"
)

type Language byte

const (
	LanguageRU Language = iota
)

var (
	strToLangMap = map[string]Language{
		"ru": LanguageRU,
	}
	langToStrMap = map[Language]string{
		LanguageRU: "ru",
	}
)

func ParseLanguage(s string) (Language, error) {
	var (
		lang Language
		err  error
	)
	lang, ok := strToLangMap[s]
	if !ok {
		err = fmt.Errorf("unknown language: %s", s)
	}

	return lang, err
}

func (l Language) String() string {
	return langToStrMap[l]
}

func (l Language) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

func (l *Language) UnmarshalJSON(data []byte) (err error) {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if *l, err = ParseLanguage(s); err != nil {
		return err
	}

	return nil
}
