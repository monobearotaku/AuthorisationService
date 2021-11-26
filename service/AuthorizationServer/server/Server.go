package server

import (
	"context"
	"database/sql"
	pb "github.com/ash0tych/gRPC_MusicService/proto/compiled/authorization"
	dbComm "github.com/ash0tych/gRPC_MusicService/service/AuthorizationServer/dbCommands"
	valid "github.com/ash0tych/gRPC_MusicService/service/AuthorizationServer/validation"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"time"
)

const signKey string = "Hello World"

type Server struct {
	DatabaseConn *sql.DB
	pb.UnimplementedUserServiceServer
}

type Claims struct {
	jwt.StandardClaims
	Email string
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

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: in.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signKey))

	if err != nil {
		return nil, err
	}

	return &pb.UserRequest{
		Ok:    true,
		Err:   nil,
		Token: tokenString,
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

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: in.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signKey))

	return &pb.UserRequest{
		Ok:    true,
		Err:   nil,
		Token: tokenString,
	}, nil
}
