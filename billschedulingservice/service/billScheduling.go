package service

import (
	"billingsystemexperiment/billschedulingservice/provider"
	"time"
)

func Schedule() {

	customers := provider.GetToBillCustomers(time.Now().Day())

	svc := provider.GetService()

	for _, customer := range customers {
		provider.SendMessage(svc, customer)
		time.Sleep(1 * time.Second) // Sending messages too fast can cause them to be lost
	}

}
