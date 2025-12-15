package db

import (
	"Ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
	if !cnf.EnableSSLMODE {
		connString += " sslmode=disable"
	}
	return connString
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbsource := GetConnectionString(cnf)

	dbcon, err := sqlx.Connect("postgres", dbsource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbcon, nil
}
