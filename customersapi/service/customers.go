package service

import (
	"billingsystemexperiment/customersapi/domain"
	"billingsystemexperiment/customersapi/provider"
	"billingsystemexperiment/customersapi/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {

	customers := provider.GetAll()

	customersResponse := domain.CustomersResponse{
		Customers: customers,
	}

	output, err := json.Marshal(&customersResponse)
	utils.CheckError(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, string(output))

}

func Post(writer http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)
	utils.CheckError(err)

	var customerRequest domain.CustomerRequest
	json.Unmarshal(body, &customerRequest)

	customer := provider.Post(customerRequest.Customer)

	customerResponse := domain.CustomerResponse{
		Customer: customer,
	}

	output, err := json.Marshal(&customerResponse)
	utils.CheckError(err)

	var o = string(output)
	fmt.Println(o)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, string(output))

}

func GetOne(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customerId"]

	customer, err := provider.GetOne(customerId)
	if err == sql.ErrNoRows {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		utils.CheckError(err)

		customerResponse := domain.CustomerResponse{
			Customer: customer,
		}

		output, err := json.Marshal(&customerResponse)
		utils.CheckError(err)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, string(output))
	}

}

func Put(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customerId"]

	body, err := ioutil.ReadAll(request.Body)
	utils.CheckError(err)

	var customerRequest domain.CustomerRequest
	json.Unmarshal(body, &customerRequest)

	_, err = provider.GetOne(customerId)
	if err == sql.ErrNoRows {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		utils.CheckError(err)

		provider.Put(customerId, customerRequest.Customer)

		writer.WriteHeader(http.StatusNoContent)
	}
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customerId"]

	_, err := provider.GetOne(customerId)
	if err == sql.ErrNoRows {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		utils.CheckError(err)

		provider.Delete(customerId)

		writer.WriteHeader(http.StatusNoContent)
	}
}
