package domain

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	

	FindAllSql := "select * from customers"

	rows, err := d.client.Query(FindAllSql)

	if err != nil {
		log.Println("Error while quering customer table")
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scaning customer table")
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}


func NewCostumerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}