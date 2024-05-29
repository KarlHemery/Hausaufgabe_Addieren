package services

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func NewDB(host, user, password, dbname string) *DB {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	return &DB{conn: db}
}

func (db *DB) CalculateSum(num1, num2 float64) (float64, error) {
	var sum float64
	query := "SELECT $1::numeric + $2::numeric"
	err := db.conn.QueryRow(query, num1, num2).Scan(&sum)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %v", err)
	}

	return sum, nil
}

func (db *DB) Close() {
	db.conn.Close()
}
