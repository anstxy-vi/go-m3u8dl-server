package utils

import (
	"golang.org/x/text/message"
)

func ToCurrency(amount float64) string {
	p := message.NewPrinter(message.MatchLanguage("en"))

	return p.Sprintf("$%.2f", amount)
}
