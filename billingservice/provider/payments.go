package provider

import (
	"billingsystemexperiment/billingservice/domain"
	"billingsystemexperiment/billingservice/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func ChargeCustomer(customer domain.Customer) {

	year, month, day := time.Now().Date()

	date := fmt.Sprintf("%d-%d-%d", year, month, day)

	payment := domain.Payment{
		Amount:          customer.PaymentAmount,
		Date:            date,
		CreditCardToken: customer.CreditCardtoken,
	}

	paymentRequest := domain.PaymentRequest{Payment: payment}

	requestBody, err := json.Marshal(paymentRequest)

	url := "http://localhost:8081/api/v1/customers/" + customer.Id + "/payments"

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	utils.CheckError(err)

	if response.StatusCode != http.StatusCreated {
		log.Fatal("Error calling the payments API: " + string(response.StatusCode))
	}

}
