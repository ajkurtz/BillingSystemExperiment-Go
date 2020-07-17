package provider

import (
	"billingsystemexperiment/paymentsapi/utils"
	"net/http"
)

func CustomerIsValid(customerId string) bool {

	url := "http://localhost:8080/api/v1/customers/" + customerId

	response, err := http.Get(url)
	utils.CheckError(err)

	if response.StatusCode != http.StatusOK {
		return false
	}

	return true

}
