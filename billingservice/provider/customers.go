package provider

import (
	"billingsystemexperiment/billingservice/domain"
	"billingsystemexperiment/billingservice/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetCustomer(customerId string) domain.Customer {

	url := "http://localhost:8080/api/v1/customers/" + customerId

	response, err := http.Get(url)
	utils.CheckError(err)

	if response.StatusCode != http.StatusOK {
		log.Fatal("Error calling the customers API: " + string(response.StatusCode))
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	utils.CheckError(err)

	var customerResponse domain.CustomerResponse
	err = json.Unmarshal(body, &customerResponse)
	utils.CheckError(err)

	return customerResponse.Customer
}
