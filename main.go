package main

import (
	pb "github.com/HuiguoRose/shippy-service-user/proto/user"
	"github.com/micro/go-micro"
	"log"
)


func main() {
	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.user"),
	)

	// Init will parse the command line flags.
	srv.Init()
	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer func() {
		_ = db.Close()
	}()
	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(User{})

	repository := &UserRepository{db}

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &handler{
		repository: repository,
	})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
