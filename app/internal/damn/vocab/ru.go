package vocab

import "ktolstikhin/damn/internal/util"

var (
	masculineAbusiveAdjectivesRU = []string{
		"толстый",
		"жирный",
		"тощий",
		"лысоголовый",
		"конопатый",
		"курносый",
		"косоглазый",
		"пучеглазый",
		"кривоногий",
		"косолапый",
		"хромоногий",
		"криволапый",
		"криворогий",
		"плоскожопый",
		"толстожопый",
		"волосатый",
		"грязный",
		"вонючий",
		"потный",
		"зловонный",
		"ужасный",
		"неопрятный",
		"глухой",
		"глухонемой",
		"слепой",
		"неприятный",
		"противный",
		"рваный",
		"потасканный",
		"второсортный",
		"тупой",
		"облезлый",
		"плешивый",
		"ботоксный",
		"токсичный",
		"высохший",
		"тупорогий",
		"мохнорылый",
		"лысоногий",
		"жирноногий",
		"жирножопый",
		"тощеногий",
		"кривоклювый",
		"гнойный",
		"депрессивный",
		"душевнобольной",
		"всратый",
		"страшный",
		"конченный",
		"бухой",
		"вмазанный",
		"под кайфом",
	}
	masculineObsceneAdjectivesRU = append([]string{
		"хуёвый",
		"хуеблядский",
		"хуеносый",
		"хуерукий",
		"хуеногий",
		"хуеголовый",
		"хуежопый",
		"ебанутый",
		"ёбнутый",
		"ебучий",
		"лихоеблимудый",
		"лихоебучий",
		"пизданутый",
		"пиздоблядский",
		"пиздоногий",
		"пиздорукий",
		"пиздоголовый",
		"пиздокрылый",
		"уебищный",
		"толстомудый",
		"опиздюленный",
		"пиздоебучий",
		"хуячий",
		"аналочленный",
		"кривохуевый",
		"халдоблядский",
		"блядский",
		"охуевший",
		"спермоблядский",
		"долбоблядский",
		"ублюдский",
		"хитровыебанный",
		"блядовыебанный",
		"пиздомудохвостный",
		"пиздопидормотный",
		"хуеклювый",
		"херопиздокрылый",
	}, masculineAbusiveAdjectivesRU...)
	masculineAbusiveNounsRU = []string{
		"кабан",
		"бычара",
		"козёл",
		"осёл",
		"тюлень",
		"мудозвон",
		"мудила",
		"мудак",
		"дебил",
		"ушлёпок",
		"идиот",
		"маргинал",
		"торчок",
		"наркоман",
		"бич",
		"бичара",
		"алкаш",
		"алконавт",
		"имбецил",
		"даун",
		"борщевик",
		"гриб",
		"холодец",
		"баклан",
		"говноед",
		"шизик",
		"шизомудозвон",
		"шизоид",
		"психопат",
		"спермобак",
		"говноед",
		"жиробас",
		"хер",
		"херувим",
		"бобёр",
		"бармалей",
		"Валера",
		"Артём",
		"Олег",
		"выродок",
		"мутант",
		"осётр",
		"карась",
		"слон",
		"носорог",
	}
	masculineObsceneNounsRU = append([]string{
		"хуй",
		"пиздохуемот",
		"херопиздокрыл",
		"пидорас",
		"пидрила",
		"хуеплёт",
		"хуешланг",
		"хуеклюв",
		"уёбок",
		"пиздобол",
		"пидормот",
		"пиздюк",
		"пиздомудохвост",
		"еблан",
		"хуйлан",
		"жопоблядник",
		"хуепидор",
		"пиздопидормот",
		"хероблядник",
		"шизопиздокрыл",
		"ослаёб",
		"овцеёб",
	}, masculineAbusiveNounsRU...)

	feminineAbusiveAdjectivesRU = []string{}
	feminineObsceneAdjectivesRU = append([]string{}, feminineAbusiveAdjectivesRU...)
	feminineAbusiveNounsRU      = []string{}
	feminineObsceneNounsRU      = append([]string{}, feminineAbusiveNounsRU...)

	conjunctionsAbusiveRU = []string{
		"и",
		"внатуре",
		"нахер",
		"вообще",
		"в жопу",
		"в говно",
		"в дрова",
	}
	conjunctionsObsceneRU = append([]string{
		"нах",
		"нахуй",
		"бля",
		"блядь",
		"впизду",
		"без пизды",
		"ебать",
		"ебать-копать",
		"ебать-колоть",
		"ебать-не-встать",
		"ебать-не-переебать",
	}, conjunctionsAbusiveRU...)

	masculineAbusiveAdditionsRU = map[string][]string{
		"орущий": {
			"на море",
			"себе в жопу",
			"себе в ухо",
			"в небо",
			"на говно",
			"в истерике",
			"в ужасе",
			"в страхе",
			"и срущий",
			"и злоебучий",
		},
		"нюхающий": {
			"говно",
			"жопу орангутана",
			"жопу носорога",
			"жопу скунса",
			"свои подмышки",
			"фекалии копибары",
			"вонючие носки",
		},
		"застрявший": {
			"в жопе бомжа",
			"в жопе мамонта",
			"в жопе динозавра",
		},
		"жрущий": {
			"гавно мамонта",
			"фекалии носорога",
			"дохлых медуз",
			"склизкие водоросли",
			"вонючие носки",
		},
		"срущий": {
			"гвоздями",
			"дикобразами",
			"всратыми пауками",
		},
	}
	masculineObsceneAdditionsRU = util.MergeMaps(masculineAbusiveAdditionsRU, map[string][]string{
		"сосущий": {
			"хуй",
			"хуй обезьяны",
			"хуй медузы",
			"хуй бомжа",
			"сиськи хуебляди",
		},
		"ебущий": {
			"дикобразов",
			"обезьян",
			"сраных обезьян",
			"бомжей",
			"грязных бомжей",
			"пиздохуемразей",
		},
	})
)
