package responses

import (
	"context"
	"net/http"
	"strings"

	"github.com/itsjoniur/currency/internal/providers"
	"github.com/itsjoniur/currency/pkg/utils"
	"github.com/unrolled/render"
)

type Currencies struct {
	Status     bool                       `json:"status" default:"true"`
	Currencies map[string]*CurrencyDetail `json:"Currencies"`
	Coins      map[string]*CoinDetail     `json:"coins"`
	Golds      map[string]*GoldDetail     `json:"golds"`
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

func (c Currencies) Create(currencies []*providers.Currency, coins []*providers.Coin, golds []*providers.Gold, lastModify string) *Currencies {
	c.Status = true
	c.CreateCurrencies(currencies)
	c.CreateCoins(coins)
	c.CreateGolds(golds)
	c.LastModify = lastModify
	return &c
}

func (c *Currencies) CreateCurrencies(cs []*providers.Currency) {
	c.Currencies = make(map[string]*CurrencyDetail)
	for i := 0; i < len(cs); i++ {
		c.Currencies[strings.ToLower(cs[i].Code)] = CurrencyDetail{}.Create(cs[i])
	}
}

func (c *Currencies) CreateCoins(coins []*providers.Coin) {
	c.Coins = make(map[string]*CoinDetail)
	for i := 0; i < len(coins); i++ {
		name := strings.Replace(strings.ToLower(coins[i].Name), " ", "", 1)
		c.Coins[name] = CoinDetail{}.Create(coins[i])
	}
}

func (c *Currencies) CreateGolds(golds []*providers.Gold) {
	c.Golds = make(map[string]*GoldDetail)
	for i := 0; i < len(golds); i++ {
		c.Golds[strings.ToLower(golds[i].Name)] = GoldDetail{}.Create(golds[i])
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

func RenderCurrenciesResponse(ctx context.Context, w http.ResponseWriter, currencies map[string]string) {
	var cs []*providers.Currency
	var coins []*providers.Coin
	var golds []*providers.Gold
	r := ctx.Value(3).(*render.Render)

	for _, i := range utils.MapKeyToSlice(providers.Currencies) {
		c, _ := providers.GetCurrency(ctx, i)
		if c != nil {
			cs = append(cs, c)
		}
	}

	for _, i := range utils.MapKeyToSlice(providers.Coins) {
		c, _ := providers.GetCoin(ctx, i)
		if c != nil {
			coins = append(coins, c)
		}
	}

	for _, i := range utils.MapKeyToSlice(providers.Golds) {
		g, _ := providers.GetGold(ctx, i)
		if g != nil {
			golds = append(golds, g)
		}
	}

	csr := &Currencies{}
	res := csr.Create(cs, coins, golds, currencies["last_modified"])

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
