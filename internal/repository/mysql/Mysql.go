package mysql

import (
	"database/sql"
	"fmt"
	_interface "github.com/xvbnm48/bookStore/internal/repository/interface"
	"log"
	"time"
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url, port, schema, user, password string) (_interface.ReadWriter, error) {
	Mysqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", url, port, user, password, schema)
	db, err := sql.Open("mysql", Mysqlconn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &dbReadWriter{db: db}, nil
}

// Close is used for closing the sql connection
func (rw *dbReadWriter) Close() error {
	if rw.db != nil {
		if err := rw.db.Close(); err != nil {
			return err
		}
		rw.db = nil
	}

	return nil
}
