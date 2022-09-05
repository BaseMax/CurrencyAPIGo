package utils

import "github.com/itsjoniur/currency/internal/providers"

func MapKeyToSlice(m providers.Map) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
