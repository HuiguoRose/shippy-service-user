package main

import (
	"context"
	"errors"
	pb "github.com/HuiguoRose/shippy-service-user/proto/user"
	"github.com/micro/go-micro/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
)


type handler struct {
	repo         repository
	tokenService AuthAble
	PubSub       micro.Event
}

func (s *handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := s.repo.Create(ctx, MarshalUser(req)); err != nil {
		return err
	}
	res.User = req
	if err := s.PubSub.Publish(ctx, req); err != nil {
		return err
	}
	return nil
}

func (s *handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Get(ctx, req.Id)
	if err != nil {
		return err
	}
	res.User = UnmarshalUser(user)
	return nil
}

func (s *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Users = UnmarshalUserCollection(users)
	return nil
}

func (s *handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with:", req.Email, req.Password)
	//hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	return err
	//}
	//req.Password = string(hashedPass)
	user, err := s.repo.GetByEmail(ctx, req.Email)
	log.Println(user)
	if err != nil {
		return err
	}
	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}
	token, err := s.tokenService.Encode(UnmarshalUser(user))
	if err != nil {
		return err
	}
	res.Token = token
	return nil

}

func (s *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	// Decode token
	claims, err := s.tokenService.Decode(req.Token)

	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true
	return nil
}

