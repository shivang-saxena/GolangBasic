package main

import (
	"context"
	pb "crud/pb"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {

}

//*Post type map to store product information
var postMap = map[string]*pb.Post{}

//creation of new post to store into map
func (s *server) CreatePost(ctx context.Context, in *pb.Post) (*pb.Post, error) {
	fmt.Println("CreatePost: ",in.Id)
	postMap[in.GetId()] = in
	createdOn, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return nil, err
	}

	in.CreatedOn = createdOn
	return in, nil
}


//Get post from the Post map
func (s *server) GetPost(ctx context.Context, in *pb.PostRequest) (*pb.Post, error) {
	fmt.Println("GetPost: ",in.Id)
	return  postMap[in.GetId()], nil
}

// Delete Post from the Post type map
func (s *server) DeletePost(ctx context.Context, in *pb.PostRequest)(*empty.Empty, error) {
	fmt.Println("DeletePost: ",in.Id)
	delete(postMap, in.Id)
	return &empty.Empty{},nil
}

//Update the previous stored post in Map
func (s *server) UpdatePost(ctx context.Context, in *pb.Post) (*pb.Post, error) {
	fmt.Println("UpdatePost: ",in.Id)
	postMap[in.GetId()] = in
	updatedOn, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return nil, err
	}

	in.UpdatedOn = updatedOn
	return in, nil

}

//List all the post that store into tha map
func (s *server) ListPost(ctx context.Context, em *empty.Empty) (*pb.ListPostResponse, error) {
	fmt.Println("ListPost: ")
	post := make([]*pb.Post,0)
	for _ , val := range postMap{
		post = append(post , val)
	}
	fmt.Println(post)
	return &pb.ListPostResponse{
		Post:            post    ,
	},nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		
	}
}
