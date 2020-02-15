package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
	"post-service/pb"
	"post-service/server"
	"post-service/store"
)

func main() {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "root"
		dbname   = "crud_golang"
	)

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()


	PostsServiceServer := server.NewPostServer(store.NewPostStore(db))
	pb.RegisterPostsServiceHandlerServer(ctx, mux, PostsServiceServer)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	 err = http.ListenAndServe(":8080", mux)
	 if err != nil {
	 	panic(err)
	 }
}
