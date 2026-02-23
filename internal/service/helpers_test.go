package service

import (
	"strings"
	"testing"
)

func TestIsValidUrl(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want bool
	}{
		{"Valid HTTPS", "https://google.com", true},
		{"Valid HTTP", "http://ya.ru", true},
		{"Valid with path", "https://github.com/LeezyWannaFall", true},
		{"Missing scheme", "google.com", false},
		{"Missing host", "https://", false},
		{"Empty string", "", false},
		{"Random text", "just-random-text", false},
		{"Spaces", "https:// google.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidUrl(tt.url); got != tt.want {
				t.Errorf("isValidUrl(%s) = %v, want %v", tt.url, got, tt.want)
			}
		})
	}
}

func TestGenerateShortLink(t *testing.T) {
	link := GenerateShortLink()
	if len(link) != 10 {
		t.Errorf("Expected length 10, got %d", len(link))
	}

	for _, char := range link {
		if !strings.ContainsRune(alphabet, char) {
			t.Errorf("Character %c not found in alphabet", char)
		}
	}
}