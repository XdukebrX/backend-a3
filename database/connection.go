package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() (*sql.DB, error) {
	stringConnect := "root:@/a3-clavison?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConnect)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to DataBase is open ")
	return db, nil
}
