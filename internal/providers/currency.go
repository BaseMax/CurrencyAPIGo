package providers

import (
	"context"
	"errors"
	"fmt"
	"strconv"
)

type Currency struct {
	Code  string
	Name  string
	Sell  int
	Buy   int
	Price float64
}

func GetCurrency(ctx context.Context, code string) (*Currency, error) {
	currencies, err := GetCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	if currencies == nil {
		return nil, errors.New("data is empty")
	}

	curr := CurrencyKeys[code]
	if curr == (CurrencyData{}) {
		return nil, errors.New(fmt.Sprintf("currency %s not found", code))
	}

	sell, err := strconv.Atoi(currencies[curr.Sell])
	if err != nil {
		return nil, err
	}

	buy, err := strconv.Atoi(currencies[curr.Buy])
	if err != nil {
		return nil, err
	}

	return &Currency{
		Code: code,
		Name: curr.Name,
		Sell: sell,
		Buy:  buy,
	}, nil
}

func GetGoldCoin(ctx context.Context, code string) (*Currency, error) {
	currencies, err := GetCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	if currencies == nil {
		return nil, errors.New("data is empty")
	}

	coin := CurrencyKeys[code]
	if coin == (CurrencyData{}) {
		return nil, errors.New(fmt.Sprintf("coin %s not found", code))
	}

	sell, err := strconv.Atoi(currencies[coin.Sell])
	if err != nil {
		return nil, err
	}

	buy, err := strconv.Atoi(currencies[coin.Buy])
	if err != nil {
		return nil, err
	}

	return &Currency{
		Code: code,
		Name: coin.Name,
		Sell: sell,
		Buy:  buy,
	}, nil
}

func GetGold(ctx context.Context, code string) (*Currency, error) {
	currencies, err := GetCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	if currencies == nil {
		return nil, errors.New("data is empty")
	}

	gold := CurrencyKeys[code]
	if gold == (CurrencyData{}) {
		return nil, errors.New(fmt.Sprintf("gold %s not found", code))
	}

	price, err := strconv.ParseFloat(currencies[gold.Sell], 64)
	if err != nil {
		return nil, err
	}

	return &Currency{
		Code:  code,
		Name:  gold.Name,
		Price: price,
	}, nil
}

func GetCryptoCurrency(ctx context.Context, code string) (*Currency, error) {
	currencies, err := GetCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	if currencies == nil {
		return nil, errors.New("data is empty")
	}

	crypto := CurrencyKeys[code]
	if crypto == (CurrencyData{}) {
		return nil, errors.New(fmt.Sprintf("crypto %s not found", code))
	}

	price, err := strconv.ParseFloat(currencies[crypto.Sell], 64)
	if err != nil {
		return nil, err
	}

	return &Currency{
		Code:  code,
		Name:  crypto.Name,
		Price: price,
	}, nil
}
