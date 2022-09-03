package main

import (
	"context"
	"strconv"
	"strings"
)

var (
	Currencies map[string]string = map[string]string{
		"usd": "US Dollar",
		"eur": "Euro",
		"gbp": "British Pound",
		"chf": "Swiss Franc",
		"cad": "Canadian Dollar",
		"aud": "Australian Dollar",
		"sek": "Swedish Krona",
		"nok": "Norwegian Krone",
		"rub": "Russian Ruble",
		"thb": "Thai Baht",
		"sgd": "Singapore Dollar",
		"hkd": "Hong Kong Dollar",
		"azn": "Azerbaijani Manat",
		"amd": "Armenian Dram",
		"dkk": "Danish Krone",
		"aed": "UAE Dirham",
		"jpy": "Japanese Yen",
		"try": "Turkish Lira",
		"cny": "Chinese Yuan",
		"sar": "KSA Rial",
		"inr": "Indian Rupee",
		"myr": "Ringgit",
		"afn": "Afghan Afghani",
		"kwd": "Kuwaiti Dinar",
		"iqd": "Iraqi Dinar",
		"bhd": "Bahraini Dinar",
		"omr": "Omani Rial",
		"qar": "Qatari Rial",
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

func GetCurrency(ctx context.Context, name string) (*Currency, error) {
	currencies, err := GetCurrencies(ctx)
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
		Code: strings.ToTitle(name),
		Name: Currencies[name],
		Sell: sell,
		Buy:  buy,
	}, nil
}

func GetCoin(ctx context.Context, name string) (*Coin, error) {
	currencies, err := GetCurrencies(ctx)
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

func GetGold(ctx context.Context, name string) (*Gold, error) {
	currencies, err := GetCurrencies(ctx)
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
