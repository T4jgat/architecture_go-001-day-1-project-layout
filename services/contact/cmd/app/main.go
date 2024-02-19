package main

import (
	"architecture_go/pkg/store/postgres"
	"fmt"
	"log"
)

func main() {

	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "1423"
	dbname := "postgres"

	db, err := postgres.Connect(host, port, user, password, dbname)
	if err != nil {
		log.Fatal(err)
		return
	}

	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var barcode int
		var firstname string
		err := rows.Scan(&barcode, &firstname)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("ID:", barcode, "Name:", firstname)
	}

	// Закрытие подключения
	db.Close()

	fmt.Println("Hello World!")
}
