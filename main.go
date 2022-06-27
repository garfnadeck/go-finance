package main


import(
	"database/sql"
	"log"

	"github.com/garfnadeck/go-finance/api"
	_ "github.com/lib/pq"
	db "github.com/garfnadeck/go-finance/db/sqlc"

)

const (
	dbDriver = "postgres"
	dbSource ="postgres://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main(){
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