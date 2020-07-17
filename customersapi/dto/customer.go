package dto

import "database/sql"

type Customer struct {
	Id              sql.NullString  `json:"id,omitempty"`
	FirstName       sql.NullString  `json:"firstName,omitempty"`
	LastName        sql.NullString  `json:"lastName,omitempty"`
	Street1         sql.NullString  `json:"street1,omitempty"`
	Street2         sql.NullString  `json:"street2,omitempty"`
	City            sql.NullString  `json:"city,omitempty"`
	State           sql.NullString  `json:"state,omitempty"`
	Zip             sql.NullString  `json:"zip,omitempty"`
	CreditCardtoken sql.NullString  `json:"creditCardtoken,omitempty"`
	PaymentDay      sql.NullInt64   `json:"paymentDay,omitempty"`
	PaymentAmount   sql.NullFloat64 `json:"paymentAmount,omitempty"`
}
