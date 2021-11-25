package server

import (
	"context"
	"database/sql"
	pb "github.com/ash0tych/gRPC_1/proto/compiled/authorization"
	dbComm "github.com/ash0tych/gRPC_1/service/server/dbCommands"
	valid "github.com/ash0tych/gRPC_1/service/server/validation"
	_ "github.com/lib/pq"
	"log"
)

type Server struct {
	DatabaseConn *sql.DB
}

func (s *Server) Start() {
	err := dbComm.ConnectToDb(s.DatabaseConn)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Close() error {
	err := dbComm.CloseDb(s.DatabaseConn)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s *Server) IdentifyUser(ctx context.Context, in *pb.UserData) (*pb.UserRequest, error) {
	err, UserError := valid.CorrectInput(in)
	if !err {
		return &pb.UserRequest{
			Ok:  err,
			Err: UserError,
		}, nil
	}

	return &pb.UserRequest{
		Ok:  true,
		Err: nil,
	}, nil
}
