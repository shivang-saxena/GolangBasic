package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"log"
	"post-service/pb"
	"time"
)

type PostPgStore struct {
	*sql.DB
}

func NewPostStore(db *sql.DB) PostStore{
	return &PostPgStore{db}
}

func (ob *PostPgStore) CreatePost(ctx context.Context, in *pb.Post) (*pb.Post, error){
	id, _ := uuid.NewUUID()
	ut := new(time.Time)
	err := ob.QueryRowContext(ctx, `INSERT INTO posts(postid,title,description,createdon,updatedon) VALUES($1, $2, $3, $4,$5) RETURNING postid`, id, in.Title, in.Description, time.Now(),ut).Scan(&in.Id)
	if err != nil {
		return in, err
	}
	in.CreatedOn,_ = ptypes.TimestampProto(time.Now())
	in.UpdatedOn,_ = ptypes.TimestampProto(*ut)
	return in,nil
}

func (ob *PostPgStore) GetPost(ctx context.Context, id string)(*pb.Post, error){
	rows, err := ob.QueryContext(ctx,"SELECT postid,title,description,createdon,updatedon FROM posts WHERE postid=$1;", id)
	if err != nil {
		return nil, err
	}

	var createdTime time.Time
	var updatedTime time.Time
	post := &pb.Post{}
	for rows.Next() {
		if err := rows.Scan(&post.Id,&post.Title,&post.Description,&createdTime,&updatedTime); err != nil {
			log.Fatal(err)
		}
		post.CreatedOn,_ = ptypes.TimestampProto(createdTime)
		post.UpdatedOn, _ = ptypes.TimestampProto(updatedTime)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return post, nil
}

func (ob *PostPgStore) DeletePost(ctx context.Context, id string) (bool, error){
	sqlStatement := "DELETE FROM posts WHERE postid = $1;"
	_, err := ob.ExecContext(ctx, sqlStatement, id)
	if err != nil {
		panic(err)
	}
	//count , err := res.RowsAffected()
	if err != nil {
		panic(err)
	} else {
		return true, nil
	}

}

func (ob *PostPgStore) UpdatePost(ctx context.Context,in *pb.Post) (*pb.Post, error){
	sqlStatement := "UPDATE posts SET title = $2, description = $3, updatedon= $4  WHERE postid = $1;"
	res, err := ob.ExecContext(ctx,sqlStatement, in.Id, in.Title, in.Description, time.Now())
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		fmt.Println(count)
		return in, nil
	} else {
		return in, nil
	}
}

func (ob *PostPgStore) ListPosts(ctx context.Context,em *empty.Empty) ([]*pb.Post, error){
	rows, err := ob.QueryContext(ctx,"SELECT postid,title,description,createdon,updatedon FROM posts")
	if err != nil {
		fmt.Println(err)
	}
	var createdTime time.Time
	var updatedTime time.Time
	posts := make([]*pb.Post, 0)
	for rows.Next() {
		post := &pb.Post{}
		if err := rows.Scan(&post.Id, &post.Title, &post.Description, &createdTime,&updatedTime); err != nil {
			log.Fatal(err)
		}
		post.CreatedOn,_ = ptypes.TimestampProto(createdTime)
		post.UpdatedOn,_ = ptypes.TimestampProto(updatedTime)
		posts = append(posts,post)
	}
	return posts, nil
}
