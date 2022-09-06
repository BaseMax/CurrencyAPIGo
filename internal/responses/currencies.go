package responses

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
	"golang.org/x/exp/slices"

	"github.com/itsjoniur/currency/internal/providers"
	"github.com/itsjoniur/currency/pkg/utils"
)

type Currencies struct {
	Status     bool                       `json:"status" default:"true"`
	Currencies map[string]*CurrencyDetail `json:"currencies"`
	Coins      map[string]*CoinDetail     `json:"coins"`
	Golds      map[string]*GoldDetail     `json:"golds"`
	Cryptos    map[string]*CryptoDetail   `json:"cryptos"`
	LastModify string                     `json:"last_modify"`
}

type CurrencyResponse struct {
	Status   bool `json:"status" default:"true"`
	Currency *CurrencyDetail
}

type CoinResponse struct {
	Status bool        `json:"status" default:"true"`
	Coin   *CoinDetail `json:"coin"`
}

type GoldResponse struct {
	Status bool        `json:"status"`
	Gold   *GoldDetail `json:"gold"`
}

type CryptoResponse struct {
	Status bool          `json:"status"`
	Crypto *CryptoDetail `json:"crypto"`
}

type CurrencyDetail struct {
	Currency string `json:"currency"`
	Sell     int    `json:"sell"`
	Buy      int    `json:"buy"`
}

type CoinDetail struct {
	Name string `json:"name"`
	Sell int    `json:"sell"`
	Buy  int    `json:"buy"`
}

type GoldDetail struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CryptoDetail struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (c Currencies) Create(items map[string][]any, lastModify string) *Currencies {
	c.Status = true
	c.Currencies = make(map[string]*CurrencyDetail)
	c.Coins = make(map[string]*CoinDetail)
	c.Golds = make(map[string]*GoldDetail)
	c.Cryptos = make(map[string]*CryptoDetail)

	for k := range items {
		c.CreateCurrencies(items[k])
	}

	c.LastModify = lastModify
	return &c
}

func (c *Currencies) CreateCurrencies(currencies []any) {
	for _, crncy := range currencies {
		switch crncy.(type) {
		case *providers.Currency:
			currency := crncy.(*providers.Currency)
			c.Currencies[currency.Code] = CurrencyDetail{}.Create(currency)
		case *providers.Coin:
			coin := crncy.(*providers.Coin)
			c.Coins[coin.Code] = CoinDetail{}.Create(coin)
		case *providers.Gold:
			gold := crncy.(*providers.Gold)
			c.Golds[gold.Code] = GoldDetail{}.Create(gold)
		case *providers.Crypto:
			crypto := crncy.(*providers.Crypto)
			c.Cryptos[crypto.Code] = CryptoDetail{}.Create(crypto)
		}
	}
}

func (c CurrencyDetail) Create(currency *providers.Currency) *CurrencyDetail {
	c.Currency = currency.Name
	c.Sell = currency.Sell
	c.Buy = currency.Buy

	return &c
}

func (c CoinDetail) Create(coin *providers.Coin) *CoinDetail {
	c.Name = coin.Name
	c.Sell = coin.Sell
	c.Buy = coin.Buy

	return &c
}

func (c GoldDetail) Create(gold *providers.Gold) *GoldDetail {
	c.Name = gold.Name
	c.Price = gold.Price

	return &c
}

func (c CryptoDetail) Create(crypto *providers.Crypto) *CryptoDetail {
	c.Name = crypto.Name
	c.Price = crypto.Price

	return &c
}

func RenderCurrenciesResponse(ctx context.Context, w http.ResponseWriter, currencies map[string]string) {
	items := map[string][]any{}
	r := ctx.Value(3).(*render.Render)

	for _, k := range utils.MapKeyToSlice(providers.CurrencyKeys) {
		if slices.Contains(providers.CurrencyList, k) {
			c, _ := providers.GetCurrency(ctx, k)
			items["currency"] = append(items["currency"], c)
		} else if slices.Contains(providers.GoldCoinList, k) {
			gc, _ := providers.GetGoldCoin(ctx, k)
			items["coins"] = append(items["coins"], gc)
		} else if slices.Contains(providers.GoldList, k) {
			g, _ := providers.GetGold(ctx, k)
			items["golds"] = append(items["golds"], g)
		} else if slices.Contains(providers.CryptoCurrencyList, k) {
			cc, _ := providers.GetCryptoCurrency(ctx, k)
			items["cryptos"] = append(items["cyptos"], cc)
		}
	}

	csr := &Currencies{}
	res := csr.Create(items, currencies["last_modified"])

	r.JSON(w, http.StatusOK, res)
}

func RenderCoinResponse(ctx context.Context, w http.ResponseWriter, coin *providers.Coin) {
	r := ctx.Value(3).(*render.Render)

	res := CoinResponse{
		Status: true,
		Coin:   CoinDetail{}.Create(coin),
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

func RenderGoldResponse(ctx context.Context, w http.ResponseWriter, gold *providers.Gold) {
	r := ctx.Value(3).(*render.Render)

	res := GoldResponse{
		Status: true,
		Gold:   GoldDetail{}.Create(gold),
	}

	r.JSON(w, http.StatusOK, res)
}

func RenderCryptoCurrencyResponse(ctx context.Context, w http.ResponseWriter, crypto *providers.Crypto) {
	r := ctx.Value(3).(*render.Render)

	res := CryptoResponse{
		Status: true,
		Crypto: CryptoDetail{}.Create(crypto),
	}

	r.JSON(w, http.StatusOK, res)
}
