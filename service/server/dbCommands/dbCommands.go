package dbCommands

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectToDb() (*sql.DB, error) {
	conn := fmt.Sprintf("host=localhost port=8080 user=postgres password=12345678 dbname=UserAuthorization sslmode=disable")
	bufferDb, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = bufferDb.Ping()
	if err != nil {
		return nil, err
	}

	return bufferDb, nil
}

func CloseDb(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

func CheckCorrectPassword(db *sql.DB, login, password string) (bool, error) {
	query := `SELECT id, password FROM "Users" WHERE login=$1`
	rows, err := db.Query(query, login)
	if err != nil {
		return false, err
	}

	var dbId string
	var dbPassword string

	rows.Next()
	err = rows.Scan(&dbId, &dbPassword)
	if err != nil {
		return false, err
	}

	if dbPassword != password {
		return false, nil
	}

	return true, nil
}

func CheckLoginInDb(db *sql.DB, login string) (bool, error) {
	query := `SELECT id FROM "Users" WHERE login=$1`
	rows, err := db.Query(query, login)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func CreateUserInDb(db *sql.DB, login, password string) error {
	query := `INSERT into "Users" (login, password) values ($1, $2)`
	_, err := db.Exec(query, login, password)
	if err != nil {
		return err
	}
	return nil
}
