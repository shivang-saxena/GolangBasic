package store

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"post-service/pb"
)


type PostStore interface {
	CreatePost(ctx context.Context, in *pb.Post) (*pb.Post, error)
	GetPost(ctx context.Context, id string)(*pb.Post, error)
	DeletePost(ctx context.Context, id string) (bool, error)
	UpdatePost(ctx context.Context,in *pb.Post) (*pb.Post, error)
	ListPosts(ctx context.Context,em *empty.Empty) ([]*pb.Post, error)
}


