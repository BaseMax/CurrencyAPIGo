package main

import (
	"strconv"
	"strings"
)

var (
	Currencies map[string][2]string = map[string][2]string{
		"usd": {"USD", "US Dollar"},
		"eur": {"EUR", "Euro"},
		"gbp": {"GBP", "British Pound"},
		"chf": {"CHF", "Swiss Franc"},
		"cad": {"CAD", "Canadian Dollar"},
		"aud": {"AUD", "Australian Dollar"},
		"sek": {"SEK", "Swedish Krona"},
		"nok": {"NOK", "Norwegian Krone"},
		"rub": {"RUB", "Russian Ruble"},
		"thb": {"THB", "Thai Baht"},
		"sgd": {"SGD", "Singapore Dollar"},
		"hkd": {"HKD", "Hong Kong Dollar"},
		"azn": {"AZN", "Azerbaijani Manat"},
		"amd": {"AMD", "Armenian Dram"},
		"dkk": {"DKK", "Danish Krone"},
		"aed": {"AED", "UAE Dirham"},
		"jpy": {"JPY", "Japanese Yen"},
		"try": {"TRY", "Turkish Lira"},
		"cny": {"CNY", "Chinese Yuan"},
		"sar": {"SAR", "KSA Rial"},
		"inr": {"INR", "Indian Rupee"},
		"myr": {"MYR", "Ringgit"},
		"afn": {"AFN", "Afghan Afghani"},
		"kwd": {"KWD", "Kuwaiti Dinar"},
		"iqd": {"IQD", "Iraqi Dinar"},
		"bhd": {"BHD", "Bahraini Dinar"},
		"omr": {"OMR", "Omani Rial"},
		"qar": {"QAR", "Qatari Rial"},
	}

	Coins map[string]string = map[string]string{
		"azadi1":   "Azadi",
		"azadi1_2": "Azadi 1/2",
		"azadi1_4": "Azadi 1/4",
		"emami":    "Emami",
		"azadi1g":  "Gerami",
	}
)

type Currency struct {
	Code string
	Name string
	Sell int
	Buy  int
}

type Coin struct {
	Name string
	Sell int
	Buy  int
}

type Gold struct {
	Name  string
	Price int
}

func GetCurrency(name string) (*Currency, error) {
	currencies, err := GetCurrencies()
	if err != nil || currencies == nil {
		return nil, err
	}

	sell, err := strconv.Atoi(currencies[name+"1"])
	if err != nil {
		return nil, err
	}

	buy, err := strconv.Atoi(currencies[name+"2"])
	if err != nil {
		return nil, err
	}

	return &Currency{
		Code: Currencies[name][0],
		Name: Currencies[name][1],
		Sell: sell,
		Buy:  buy,
	}, nil
}

func GetCoin(name string) (*Coin, error) {
	currencies, err := GetCurrencies()
	if err != nil || currencies == nil {
		return nil, err
	}

	sell, err := strconv.Atoi(currencies[name])
	if err != nil {
		return nil, err
	}

	buy, err := strconv.Atoi(currencies[name+"_2"])
	if err != nil {
		return nil, err
	}

	return &Coin{
		Name: Coins[name],
		Sell: sell,
		Buy:  buy,
	}, nil
}

func GetGold(name string) (*Gold, error) {
	currencies, err := GetCurrencies()
	if err != nil || currencies == nil {
		return nil, err
	}

	price, err := strconv.Atoi(currencies[name])
	if err != nil {
		return nil, err
	}

	return &Gold{
		Name:  strings.ToTitle(name),
		Price: price,
	}, nil
}
