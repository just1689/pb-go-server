package mysql

import "database/sql"

func Open(connection string) (*sql.DB, error) {
	return sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/parabible_test")
}

func Connect(db *sql.DB) error {
	return db.Ping()
}
