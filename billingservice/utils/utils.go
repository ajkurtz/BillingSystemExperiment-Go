package utils

import (
	"log"
)

const key = "12345678901234567890123456789012"

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
