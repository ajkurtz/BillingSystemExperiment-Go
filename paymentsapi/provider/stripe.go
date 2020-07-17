package provider

import (
	"fmt"
	"billingsystemexperiment/paymentsapi/domain"
	"billingsystemexperiment/paymentsapi/utils"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func Charge(payment domain.Payment) error {

	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"

	token := utils.Decrypt(payment.CreditCardToken)

	cents := int64(payment.Amount * 100.0)

	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(cents),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Source:   &stripe.SourceParams{Token: stripe.String(token)},
	}
	_, err := charge.New(params)
	fmt.Println(err) // Ignore the error for this example since we don't have a valid token

	return nil
}
