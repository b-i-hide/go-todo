package db

import "database/sql"

var db *sql.DB

func DbInit() (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", "root@/go_todo_development")
	return db, err
}

func DbClose() {
	if db != nil {
		db.Close()
	}
}

func DbConn() *sql.DB {
	return db
}
