package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Yepez1997/goProjects/src/gRPC/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// globally accesible collection
var collection *mongo.Collection

// all services - added on needed basis
type server struct{}

// blogItem api definition - goes along well with the protocal buffer message
// TODO: look over primitive type
type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

// CreateBlog - BlogService unary rpc call
func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	// get the data
	fmt.Println("Creating Blog Request")
	blog := req.GetBlog()
	// create the data
	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	res, err := collection.InsertOne(context.Background(), data)
	// send the data to the mongo db and specify errors
	if err != nil {
		log.Fatalf("Error while inserting into the collection: %v", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	// cast the interface
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Content:  blog.GetContent(),
			Title:    blog.GetTitle(),
		},
	}, nil

}

// ReadBlog - gets the blog in the blog database
func (*server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	fmt.Println("Read Blog Request")
	blogID := req.GetBlogId()
	// check if object id is present
	oid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot Parse ID"),
		)
	}

	// retrieve data from mongo db
	data := &blogItem{}
	// pass in the filter (look at documentation for this)
	filter := bson.M{"_id": oid}
	res := collection.FindOne(context.Background(), filter)
	// reminds me of c
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Cannot find blog with the specified ID"),
		)
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, nil

}

// UpdateBlog
func (*server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	fmt.Println("Updating Blog Request")
	// parse the object
	blog := req.GetBlog()
	oid, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("The blog was not found ..."),
		)
	}

	data := &blogItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(context.Background(), filter)
	// reminds me of c
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Cannot find blog with the specified ID"),
		)
	}

	// update the responses i.e update the internal struct
	data.AuthorID = blog.GetAuthorId()
	data.Content = blog.GetContent()
	data.Title = blog.GetTitle()

	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error while Updating: %v", updateErr),
		)
	}

	return &blogpb.UpdateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, nil

}

// DeleteBlog
func (*server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	fmt.Println("Deleting Blog Request")
	oid, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("The blog was not found ..."),
		)
	}
	filter := bson.M{"_id": oid}
	delRes, delErr := collection.DeleteOne(context.Background(), filter)
	if delErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete in MongoDB: %v", delErr),
		)
	}

	if delRes.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find document in MongoDB: %v", delErr),
		)
	}

	return &blogpb.DeleteBlogResponse{
		BlogId: req.GetBlogId(),
	}, nil

}

func ListBlog(req *blogpb.ListBlogRequest, stream *blogpb.BlogService_ListBlogServer) error {
	fmt.Println("List Blog Request")
	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cur.Close(context.Background())
	// for each time we can view the next element - take some data and decode it into the obj
	for cur.Next(context.Background()) {
		data := &blogItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}
		// send the data 
		stream.Send(&blogpb.ListBlogResponse{
			Blog: &blogpb.Blog{
				Id: data.ID.Hex(),
				AuthorId: data.AuthorID,
				Title: data.Title,
				Content: data.Content,
			}
		})
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Internal Error: %v", err),
		)
	}
	return nil
}

func main() {

	// get the  go code we get the file name and line number - if crashes
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// connectiing to mongodb; client represents a client object ot mongodb
	fmt.Println("Connecting to MongoDB")
	// connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// open up a connection
	// from the client collect the database and chose the blog collection
	fmt.Println("Blog Service Started ...")
	collection = client.Database("mydb").Collection("blog")

	fmt.Print("Blog Server Started ...\n")
	// create connection; and port binding
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
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
	client.Disconnect(context.TODO())
	fmt.Println("## End of the program ##")

}
