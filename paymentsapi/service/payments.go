package service

import (
	"billingsystemexperiment/paymentsapi/domain"
	"billingsystemexperiment/paymentsapi/provider"
	"billingsystemexperiment/paymentsapi/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customerId"]

	if provider.CustomerIsValid(customerId) {
		payments := provider.GetAll(customerId)

		paymentsResponse := domain.PaymentsResponse{
			Payments: payments,
		}

		output, err := json.Marshal(&paymentsResponse)
		utils.CheckError(err)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, string(output))
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func Post(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customerId"]

	if provider.CustomerIsValid(customerId) {
		body, err := ioutil.ReadAll(request.Body)
		utils.CheckError(err)

		var paymentRequest domain.PaymentRequest
		json.Unmarshal(body, &paymentRequest)

		err = provider.Charge(paymentRequest.Payment)
		utils.CheckError(err)

		payment := provider.Post(customerId, paymentRequest.Payment)

		paymentResponse := domain.PaymentResponse{
			Payment: payment,
		}

		output, err := json.Marshal(&paymentResponse)
		utils.CheckError(err)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		fmt.Fprintf(writer, string(output))
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}

}

func GetOne(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customerId"]
	paymentId := vars["paymentId"]

	if provider.CustomerIsValid(customerId) {

		payment, err := provider.GetOne(customerId, paymentId)
		if err == sql.ErrNoRows {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			utils.CheckError(err)

			paymentResponse := domain.PaymentResponse{
				Payment: payment,
			}

			output, err := json.Marshal(&paymentResponse)
			utils.CheckError(err)

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			fmt.Fprintf(writer, string(output))
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func Put(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customerId"]
	paymentId := vars["paymentId"]

	if provider.CustomerIsValid(customerId) {
		body, err := ioutil.ReadAll(request.Body)
		utils.CheckError(err)

		var paymentRequest domain.PaymentRequest
		json.Unmarshal(body, &paymentRequest)

		_, err = provider.GetOne(customerId, paymentId)
		if err == sql.ErrNoRows {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			utils.CheckError(err)

			provider.Put(customerId, paymentId, paymentRequest.Payment)

			writer.WriteHeader(http.StatusNoContent)
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customerId"]
	paymentId := vars["paymentId"]

	if provider.CustomerIsValid(customerId) {
		_, err := provider.GetOne(customerId, paymentId)
		if err == sql.ErrNoRows {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			utils.CheckError(err)

			provider.Delete(customerId, paymentId)

			writer.WriteHeader(http.StatusNoContent)
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}
