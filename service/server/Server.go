package server

import (
	"context"
	"database/sql"
	pb "github.com/ash0tych/gRPC_MusicService/proto/compiled/authorization"
	dbComm "github.com/ash0tych/gRPC_MusicService/service/server/dbCommands"
	valid "github.com/ash0tych/gRPC_MusicService/service/server/validation"
	_ "github.com/lib/pq"
)

type Server struct {
	DatabaseConn *sql.DB
	pb.UnimplementedUserServiceServer
}

func (s *Server) Start() error {
	err := dbComm.ConnectToDb(s.DatabaseConn)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Close() {
	_ = dbComm.CloseDb(s.DatabaseConn)
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

func (s *Server) CreateUser(ctx context.Context, in *pb.UserData) (*pb.UserRequest, error) {
	return &pb.UserRequest{
		Ok:  true,
		Err: nil,
	}, nil
}
