package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

func InitDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("sqlite3", "events.db")
		if err != nil {
			panic("Can't connect to DB")
		}

		if err = db.Ping(); err != nil {
			panic("Can't verify connection to DB")
		}

		createEventTable()
	})

}

func GetDB() *sql.DB {
	if db == nil {
		InitDB()
	}
	return db
}

func createEventTable() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                description TEXT NOT NULL,
                location TEXT NOT NULL,
                date_time DATETIME NOT NULL,
                user_id INTEGER
        );`

	_, err := db.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Can't create table: %v", err)
	}
}
