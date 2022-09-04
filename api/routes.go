package api

import (
	"net/http"
	"strings"

	"github.com/itsjoniur/currency/internal/providers"
	"github.com/itsjoniur/currency/internal/responses"
	"github.com/itsjoniur/currency/pkg/utils"

	"golang.org/x/exp/slices"
)

func RootHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("documentation"))
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

	if slices.Contains(utils.MapKeyToSlice(providers.Currencies), name) {
		currency, err := providers.GetCurrency(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
		}
		responses.RenderCurrencyResponse(req.Context(), w, currency)
		return
	} else if slices.Contains(utils.MapKeyToSlice(providers.Coins), name) {
		coin, err := providers.GetCoin(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
			return
		}
		responses.RenderCoinResponse(req.Context(), w, coin)
		return
	} else if slices.Contains(utils.MapKeyToSlice(providers.Golds), name) {
		gold, err := providers.GetGold(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
			return
		}
		responses.RenderGoldResponse(req.Context(), w, gold)
	} else {
		responses.NotFoundError(req.Context(), w)
	}

}
