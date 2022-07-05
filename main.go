package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/garfnadeck/go-finance/api"
	db "github.com/garfnadeck/go-finance/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main(){
	err := godotenv.Load()
	if err !=nil{
		log.Fatal("error loading .env file")
	}

	dbDriver:= os.Getenv("DB_DRIVER")
	dbSource:= os.Getenv("DB_SOURCE")
	serverAddress:= os.Getenv("DB_SERVER")
	

	conn, err := sql.Open(dbDriver, dbSource)
	if err !=nil{
		log.Fatal("cannot connect to the db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err !=nil{
				log.Fatal("cannot connect to the API: ", err)
	}
}