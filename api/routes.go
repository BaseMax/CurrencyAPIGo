package api

import (
	"encoding/json"
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
	/*
		NOTE: This function not completed and it is just created to show expected result to us
		TODO(itsjoniur): fix map[any]any in MapKeyToSlice function
		Why we need MapKeyToSlice function?
		Answer:
			we need this function to turn keys of a map as a slice
			then check if the slice contains given query's value or not.
			so we can know what type of currency want to get and we'll return it

	*/
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
	} else {
		gold, err := providers.GetGold(req.Context(), name)
		if err != nil {
			responses.NotFoundError(req.Context(), w)
			return
		}
		json.NewEncoder(w).Encode(gold)
	}

}
