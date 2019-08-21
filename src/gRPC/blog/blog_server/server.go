package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/Yepez1997/goProjects/src/gRPC/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

// globally accesible collection
var collection *mongo.Collection

// all services - added on needed basis
type server struct{}

// blogItem api definition - goes along well with the protocal buffer message
// TODO: look over primitive type
type blogItem struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	AuthorID string             `bson:"_author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func main() {

	// get the  go code we get the file name and line number - if crashes
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// connectiing to mongodb; client represents a client object ot mongodb
	fmt.Println("Connecting to MongoDB ...")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error while connecting to the client: %v", err)
	}

	// open up a connection
	// from the client collect the database and chose the blog collection
	// make it global
	fmt.Println("Blog Service Started ...")
	collection = client.Database("mydb").Collection("blog")

	fmt.Print("Blog Server Started ...\n")
	// create connection; and port binding
	listener, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// adding authentication
	// server certificate
	// certFile := "../../ssl/server.crt"
	// // .pem file is a file grpc can read
	// keyFile := "../../ssl/server.pem"
	// cred, SSLerr := credentials.NewServerTLSFromFile(certFile, keyFile)
	// if SSLerr != nil {
	// 	fmt.Printf("Error loading credentials: %v", err)
	// 	return
	// }
	// pass in the credentials to the grpc call
	// grpc.Creds(cred)

	s := grpc.NewServer()
	// the path to the protocol buffer
	blogpb.RegisterBlogServiceServer(s, &server{})

	// adding reflection service to the server
	//reflection.Register(s)
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// interrupt on signal; wait for control c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until the signal is received
	<-ch
	fmt.Println("Stopping the server ...")
	s.Stop()
	fmt.Println("Stopping the listener ...")
	listener.Close()
	fmt.Println("Closing MongoDB ...")
	client.Disconnect(ctx)
	fmt.Println("## End of the program ##")

}
