package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectDatabase() *sql.DB {
	conexao := "user=postgres dbname=lojinha_online password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
