package service

import (
	"math/rand/v2"
	"net/url"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func GenerateShortLink() string {
	result := make([]byte, 10)

	for i := 0; i < 10; i++ {
		result[i] = alphabet[rand.IntN(len(alphabet))]
	}
 	
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