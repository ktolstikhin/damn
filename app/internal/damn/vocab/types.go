package vocab

type Gender byte

const (
	GenderMasculine Gender = iota
	GenderFeminine
)

var StrToGenderMap = map[string]Gender{
	"m": GenderMasculine,
	"f": GenderFeminine,
}

type Language byte

const (
	LanguageRU Language = iota
)

var StrToLanguageMap = map[string]Language{
	"ru": LanguageRU,
}
