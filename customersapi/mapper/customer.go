package mapper

import (
	"billingsystemexperiment/customersapi/domain"
	"billingsystemexperiment/customersapi/dto"
)

func MapToDomain(sourceCustomer dto.Customer) domain.Customer {
	var destCustomer domain.Customer

	if sourceCustomer.Id.Valid {
		destCustomer.Id = sourceCustomer.Id.String
	}

	if sourceCustomer.FirstName.Valid {
		destCustomer.FirstName = sourceCustomer.FirstName.String
	}

	if sourceCustomer.LastName.Valid {
		destCustomer.LastName = sourceCustomer.LastName.String
	}

	if sourceCustomer.Street1.Valid {
		destCustomer.Street1 = sourceCustomer.Street1.String
	}

	if sourceCustomer.Street2.Valid {
		destCustomer.Street2 = sourceCustomer.Street2.String
	}

	if sourceCustomer.City.Valid {
		destCustomer.City = sourceCustomer.City.String
	}

	if sourceCustomer.State.Valid {
		destCustomer.State = sourceCustomer.State.String
	}

	if sourceCustomer.Zip.Valid {
		destCustomer.Zip = sourceCustomer.Zip.String
	}

	if sourceCustomer.CreditCardtoken.Valid {
		destCustomer.CreditCardtoken = sourceCustomer.CreditCardtoken.String
	}

	if sourceCustomer.PaymentDay.Valid {
		destCustomer.PaymentDay = int(sourceCustomer.PaymentDay.Int64)
	}

	if sourceCustomer.PaymentAmount.Valid {
		destCustomer.PaymentAmount = sourceCustomer.PaymentAmount.Float64
	}

	return destCustomer
}
