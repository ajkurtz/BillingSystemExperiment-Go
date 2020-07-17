package provider

import (
	"billingsystemexperiment/customersapi/domain"
	"billingsystemexperiment/customersapi/dto"
	"billingsystemexperiment/customersapi/mapper"
	"billingsystemexperiment/customersapi/utils"
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

func GetAll() []domain.Customer {

	db := connect()

	sql := `SELECT
					id, firstName, lastName, street1, street2, city, state, zip, creditCardToken, paymentAmount, paymentDay
					FROM customers `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	result, err := statement.Query()
	utils.CheckError(err)
	defer result.Close()

	customers := []domain.Customer{}
	customer := dto.Customer{}

	for result.Next() {
		result.Scan(&customer.Id,
			&customer.FirstName,
			&customer.LastName,
			&customer.Street1,
			&customer.Street2,
			&customer.City,
			&customer.State,
			&customer.Zip,
			&customer.CreditCardtoken,
			&customer.PaymentAmount,
			&customer.PaymentDay)
		customers = append(customers, mapper.MapToDomain(customer))
	}

	return customers
}

func Post(newCustomer domain.Customer) domain.Customer {

	id := uuid.New()

	encryptedToken := utils.Encrypt(newCustomer.CreditCardtoken)

	db := connect()

	sql := `INSERT INTO customers
					(id, firstName, lastName, street1, street2, city, state, zip, creditCardToken, paymentAmount, paymentDay)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	_, err = statement.Exec(id, newCustomer.FirstName, newCustomer.LastName, newCustomer.Street1, newCustomer.Street2,
		newCustomer.City, newCustomer.State, newCustomer.Zip, encryptedToken, newCustomer.PaymentAmount, newCustomer.PaymentDay)
	utils.CheckError(err)

	customer := domain.Customer{
		Id: id.String(),
	}

	return customer
}

func GetOne(customerId string) (domain.Customer, error) {

	db := connect()

	sql := `SELECT
					id, firstName, lastName, street1, street2, city, state, zip, creditCardToken, paymentAmount, paymentDay
					FROM customers
					WHERE id = ? `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	result := statement.QueryRow(customerId)

	customer := dto.Customer{}

	err = result.Scan(&customer.Id,
		&customer.FirstName,
		&customer.LastName,
		&customer.Street1,
		&customer.Street2,
		&customer.City,
		&customer.State,
		&customer.Zip,
		&customer.CreditCardtoken,
		&customer.PaymentAmount,
		&customer.PaymentDay)

	return mapper.MapToDomain(customer), err
}

func Put(customerId string, newCustomer domain.Customer) {

	encryptedToken := utils.Encrypt(newCustomer.CreditCardtoken)

	db := connect()

	sql := `UPDATE customers
          SET firstName = ?,
          lastName = ?,
          street1 = ?,
          street2 = ?,
          city = ?,
          state = ?,
          zip = ?,
          creditCardToken = ?,
          paymentDay = ?,
          paymentAmount = ?
          WHERE id = ? `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	_, err = statement.Exec(newCustomer.FirstName, newCustomer.LastName, newCustomer.Street1, newCustomer.Street2,
		newCustomer.City, newCustomer.State, newCustomer.Zip, encryptedToken, newCustomer.PaymentDay, newCustomer.PaymentAmount, customerId)
	utils.CheckError(err)

}

func Delete(customerId string) {

	db := connect()

	sql := `DELETE FROM customers
          WHERE id = ? `

	statement, err := db.Prepare(sql)
	utils.CheckError(err)
	defer statement.Close()

	_, err = statement.Exec(customerId)
	utils.CheckError(err)

}
