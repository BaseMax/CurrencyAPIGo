package providers

type Map map[string]map[string]string
type List []string

var CurrencyList List = List{
	"usd", "eur", "gbp", "chf",
	"cad", "aud", "sek", "nok",
	"rub", "thb", "sgd", "hkd",
	"azn", "amd", "dkk", "aed",
	"jpy", "try", "cny", "sar",
	"inr", "myr", "afn", "kwd",
	"iqd", "bhd", "omr", "qar",
}

var GoldCoinList List = List{
	"azadi",
	"azadi12",
	"azadi14",
	"emami",
	"gerami",
}

var GoldList List = List{
	"gram",
	"mithqal",
	"ounce",
}

var CryptoCurrencyList List = List{
	"bitcoin",
}

var CurrencyKeys Map = Map{
	"usd": {
		"name": "US Dollar",
		"sell": "usd1",
		"buy":  "usd2",
	},
	"eur": {
		"name": "Euro",
		"sell": "eur1",
		"buy":  "eur2",
	},
	"gbp": {
		"name": "British Pound",
		"sell": "gbp1",
		"buy":  "gbp2",
	},
	"chf": {
		"name": "Swiss Franc",
		"sell": "chf1",
		"buy":  "chf2",
	},
	"cad": {
		"name": "Canadian Dollar",
		"sell": "cad1",
		"buy":  "cad2",
	},
	"aud": {
		"name": "Australian Dollar",
		"sell": "aud1",
		"buy":  "aud2",
	},
	"sek": {
		"name": "Swedish Krone",
		"sell": "sek1",
		"buy":  "sek2",
	},
	"nok": {
		"name": "Norwegian Krone",
		"sell": "nok1",
		"buy":  "nok2",
	},
	"rub": {
		"name": "Russian Ruble",
		"sell": "rub1",
		"buy":  "rub2",
	},
	"thb": {
		"name": "Thai Baht",
		"sell": "thb1",
		"buy":  "thb2",
	},
	"sgd": {
		"name": "Singapore Dollar",
		"sell": "sgd1",
		"buy":  "sgd2",
	},
	"hkd": {
		"name": "Hong Kong Dollar",
		"sell": "hkd1",
		"buy":  "hkd2",
	},
	"azn": {
		"name": "Azerbaijani Manat",
		"sell": "azn1",
		"buy":  "azn2",
	},
	"amd": {
		"name": "Armenian Dram",
		"sell": "amd1",
		"buy":  "amd2",
	},
	"dkk": {
		"name": "Danish Krone",
		"sell": "dkk1",
		"buy":  "dkk2",
	},
	"aed": {
		"name": "UAE Dirham",
		"sell": "aed1",
		"buy":  "aed2",
	},
	"jpy": {
		"name": "Japanese Yen",
		"sell": "jpy1",
		"buy":  "jpy2",
	},
	"try": {
		"name": "Turkish Lira",
		"sell": "try1",
		"buy":  "try2",
	},
	"cny": {
		"name": "Chinese Yuan",
		"sell": "cny1",
		"buy":  "cny2",
	},
	"sar": {
		"name": "KSA Rial",
		"sell": "sar1",
		"buy":  "sar2",
	},
	"inr": {
		"name": "Indian Rupee",
		"sell": "inr1",
		"buy":  "inr2",
	},
	"myr": {
		"name": "Ringgit",
		"sell": "myr1",
		"buy":  "myr2",
	},
	"afn": {
		"name": "Afghan Afghani",
		"sell": "afn1",
		"buy":  "afn2",
	},
	"kwd": {
		"name": "Kuwaiti Dinar",
		"sell": "kwd1",
		"buy":  "kwd2",
	},
	"iqd": {
		"name": "Iraqi Dinar",
		"sell": "iqd1",
		"buy":  "iqd2",
	},
	"bhd": {
		"name": "Bahraini Dinar",
		"sell": "bhd1",
		"buy":  "bhd2",
	},
	"omr": {
		"name": "Omani Rial",
		"sell": "omr1",
		"buy":  "omr2",
	},
	"qar": {
		"name": "Qatari Rial",
		"sell": "qar1",
		"buy":  "qar2",
	},
	// Gold Coins
	"azadi": {
		"name": "Azadi Coin",
		"sell": "azadi1",
		"buy":  "azadi12",
	},
	"azadi12": {
		"name": "Azadi Half a coin",
		"sell": "azadi1_2",
		"buy":  "azadi1_22",
	},
	"azadi14": {
		"name": "Azadi Quarter a coin",
		"sell": "azadi1_4",
		"buy":  "azadi1_42",
	},
	"emami": {
		"name": "Emami Coin",
		"sell": "emami1",
		"buy":  "emami12",
	},
	"gerami": {
		"name": "Gerami Coin",
		"sell": "azadi1g",
		"buy":  "azadi1g2",
	},
	// Golds
	// note: golds does not have buy price
	"gram": {
		"name": "Gram",
		"sell": "gol18",
		"buy":  "",
	},
	"mithqal": {
		"name": "Mithqal",
		"sell": "mithqal",
		"buy":  "",
	},
	"ounce": {
		"name": "Ounce",
		"sell": "ounce",
		"buy":  "",
	},
	// Crypto Currency
	// note: Crypto Currencies does not have buy price
	"bitcoin": {
		"name": "Bitcoin",
		"sell": "bitcoin",
		"buy":  "",
	},
}
