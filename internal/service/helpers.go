package service

import (
	"math/rand/v2"
	"net/url"
)

const (
	number = "0123456789"
	lower = "abcdefghijklmnopqrstuvwxyz"
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	all = number + lower + upper
)

func ShortLinkGenerator() string {
	result := make([]byte, 10)

	result[0] = number[rand.IntN(len(number))]
	result[1] = lower[rand.IntN(len(lower))]
	result[2] = upper[rand.IntN(len(upper))]
	result[3] = '_'

	for i := 4; i < 10; i++ {
		result[i] = all[rand.IntN(len(all))]
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
 	
	return string(result)
}

func isValidUrl(parseUrl string) bool {
	u, err := url.ParseRequestURI(parseUrl)
	if err != nil {
		return false
	}

	if u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}