package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	// import (
	// "context"
	// "log"
	// 	"net"

	// 	pb "github.com/skariyania/go-c0nceptual/example/todo-app"
	//
	// 	"google.com/golang.org/grpc"
	// )

	pb "github.com/skariyania/go-conceptual/example/todo/proto/todo"
)

const (
	PORT = ":50051"
)

type TodoServer struct {
	// pb.UnimplementedTodoServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.NewTodo) {
	log.Printf("received %v", in.GetName())
	todo := &pb.Todo{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Done:        false,
		Id:          uuid.New().String(),
	}

	return todo, nil
}

func main() {
	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to establish TCP connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTodoServer(s, &TodoServer{})

	log.Printf("server listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve request: %v", err)
	}
}
