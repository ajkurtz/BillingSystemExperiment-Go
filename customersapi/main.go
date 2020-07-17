package main

import (
	"billingsystemexperiment/customersapi/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/customers", getHandler).Methods("GET")
	r.HandleFunc("/api/v1/customers", postHandler).Methods("POST")
	r.HandleFunc("/api/v1/customers/{customerId}", getOneHandler).Methods("GET")
	r.HandleFunc("/api/v1/customers/{customerId}", putHandler).Methods("PUT")
	r.HandleFunc("/api/v1/customers/{customerId}", deleteHandler).Methods("DELETE")
	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func getHandler(writer http.ResponseWriter, request *http.Request) {
	service.GetAll(writer, request)
}

func postHandler(writer http.ResponseWriter, request *http.Request) {
	service.Post(writer, request)
}

func getOneHandler(writer http.ResponseWriter, request *http.Request) {
	service.GetOne(writer, request)
}

func putHandler(writer http.ResponseWriter, request *http.Request) {
	service.Put(writer, request)
}

func deleteHandler(writer http.ResponseWriter, request *http.Request) {
	service.Delete(writer, request)
}
