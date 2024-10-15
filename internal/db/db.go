package db

import (
	"github.com/go-pg/pg/v9"
	"log"
)

var DB *pg.DB

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "1234321",
		Addr:     "localhost:5432",
		Database: "project",
		OnConnect: func(conn *pg.Conn) error {
			return createProductsTable(conn)
		},
	} //параметри для підключення до бази даних
	log.Println("Connecting to database")
	var db *pg.DB = pg.Connect(opts)
	log.Println("Connected to database")
	return db
}

func createProductsTable(conn *pg.Conn) error {
	_, err := conn.Exec(`CREATE TABLE IF NOT EXISTS products (
		id serial PRIMARY KEY,
		title varchar(50) NOT NULL,
		price integer NOT NULL
	)`)
	if err != nil {
		log.Println("Error while creating table products:", err)
		return err
	}

	log.Println("Table 'products' created or already exists")
	return nil
}
