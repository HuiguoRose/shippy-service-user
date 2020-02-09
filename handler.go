package main

import (
	"context"
	pb "github.com/HuiguoRose/shippy-service-user/proto/user"
	"log"
)

type handler struct {
	repository
}

func (s *handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := s.repository.Create(ctx, MarshalUser(req)); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (s *handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repository.Get(ctx, req.Id)
	if err != nil {
		return err
	}
	res.User = UnmarshalUser(user)
	return nil
}

func (s *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Users = UnmarshalUserCollection(users)
	return nil
}

func (s *handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	user, err := s.repository.GetByEmailAndPassword(ctx, MarshalUser(req))
	if err != nil {
		return err
	}
	log.Printf("=== user %v", user)
	res.Token = "testingabc"
	return nil
}

func (s *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
