package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/skariyania/go-conceptual/example/todo/proto"
)

const (
	SERVER_ADDRESS = "localhost:50051"
)

type TodoTask struct {
	Name        string
	Description string
	Done        bool
}

func main() {
	conn, err := grpc.NewClient(
		SERVER_ADDRESS,
		grpc.DialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	todoList := []TodoTask{
		{Name: "GRPC server Development", Description: "Development activity", Done: false},
		{Name: "GRPC Client Development", Description: "Development activity", Done: false},
		{Name: "GRPC PR Review", Description: "Review GRPC Pull Request", Done: false},
		{Name: "GRPC Server Local Setup", Description: "Setting up GRPC client and server in local", Done: false},
	}

	for _, todo := range todoList {
		res, err := c.CreateTodo(ctx, &pb.NewTodo{Name: todo.Name, Description: todo.Description, Done: todo.Done})

		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}

		log.Printf(`
           ID : %s
           Name : %s
           Description : %s
           Done : %v,
       `, res.GetId(), res.GetName(), res.GetDescription(), res.GetDone())
	}
}
