package email

import (
	"github.com/ferdypruis/iso3166"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

var allowedProviders = []string{
	"Gmail",
	"Yahoo",
	"Hotmail",
	"MSN",
	"Orange",
	"Comcast",
	"AOL",
	"Live",
	"RediffMail",
	"GMX",
	"Proton Mail",
	"Yandex",
	"Mail.ru",
}

type Email struct {
	Country         string
	AvgDeliveryTime int
	Provider        string
}

func New(country, provider string, avgDeliveryTime int) *Email {
	if _, err := iso3166.FromAlpha2(country); err != nil {
		return nil
	}
	if !slices.Contains(allowedProviders, provider) {
		return nil
	}
	return &Email{
		Country:         country,
		AvgDeliveryTime: avgDeliveryTime,
		Provider:        provider,
	}
}

func FromSTR(str string) *Email {
	listStr := strings.Split(str, ";")
	if len(listStr) < 3 {
		return nil
	}
	avgDeliveryTime, err := strconv.Atoi(listStr[2])
	if err != nil {
		return nil
	}
	return New(listStr[0], listStr[1], avgDeliveryTime)
}