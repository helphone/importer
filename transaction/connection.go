package transaction

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Load the drive of postgres
)

// Connection struct will hold the context of the transaction
type Connection struct {
	Tx                     *sql.Tx
	CloseConnectionChannel chan int
}

// CreateConnection will create a connection struct
// to do the all import
func CreateConnection() (*Connection, error) {
	hostname := os.Getenv("DB_PORT_5432_TCP_ADDR")
	port := os.Getenv("DB_PORT_5432_TCP_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		hostname, port, username, password, "helphone")

	database, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	tx, err := database.Begin()
	if err != nil {
		return nil, err
	}

	conn := &Connection{
		Tx: tx,
		CloseConnectionChannel: make(chan int, 1),
	}

	go func() {
		<-conn.CloseConnectionChannel
		database.Close()
	}()

	return conn, err
}

// Finish is the last function to call with giving error provided by
// every transaction to know if everything goes fine
func (c Connection) Finish(err error) {
	switch err {
	case nil:
		err = c.Tx.Commit()
	default:
		c.Tx.Rollback()
	}
	c.CloseConnectionChannel <- 1
}
