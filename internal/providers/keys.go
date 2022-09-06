package providers

type CurrencyData struct {
	Name string
	Sell string
	Buy  string
}

var CurrencyList = []string{
	"usd", "eur", "gbp", "chf",
	"cad", "aud", "sek", "nok",
	"rub", "thb", "sgd", "hkd",
	"azn", "amd", "dkk", "aed",
	"jpy", "try", "cny", "sar",
	"inr", "myr", "afn", "kwd",
	"iqd", "bhd", "omr", "qar",
}

var GoldCoinList = []string{
	"azadi",
	"azadi12",
	"azadi14",
	"emami",
	"gerami",
}

var GoldList = []string{
	"gram",
	"mithqal",
	"ounce",
}

var CryptoCurrencyList = []string{
	"bitcoin",
}

var CurrencyKeys = map[string]CurrencyData{
	"usd": CurrencyData{
		Name: "US Dollar",
		Sell: "usd1",
		Buy:  "usd2",
	},
	"eur": CurrencyData{
		Name: "Euro",
		Sell: "eur1",
		Buy:  "eur2",
	},
	"gbp": CurrencyData{
		Name: "British Pound",
		Sell: "gbp1",
		Buy:  "gbp2",
	},
	"chf": CurrencyData{
		Name: "Swiss Franc",
		Sell: "chf1",
		Buy:  "chf2",
	},
	"cad": CurrencyData{
		Name: "Canadian Dollar",
		Sell: "cad1",
		Buy:  "cad2",
	},
	"aud": CurrencyData{
		Name: "Australian Dollar",
		Sell: "aud1",
		Buy:  "aud2",
	},
	"sek": CurrencyData{
		Name: "Swedish Krone",
		Sell: "sek1",
		Buy:  "sek2",
	},
	"nok": CurrencyData{
		Name: "Norwegian Krone",
		Sell: "nok1",
		Buy:  "nok2",
	},
	"rub": CurrencyData{
		Name: "Russian Ruble",
		Sell: "rub1",
		Buy:  "rub2",
	},
	"thb": CurrencyData{
		Name: "Thai Baht",
		Sell: "thb1",
		Buy:  "thb2",
	},
	"sgd": CurrencyData{
		Name: "Singapore Dollar",
		Sell: "sgd1",
		Buy:  "sgd2",
	},
	"hkd": CurrencyData{
		Name: "Hong Kong Dollar",
		Sell: "hkd1",
		Buy:  "hkd2",
	},
	"azn": CurrencyData{
		Name: "Azerbaijani Manat",
		Sell: "azn1",
		Buy:  "azn2",
	},
	"amd": CurrencyData{
		Name: "Armenian Dram",
		Sell: "amd1",
		Buy:  "amd2",
	},
	"dkk": CurrencyData{
		Name: "Danish Krone",
		Sell: "dkk1",
		Buy:  "dkk2",
	},
	"aed": CurrencyData{
		Name: "UAE Dirham",
		Sell: "aed1",
		Buy:  "aed2",
	},
	"jpy": CurrencyData{
		Name: "Japanese Yen",
		Sell: "jpy1",
		Buy:  "jpy2",
	},
	"try": CurrencyData{
		Name: "Turkish Lira",
		Sell: "try1",
		Buy:  "try2",
	},
	"cny": CurrencyData{
		Name: "Chinese Yuan",
		Sell: "cny1",
		Buy:  "cny2",
	},
	"sar": CurrencyData{
		Name: "KSA Rial",
		Sell: "sar1",
		Buy:  "sar2",
	},
	"inr": CurrencyData{
		Name: "Indian Rupee",
		Sell: "inr1",
		Buy:  "inr2",
	},
	"myr": CurrencyData{
		Name: "Ringgit",
		Sell: "myr1",
		Buy:  "myr2",
	},
	"afn": CurrencyData{
		Name: "Afghan Afghani",
		Sell: "afn1",
		Buy:  "afn2",
	},
	"kwd": CurrencyData{
		Name: "Kuwaiti Dinar",
		Sell: "kwd1",
		Buy:  "kwd2",
	},
	"iqd": CurrencyData{
		Name: "Iraqi Dinar",
		Sell: "iqd1",
		Buy:  "iqd2",
	},
	"bhd": CurrencyData{
		Name: "Bahraini Dinar",
		Sell: "bhd1",
		Buy:  "bhd2",
	},
	"omr": CurrencyData{
		Name: "Omani Rial",
		Sell: "omr1",
		Buy:  "omr2",
	},
	"qar": CurrencyData{
		Name: "Qatari Rial",
		Sell: "qar1",
		Buy:  "qar2",
	},
	// Gold Coins
	"azadi": CurrencyData{
		Name: "Azadi Coin",
		Sell: "azadi1",
		Buy:  "azadi12",
	},
	"azadi12": CurrencyData{
		Name: "Azadi Half a coin",
		Sell: "azadi1_2",
		Buy:  "azadi1_22",
	},
	"azadi14": CurrencyData{
		Name: "Azadi Quarter a coin",
		Sell: "azadi1_4",
		Buy:  "azadi1_42",
	},
	"emami": CurrencyData{
		Name: "Emami Coin",
		Sell: "emami1",
		Buy:  "emami12",
	},
	"gerami": CurrencyData{
		Name: "Gerami Coin",
		Sell: "azadi1g",
		Buy:  "azadi1g2",
	},
	// Golds
	// note: golds does not have buy price
	"gram": CurrencyData{
		Name: "Gram",
		Sell: "gol18",
	},
	"mithqal": CurrencyData{
		Name: "Mithqal",
		Sell: "mithqal",
	},
	"ounce": CurrencyData{
		Name: "Ounce",
		Sell: "ounce",
	},
	// Crypto Currency
	// note: Crypto Currencies does not have buy price
	"bitcoin": CurrencyData{
		Name: "Bitcoin",
		Sell: "bitcoin",
	},
}
