package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gomarkdown/markdown"

	"github.com/itsjoniur/currency/internal/providers"
	"github.com/itsjoniur/currency/internal/responses"
	"github.com/itsjoniur/currency/pkg/utils"
)

func RootHandler(w http.ResponseWriter, req *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Panicln("can not found root directory")
	}

	dat, err := os.ReadFile(path.Join(wd, "README.md"))
	if err != nil {
		fmt.Println(err)
		responses.InternalServerError(req.Context(), w)
	}

	md := []byte(dat)
	docs := markdown.ToHTML(md, nil, nil)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(docs)
}

func CurrencyHandler(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("q")

	if name == "" {
		cs, err := providers.GetCurrencies(req.Context())
		if err != nil {
			fmt.Println(err)
			responses.InternalServerError(req.Context(), w)
			return
		}
		responses.RenderCurrenciesResponse(req.Context(), w, cs)
		return
	}

	switch name = strings.ToLower(name); {
	case utils.SliceContains(providers.CurrencyList, name):
		currency, err := providers.GetCurrency(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
		}
		responses.RenderCurrencyResponse(req.Context(), w, currency)

	case utils.SliceContains(providers.GoldCoinList, name):
		coin, err := providers.GetGoldCoin(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
			return
		}
		responses.RenderCoinResponse(req.Context(), w, coin)

	case utils.SliceContains(providers.GoldList, name):
		gold, err := providers.GetGold(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
			return
		}
		responses.RenderGoldResponse(req.Context(), w, gold)

	case utils.SliceContains(providers.CryptoCurrencyList, name):
		crypto, err := providers.GetCryptoCurrency(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
		}
		responses.RenderCryptoCurrencyResponse(req.Context(), w, crypto)

	default:
		responses.NotFoundError(req.Context(), w)
	}
}
