package billing

import (
	"errors"
	"github.com/esvas/FinalProject/pkg/pars"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func New(path string) (*BillingData, error) {
	content, err := pars.ReadFile(path)
	if len(content) != 6 || err != nil {
		return nil, errors.New("ошибка входных данных billing")
	}
	return &BillingData{
		CreateCustomer: check(content[5]),
		Purchase:       check(content[4]),
		Payout:         check(content[3]),
		Recurring:      check(content[2]),
		FraudControl:   check(content[1]),
		CheckoutPage:   check(content[0]),
	}, nil
}

func check(status byte) bool {
	return status == byte('1')
}