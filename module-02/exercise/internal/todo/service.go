package todo

import (
	"context"
	"github.com/cshep4/grpc-course/02-exercise-solution/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type TodoService struct {
	proto.UnimplementedTodoServiceServer
	tasks map[string]string
	maxID int
}

func New() *TodoService {
	return &TodoService{
		tasks: make(map[string]string),
	}
}

func (s *TodoService) AddTask(ctx context.Context, request *proto.AddTaskRequest) (*proto.AddTaskResponse, error) {
	if request.Task == "" {
		return nil, status.Error(codes.InvalidArgument, "task cannot be empty")
	}
	s.maxID++
	var id = strconv.Itoa(s.maxID)
	s.tasks[id] = request.GetTask()
	return &proto.AddTaskResponse{
		Id: id,
	}, nil
}

func (s *TodoService) CompleteTask(ctx context.Context, request *proto.CompleteTaskRequest) (*proto.CompleteTaskResponse, error) {
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "ID cannot be empty")
	}
	if _, ok := s.tasks[request.Id]; !ok {
		return nil, status.Error(codes.NotFound, "task not found by ID")
	}
	delete(s.tasks, request.GetId())
	return &proto.CompleteTaskResponse{}, nil
}

func (s *TodoService) ListTasks(ctx context.Context, request *proto.ListTasksRequest) (*proto.ListTasksResponse, error) {
	tasks := make([]*proto.Task, 0, len(s.tasks))
	for id, task := range s.tasks {
		tasks = append(tasks, &proto.Task{
			Id:   id,
			Task: task,
		})
	}
	return &proto.ListTasksResponse{Tasks: tasks}, nil
}
