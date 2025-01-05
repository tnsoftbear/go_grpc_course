package main

import (
	"context"
	"fmt"
	"github.com/cshep4/grpc-course/02-exercise-solution/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), // заблокируется до создания соединения
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewTodoServiceClient(conn)
	res, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "Work"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response1: %s\n", res)

	res, err = client.AddTask(ctx, &proto.AddTaskRequest{})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic(err)
		}
		fmt.Printf("Response code: %s, message: %s, details: %s, proto: %s\n", st.Code(), st.Message(), st.Details(), st.Proto())
	}
	fmt.Printf("Response2: %s\n", res)

	res, err = client.AddTask(ctx, &proto.AddTaskRequest{Task: "Run"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response3: %s\n", res)

	ctRes, err := client.CompleteTask(ctx, &proto.CompleteTaskRequest{Id: "1"})
	if err != nil {
		panic(err)
	}
	fmt.Println(ctRes)

	tasks, err := client.ListTasks(ctx, &proto.ListTasksRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(tasks.GetTasks())

}
