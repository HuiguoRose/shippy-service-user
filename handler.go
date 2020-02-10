package main

import (
	"context"
	"encoding/json"
	"errors"
	pb "github.com/HuiguoRose/shippy-service-user/proto/user"
	"github.com/micro/go-micro/v2/broker"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const topic = "user.created"

type handler struct {
	repo repository
	tokenService AuthAble
	PubSub broker.Broker
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
	if err := s.publishEvent(req); err != nil {
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
	user, err := s.repo.GetByEmail(ctx, MarshalUser(req))
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

func (s *handler) publishEvent(user *pb.User) error {
	// Marshal to JSON string
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Create a broker message
	msg := &broker.Message{
		Header: map[string]string{
			"id": user.Id,
		},
		Body: body,
	}

	// Publish message to broker
	if err := s.PubSub.Publish(topic, msg); err != nil {
		log.Printf("[pub] failed: %v", err)
	}

	return nil
}