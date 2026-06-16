package main

import (
	"database/sql"
	"log"

	"github.com/JaguarsCodehub/golang-ecom/cmd/api"
	"github.com/JaguarsCodehub/golang-ecom/config"
	"github.com/JaguarsCodehub/golang-ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	database, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer database.Close()

	initStorage(database)

	addr := ":" + config.Envs.Port
	server := api.NewAPIServer(addr, database)
	if err := server.Run(); err != nil {
		log.Fatal("Error in main:", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	log.Println("DB: Connected successfully")
}
