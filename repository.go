package main

import (
	"context"
	pb "github.com/HuiguoRose/shippy-service-user/proto/user"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Company  string ` json:"company,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func UnmarshalUser(user *User) *pb.User {
	return &pb.User{
		Id:       user.Id,
		Name:     user.Name,
		Company:  user.Company,
		Email:    user.Email,
		Password: user.Password,
	}
}

func MarshalUser(user *pb.User) *User {
	return &User{
		Id:       user.Id,
		Name:     user.Name,
		Company:  user.Company,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	return scope.SetColumn("Id", id.String())
}

type Users []*User

func UnmarshalUserCollection(users Users) []*pb.User {
	collection := make([]*pb.User, 0)
	for _, user := range users {
		collection = append(collection, UnmarshalUser(user))
	}
	return collection
}

type repository interface {
	Create(ctx context.Context, user *User) error
	Get(ctx context.Context, Id string) (*User, error)
	GetAll(ctx context.Context) (Users, error)
	GetByEmailAndPassword(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, user *User) (*User, error)
}
type UserRepository struct {
	db *gorm.DB
}

// MongoRepository implementation
//type MongoRepository struct {
//	collection *mongo.Collection
//}

// Create -
func (repository *UserRepository) Create(ctx context.Context, user *User) error {
	if err := repository.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Get -
func (repository *UserRepository) Get(ctx context.Context, Id string) (*User, error) {
	var user = &User{}
	user.Id = Id
	if err := repository.db.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Get -
func (repository *UserRepository) GetAll(ctx context.Context) (Users, error) {
	var users Users
	if err := repository.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepository) GetByEmailAndPassword(ctx context.Context, user *User) (*User, error) {
	if err := repository.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) GetByEmail(ctx context.Context, user *User) (*User, error) {
	if err := repository.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
