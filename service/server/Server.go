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

	userInDb, err := dbComm.CheckLoginInDb(s.DatabaseConn, in.Login)
	if err != nil {
		return nil, err
	}
	if !userInDb {
		return &pb.UserRequest{
			Ok: false,
			Err: &pb.UserError{
				Err: "User not found",
				Id:  4,
			},
		}, nil
	}

	access, err := dbComm.CheckCorrectPassword(s.DatabaseConn, in.Login, in.Password)
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
	correct, UserError := valid.CorrectInput(in)
	if !correct {
		return &pb.UserRequest{
			Ok:  false,
			Err: UserError,
		}, nil
	}

	userInDb, err := dbComm.CheckLoginInDb(s.DatabaseConn, in.Login)
	if err != nil {
		return nil, err
	}
	if userInDb {
		return &pb.UserRequest{
			Ok: false,
			Err: &pb.UserError{
				Err: "User already exists",
				Id:  5,
			},
		}, nil
	}

	err = dbComm.CreateUserInDb(s.DatabaseConn, in.Login, in.Password)
	if err != nil {
		return nil, err
	}

	return &pb.UserRequest{
		Ok:  true,
		Err: nil,
	}, nil
}
