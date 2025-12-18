package app

import (
	"database/sql"
	"fmt"
)

func NewDB(userDB, passDB, hostDB, portDB, nameDB string) *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", userDB, passDB, hostDB, portDB, nameDB)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db

}
