package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gomarkdown/markdown"
	"golang.org/x/exp/slices"

	"github.com/itsjoniur/currency/internal/providers"
	"github.com/itsjoniur/currency/internal/responses"
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
			responses.InternalServerError(req.Context(), w)
			return
		}
		responses.RenderCurrenciesResponse(req.Context(), w, cs)
		return
	}

	name = strings.ToLower(name)

	if slices.Contains(providers.CurrencyList, name) {
		currency, err := providers.GetCurrency(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
		}
		responses.RenderCurrencyResponse(req.Context(), w, currency)
		return
	} else if slices.Contains(providers.GoldCoinList, name) {
		coin, err := providers.GetGoldCoin(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
			return
		}
		responses.RenderCoinResponse(req.Context(), w, coin)
		return
	} else if slices.Contains(providers.GoldList, name) {
		gold, err := providers.GetGold(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
			return
		}
		responses.RenderGoldResponse(req.Context(), w, gold)

	} else if slices.Contains(providers.CryptoCurrencyList, name) {
		crypto, err := providers.GetCryptoCurrency(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
		}
		responses.RenderCryptoCurrencyResponse(req.Context(), w, crypto)
	} else {
		responses.NotFoundError(req.Context(), w)
	}
}
