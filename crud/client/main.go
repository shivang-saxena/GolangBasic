package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"os"
	"strings"
	"time"

	pb "crud/pb"

	"google.golang.org/grpc"
)


func menu(){
	fmt.Println("Press 0 for menu.")
	fmt.Println("Press 1 to create a post.")
	fmt.Println("Press 2 to get a post.")
	fmt.Println("Press 3 to delete a post.")
	fmt.Println("Press 4 to update a post.")
	fmt.Println("Press 5 to show all the post")
	fmt.Println("Press -1 to quit")
}

func main() {
	//connection dialer
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPostsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 300)
	defer cancel()

	menu()
	var flag bool = true
	for flag {
		var inp string
		fmt.Scanf("%s", &inp)
		buf := bufio.NewReader(os.Stdin)
		switch inp {
		case "0":
		menu()

		case "1":
			fmt.Print("Enter post ID : ")
			id, _ := buf.ReadString('\n')
			id = strings.Trim(id, "\n\r")

			fmt.Print("Enter post title : ")
			title, _ := buf.ReadString('\n')
			title = strings.Trim(title, "\n\r")

			fmt.Print("Enter post Description : ")
			description, _ := buf.ReadString('\n')
			description = strings.Trim(description, "\n\r")

			res, err := c.CreatePost(ctx, &pb.Post{
				Id:          id,
				Title:       title,
				Description: description,
			})
			if err != nil {
				fmt.Println("could not create:", err)
			}

			fmt.Println("Post Created:", res)

		case "2":
			fmt.Print("Enter id to get Post: ")
			id, _ := buf.ReadString('\n')
			id = strings.Trim(id, "\n\r")
			res, err := c.GetPost(ctx, &pb.PostRequest{
				Id: id,
			})
			if err != nil || res == nil {
				fmt.Println("could not Get Post:")
			}else {
				fmt.Println("Found Post: ", res)
			}

		case "3":
			fmt.Print("Enter id to Delete Post: ")
			id, _ := buf.ReadString('\n')
			id = strings.Trim(id, "\n\r")
			res, err := c.DeletePost(ctx, &pb.PostRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("could not Delete Post: %v", err)
			}

			fmt.Println("Post Deleted", res)

		case "4":
			fmt.Print("Enter post ID to Update : ")
			id, _ := buf.ReadString('\n')
			id = strings.Trim(id, "\n\r")

			fmt.Print("Enter post title : ")
			title, _ := buf.ReadString('\n')
			title = strings.Trim(title, "\n\r")

			fmt.Print("Enter post Description : ")
			description, _ := buf.ReadString('\n')
			description = strings.Trim(description, "\n\r")

			res, err := c.UpdatePost(ctx, &pb.Post{
				Id:          string(id),
				Title:       string(title),
				Description: string(description),
			})
			if err != nil {
				log.Fatalf("could not Get Post: %v", err)
			}

			fmt.Println("Post Updated", res)

		case "5":
			res, err := c.ListPost(context.Background(), &empty.Empty{})
			if err != nil {
				log.Fatalf("could not Get Post: %v", err)
			}

			log.Println("All Posts", res)
		case "-1":
			flag=false

		default:
			fmt.Println("\nWrong option!")
			menu()
		}
	}
}
