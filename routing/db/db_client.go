package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Connection struct {
	*pgxpool.Pool
}

var connection *Connection

func InitDB() error {

	if connection != nil {
		return nil
	}

	config, _ := pgxpool.ParseConfig(os.Getenv("PGCON"))

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	connection = &Connection{pool}

	if err != nil {
		log.Fatalln("Unable to connect to database \n", err)
		return err
	}

	log.Println("DB connection established")


	return nil
}

func GetConnection() *Connection {
	if connection == nil {
		log.Println("no active db connection... setting up a new one")
		InitDB()
	}
	return connection
}
