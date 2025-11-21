package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/pdhawan2001/Go-REST-API/cmd/api"
	"github.com/pdhawan2001/Go-REST-API/config"
	"github.com/pdhawan2001/Go-REST-API/db"
)

// func (receiver) MethodName(params) returnTypes

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
