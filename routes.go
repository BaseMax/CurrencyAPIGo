package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
		w.Write([]byte("currencies"))
		return
	}

	name = strings.ToLower(name)

	golds := []string{"bitcoin", "ounce", "mithqal", "gram"}
	coins := []string{"azadi1", "azadi1_2", "azadi_4", "emami", "azadi1g"}

	if slices.Contains(golds, name) {
		g, err := GetGold(name)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("not found"))
			return
		}
		json.NewEncoder(w).Encode(g)
		return
	} else if slices.Contains(coins, name) {
		c, err := GetCoin(name)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("not found"))
			return
		}
		json.NewEncoder(w).Encode(c)
		return
	}

	uc, err := GetCurrency(name)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("not found"))
		return
	}
	json.NewEncoder(w).Encode(uc)

}
