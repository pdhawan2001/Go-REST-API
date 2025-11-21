package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/pdhawan2001/Go-REST-API/cmd/api"
	"github.com/pdhawan2001/Go-REST-API/db"
)

// func (receiver) MethodName(params) returnTypes

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "asd",
		Addr:                 "127.0.1:3306",
		DBName:               "GO-API",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
