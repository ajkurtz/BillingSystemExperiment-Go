package provider

import (
	"billingsystemexperiment/customersapi/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "ajk:password@(localhost:3306)/billing")
	utils.CheckError(err)

	err = db.Ping()
	utils.CheckError(err)

	return db
}

func GetToBillCustomers(paymentDay int) []string {

	db := connect()

	sql := `SELECT
					id
					FROM customers
					WHERE paymentDay = ? `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	result, err := statement.Query(paymentDay)
	utils.CheckError(err)
	defer result.Close()

	customers := []string{}
	var customer string

	for result.Next() {
		result.Scan(&customer)
		customers = append(customers, customer)
	}

	return customers
}
