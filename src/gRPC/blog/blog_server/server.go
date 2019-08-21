package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/Yepez1997/goProjects/src/gRPC/blog/blogpb"
	"google.golang.org/grpc"
)

// all services - added on needed basis
type server struct{}

func main() {

	// get the  go code we get the file name and line number - if crashes
	log.SetFlags(log.LstdFlags | log.Lshortfile)



	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
	"mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	))
	if err != nil { log.Fatal(err) }

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
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Stopping the listener")
	listener.Close()

}
