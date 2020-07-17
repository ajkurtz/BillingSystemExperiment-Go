package domain

type Payment struct {
	Id              string  `json:"id,omitempty"`
	Date            string  `json:"date,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	CreditCardToken string  `json:"creditCardToken,omitempty"`
}
