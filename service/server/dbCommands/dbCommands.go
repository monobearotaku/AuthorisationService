package dbCommands

import (
	"database/sql"
	"fmt"
	pb "github.com/ash0tych/gRPC_MusicService/proto/compiled/authorization"
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

func CheckUserInDb(db *sql.DB, login, password string) (bool, *pb.UserError, error) {
	query := `SELECT id, password FROM "Users" WHERE login=$1`
	rows, err := db.Query(query, login)
	if err != nil {
		return false, nil, err
	}

	if rows.Next() {
		var dbPassword string
		var dbId string

		err = rows.Scan(&dbId, &dbPassword)
		if err != nil {
			return false, nil, err
		}

		if dbPassword == password {
			return true, nil, nil
		}

		if dbPassword != "" {
			panic(nil)
			return false, &pb.UserError{
				Err: "Wrong password",
				Id:  3,
			}, nil
		}
	}
	return false, &pb.UserError{
		Err: "User not found",
		Id:  4,
	}, nil
}
