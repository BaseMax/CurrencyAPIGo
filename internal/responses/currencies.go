package responses

import (
	"context"
	"net/http"

	"github.com/unrolled/render"

	"github.com/BaseMax/CurrencyAPIGo/internal/providers"
	"github.com/BaseMax/CurrencyAPIGo/pkg/utils"
)

type Currencies struct {
	Status     bool                       `json:"status" default:"true"`
	Currencies map[string]*CurrencyDetail `json:"currencies"`
	Coins      map[string]*CurrencyDetail `json:"coins"`
	Golds      map[string]*CurrencyDetail `json:"golds"`
	Cryptos    map[string]*CurrencyDetail `json:"cryptos"`
	LastModify string                     `json:"last_modify"`
}

type CurrencyResponse struct {
	Status   bool            `json:"status" default:"true"`
	Currency *CurrencyDetail `json:"currency"`
}

type CurrencyDetail struct {
	Name  string  `json:"name,omitempty"`
	Sell  int     `json:"sell,omitempty"`
	Buy   int     `json:"buy,omitempty"`
	Price float64 `json:"price,omitempty"`
}

func (c Currencies) Create(items map[string][]any, lastModify string) *Currencies {
	c.Status = true
	c.Currencies = make(map[string]*CurrencyDetail)
	c.Coins = make(map[string]*CurrencyDetail)
	c.Golds = make(map[string]*CurrencyDetail)
	c.Cryptos = make(map[string]*CurrencyDetail)

	for k := range items {
		c.CreateCurrencies(items[k])
	}

	c.LastModify = lastModify
	return &c
}

func (c *Currencies) CreateCurrencies(currencies []any) {
	for _, currency := range currencies {
		switch curr := currency.(*providers.Currency); {

		case utils.SliceContains(providers.CurrencyList, curr.Code):
			c.Currencies[curr.Code] = CurrencyDetail{}.Create(curr)

		case utils.SliceContains(providers.GoldCoinList, curr.Code):
			c.Coins[curr.Code] = CurrencyDetail{}.Create(curr)

		case utils.SliceContains(providers.GoldList, curr.Code):
			c.Golds[curr.Code] = CurrencyDetail{}.Create(curr)

		case utils.SliceContains(providers.CryptoCurrencyList, curr.Code):
			c.Cryptos[curr.Code] = CurrencyDetail{}.Create(curr)
		}
	}
}

func (c CurrencyDetail) Create(currency *providers.Currency) *CurrencyDetail {
	c.Name = currency.Name
	c.Sell = currency.Sell
	c.Buy = currency.Buy
	c.Price = currency.Price

	return &c
}

func RenderCurrenciesResponse(ctx context.Context, w http.ResponseWriter, currencies map[string]string) {
	items := map[string][]any{}
	r := ctx.Value(3).(*render.Render)

	for _, k := range utils.MapKeyToSlice(providers.CurrencyKeys) {
		if utils.SliceContains(providers.CurrencyList, k) {
			c, _ := providers.GetCurrency(ctx, k)
			items["currency"] = append(items["currency"], c)
		} else if utils.SliceContains(providers.GoldCoinList, k) {
			gc, _ := providers.GetGoldCoin(ctx, k)
			items["coins"] = append(items["coins"], gc)
		} else if utils.SliceContains(providers.GoldList, k) {
			g, _ := providers.GetGold(ctx, k)
			items["golds"] = append(items["golds"], g)
		} else if utils.SliceContains(providers.CryptoCurrencyList, k) {
			cc, _ := providers.GetCryptoCurrency(ctx, k)
			items["cryptos"] = append(items["cyptos"], cc)
		}
	}

	csr := &Currencies{}
	res := csr.Create(items, currencies["last_modified"])

	r.JSON(w, http.StatusOK, res)
}

func RenderCoinResponse(ctx context.Context, w http.ResponseWriter, coin *providers.Currency) {
	r := ctx.Value(3).(*render.Render)

	res := CurrencyResponse{
		Status:   true,
		Currency: CurrencyDetail{}.Create(coin),
	}

	r.JSON(w, http.StatusOK, res)
}

func RenderCurrencyResponse(ctx context.Context, w http.ResponseWriter, currency *providers.Currency) {
	r := ctx.Value(3).(*render.Render)

	res := CurrencyResponse{
		Status:   true,
		Currency: CurrencyDetail{}.Create(currency),
	}

	r.JSON(w, http.StatusOK, res)
}

func RenderGoldResponse(ctx context.Context, w http.ResponseWriter, gold *providers.Currency) {
	r := ctx.Value(3).(*render.Render)

	res := CurrencyResponse{
		Status:   true,
		Currency: CurrencyDetail{}.Create(gold),
	}

	r.JSON(w, http.StatusOK, res)
}

func RenderCryptoCurrencyResponse(ctx context.Context, w http.ResponseWriter, crypto *providers.Currency) {
	r := ctx.Value(3).(*render.Render)

	res := CurrencyResponse{
		Status:   true,
		Currency: CurrencyDetail{}.Create(crypto),
	}

	r.JSON(w, http.StatusOK, res)
}
