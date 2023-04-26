package sms

import (
	"github.com/ferdypruis/iso3166"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

var allowedProviders = []string{"Topolo", "Rond", "Kildy"}

type SMS struct {
	Country     string `json:"country"`
	Bandwidth   int    `json:"bandwidth"`
	AvgRespTime int    `json:"avg_resp_time"`
	Provider    string `json:"provider"`
}

func New(country, provider string, bandwidth, avgRespTime int) *SMS {
	c, err := iso3166.FromAlpha2(country)
	if err != nil {
		return nil
	}
	if bandwidth < 0 && bandwidth > 100 {
		return nil
	}
	if !slices.Contains(allowedProviders, provider) {
		return nil
	}
	return &SMS{
		Country:     c.Name(),
		Bandwidth:   bandwidth,
		AvgRespTime: avgRespTime,
		Provider:    provider,
	}
}

func FromSTR(str string) *SMS {
	listStr := strings.Split(str, ";")
	if len(listStr) < 4 {
		return nil
	}
	bandwidth, err := strconv.Atoi(listStr[1])
	if err != nil {
		return nil
	}
	avgRespTime, err := strconv.Atoi(listStr[2])
	if err != nil {
		return nil
	}
	return New(listStr[0], listStr[3], bandwidth, avgRespTime)
}