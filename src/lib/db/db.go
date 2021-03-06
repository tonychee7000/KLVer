/*
Combine go native sql module with mysql-driver
*/
package db

import "database/sql"

// Forked from https://github.com/go-sql-driver/mysql
import _ "lib/db/go-mysql"

//Connect to MySQL server.
func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	return db, err
}
