package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
)

func isMorseCode(text string) bool {
	allowedChar := ".- "

	for _, char := range text {
		if !strings.ContainsRune(allowedChar, char) {
			return false
		}
	}
	return true
}

func Convert(text string) string {

	if isMorseCode(text) {
		return morse.ToText(text)
	}

	return morse.ToMorse(text)
}
