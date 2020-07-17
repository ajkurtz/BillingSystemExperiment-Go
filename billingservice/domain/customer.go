package domain

type Customer struct {
	Id              string  `json:"id,omitempty"`
	FirstName       string  `json:"firstName,omitempty"`
	LastName        string  `json:"lastName,omitempty"`
	Street1         string  `json:"street1,omitempty"`
	Street2         string  `json:"street2,omitempty"`
	City            string  `json:"city,omitempty"`
	State           string  `json:"state,omitempty"`
	Zip             string  `json:"zip,omitempty"`
	CreditCardtoken string  `json:"creditCardtoken,omitempty"`
	PaymentDay      int     `json:"paymentDay,omitempty"`
	PaymentAmount   float64 `json:"paymentAmount,omitempty"`
}
