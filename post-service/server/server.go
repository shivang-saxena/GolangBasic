package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "post-service/pb"
	store "post-service/store"
)

type server struct {
	Store store.PostStore
}

func NewPostServer(postStore store.PostStore) *server {
	return &server{Store: postStore}
}

//creation of new post to store into map
func (s *server) CreatePost(ctx context.Context, in *pb.Post) (*pb.Post, error) {
	err := in.Validate()
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	return s.Store.CreatePost(ctx, in)
}

func (s *server) GetPost(ctx context.Context, in *pb.GetPostRequest) (*pb.Post, error) {
	err := in.Validate()
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	return s.Store.GetPost(ctx, in.Id)
}

func (s *server) DeletePost(ctx context.Context, in *pb.DeletePostRequest) (*empty.Empty, error) {
	err := in.Validate()
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}
	_, errr := s.Store.DeletePost(ctx, in.Id)
	if errr != nil {
		return nil, status.Error(codes.Internal, errr.Error())
	}

	return &empty.Empty{}, nil
}

func (s *server) UpdatePost(ctx context.Context, in *pb.UpdatePostRequest) (*pb.Post, error) {
	err := in.Validate()
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	return s.Store.UpdatePost(ctx, in.Post)
}

func (s *server) ListPosts(ctx context.Context, em *empty.Empty) (*pb.ListPostResponse, error) {
	var posts []* pb.Post
	posts,_ =s .Store.ListPosts(ctx, em)
	return &pb.ListPostResponse{
		Posts: posts,
	},nil
}
