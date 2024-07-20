package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dsn string) {
    var err error
    db, err = sql.Open("sqlite3", dsn)
    if err != nil {
        log.Fatal(err)
    }

    createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        email TEXT UNIQUE
    );`
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
}

func GetDB() *sql.DB {
    return db
}

func Close() {
    db.Close()
}
