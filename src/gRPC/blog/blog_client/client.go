package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Yepez1997/goProjects/src/gRPC/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client Starting on port 50051...")
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
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
		return
	}
	// when whole main is done -> close the connection
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	//doUnaryCreateBlog(c)
	fmt.Println("Starting to Create the Blog ...")
	// create Blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Yepez3",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v\n", createBlogRes)

	blogID := createBlogRes.GetBlog().GetId()

	fmt.Println("Reading the blog ...")
	readBlogRes, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: blogID})
	if err != nil {
		log.Fatalf("Error while retrieving the blog: %v", err)
	}
	fmt.Printf("Blog was read: %v\n", readBlogRes)

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
