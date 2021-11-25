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
	var err error
	s.DatabaseConn, err = dbComm.ConnectToDb()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Close() {
	_ = dbComm.CloseDb(s.DatabaseConn)
}

func (s *Server) IdentifyUser(ctx context.Context, in *pb.UserData) (*pb.UserRequest, error) {
	flag, UserError := valid.CorrectInput(in)
	if !flag {
		return &pb.UserRequest{
			Ok:  flag,
			Err: UserError,
		}, nil
	}

	access, UserError, err := dbComm.CheckUserInDb(s.DatabaseConn, in.Login, in.Password)
	if err != nil {
		return nil, err
	}

	if !access {
		return &pb.UserRequest{
			Ok: false,
			Err: &pb.UserError{
				Err: "Wrong password",
				Id:  3,
			},
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
