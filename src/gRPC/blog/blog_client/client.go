package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Yepez1997/goProjects/src/gRPC/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client Starting on port 50052...")
	// create a connection to the server
	// by def grpc has ssl
	// once going in to production - remove
	//certFile := "../../ssl/ca.crt" // certificate authority trust cerificate
	//creds, err := credentials.NewClientTLSFromFile(certFile, "")
	// if err != nil {
	// 	log.Fatalf("Error loading client certificate: %v", err)
	// 	return
	// }
	// cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
		return
	}
	// when whole main is done -> close the connection
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	//doUnaryCreateBlog(c)
	fmt.Println("Starting to Create the Blog ...")
	blog := &blogpb.Blog{
		AuthorId: "Jim",
		Title:    "My Second Blog",
		Content:  "Content Of My Second Blog",
	}

	// unary api call to the server
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})

	if err != nil {
		fmt.Printf("Error from CreateBlog Client: %v", err)
	}
	fmt.Printf("Response from CreateBlog Service: %v", res.GetBlog())
	fmt.Println("Blog has been created ...")
}

// func doUnaryCreateBlog(c blogpb.BlogServiceClient) {
// 	fmt.Println("Starting to Create the Blog ...")
// 	blog := &blogpb.Blog{
// 		AuthorId: "Jim",
// 		Title:    "My Second Blog",
// 		Content:  "Content Of My Second Blog",
// 	}

// 	// unary api call to the server
// 	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})

// 	if err != nil {
// 		fmt.Printf("Error from CreateBlog Client: %v", err)
// 	}
// 	fmt.Printf("Response from CreateBlog Service: %v", res.GetBlog())
// 	fmt.Println("Blog has been created ...")
// }
