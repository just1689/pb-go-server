package mysql

import "database/sql"

const DefaultConnection = "root:toor@tcp(127.0.0.1:3306)/parabible_test"

func Open(connection string) (*sql.DB, error) {
	return sql.Open("mysql", connection)
}

func Connect(db *sql.DB) error {
	return db.Ping()
}
