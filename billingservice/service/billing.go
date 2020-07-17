package service

import (
	"billingsystemexperiment/billingservice/provider"
)

func BillCustomer() {

	svc := provider.GetService()

	message := provider.GetMessage(svc)
	for message != nil {
		customer := provider.GetCustomer(*message.Body)

		provider.ChargeCustomer(customer)

		provider.DeleteMessage(svc, message)

		message = provider.GetMessage(svc)
	}

}
