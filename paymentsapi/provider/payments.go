package provider

import (
	"billingsystemexperiment/paymentsapi/domain"
	"billingsystemexperiment/paymentsapi/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "ajk:password@(localhost:3306)/billing")
	utils.CheckError(err)

	err = db.Ping()
	utils.CheckError(err)

	return db
}

func GetAll(customerId string) []domain.Payment {

	db := connect()

	sql := `SELECT
					id, date, amount
					FROM payments
					WHERE customerId = ? `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	result, err := statement.Query(customerId)
	utils.CheckError(err)
	defer result.Close()

	payments := []domain.Payment{}

	payment := domain.Payment{}

	for result.Next() {
		result.Scan(&payment.Id,
			&payment.Date,
			&payment.Amount)
		payments = append(payments, payment)
	}

	return payments
}

func Post(customerId string, newPayment domain.Payment) domain.Payment {

	id := uuid.New()

	db := connect()

	sql := `INSERT INTO payments
					(id, customerId, date, amount)
					VALUES (?, ?, ?, ?) `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	_, err = statement.Exec(id, customerId, newPayment.Date, newPayment.Amount)
	utils.CheckError(err)

	payment := domain.Payment{
		Id: id.String(),
	}

	return payment
}

func GetOne(customerId string, paymentId string) (domain.Payment, error) {

	db := connect()

	sql := `SELECT
					id, date, amount
					FROM payments
					WHERE customerId = ?
					AND id = ? `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	result := statement.QueryRow(customerId, paymentId)

	payment := domain.Payment{}

	err = result.Scan(&payment.Id, &payment.Date, &payment.Amount)

	return payment, err
}

func Put(customerId string, paymentId string, newPayment domain.Payment) {

	db := connect()

	sql := `UPDATE payments
					SET date = ?, amount = ?
					WHERE customerId = ?
					AND id = ? `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	_, err = statement.Exec(newPayment.Date, newPayment.Amount, customerId, paymentId)
	utils.CheckError(err)

}

func Delete(customerId string, paymentId string) {

	db := connect()

	sql := `DELETE FROM payments
					WHERE customerId = ?
					AND id = ? `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	_, err = statement.Exec(customerId, paymentId)
	utils.CheckError(err)

}
