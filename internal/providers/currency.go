package providers

import (
	"context"
	"strconv"
)

type Currency struct {
	Code string
	Name string
	Sell int
	Buy  int
}

type Coin struct {
	Code string
	Name string
	Sell int
	Buy  int
}

type Gold struct {
	Code  string
	Name  string
	Price float64
}

func GetCurrency(ctx context.Context, code string) (*Currency, error) {
	currencies, err := GetCurrencies(ctx)
	if err != nil || currencies == nil {
		return nil, err
	}

	sell, err := strconv.Atoi(currencies[CurrencyKeys[code]["sell"]])
	if err != nil {
		return nil, err
	}

	buy, err := strconv.Atoi(currencies[CurrencyKeys[code]["buy"]])
	if err != nil {
		return nil, err
	}

	return &Currency{
		Code: code,
		Name: CurrencyKeys[code]["name"],
		Sell: sell,
		Buy:  buy,
	}, nil
}

func GetCoin(ctx context.Context, code string) (*Coin, error) {
	currencies, err := GetCurrencies(ctx)
	if err != nil || currencies == nil {
		return nil, err
	}

	sell, err := strconv.Atoi(currencies[CurrencyKeys[code]["sell"]])
	if err != nil {
		return nil, err
	}

	buy, err := strconv.Atoi(currencies[CurrencyKeys[code]["buy"]])
	if err != nil {
		return nil, err
	}

	return &Coin{
		Code: code,
		Name: CurrencyKeys[code]["name"],
		Sell: sell,
		Buy:  buy,
	}, nil
}

func GetGold(ctx context.Context, code string) (*Gold, error) {
	currencies, err := GetCurrencies(ctx)
	if err != nil || currencies == nil {
		return nil, err
	}

	price, err := strconv.ParseFloat(currencies[CurrencyKeys[code]["sell"]], 64)
	if err != nil {
		return nil, err
	}

	return &Gold{
		Code:  code,
		Name:  CurrencyKeys[code]["name"],
		Price: price,
	}, nil
}
