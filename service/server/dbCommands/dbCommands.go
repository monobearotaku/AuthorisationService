package dbCommands

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectToDb(db *sql.DB) error {
	var err error
	conn := fmt.Sprintf("host=localhost port=8080 user=postgres password=z]/q1937 dbname=UserAuthorization sslmode=disable")
	db, err = sql.Open("postgres", conn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func CloseDb(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

func CheckUserInDb(db *sql.DB, login, password string) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM Users WHERE login=%v AND password=%v", login, password)
	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}
	return rows.Next(), nil
}
