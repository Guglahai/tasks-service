package transportgrpc

import (
	"context"
	"fmt"

	taskpb "github.com/Guglahai/project-protos/proto/task"
	userpb "github.com/Guglahai/project-protos/proto/user"
	"github.com/Guglahai/tasks-service/internal/task"
)

type Handler struct {
	svc        task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc task.Service, userClient userpb.UserServiceClient) *Handler {
	return &Handler{
		svc:        svc,
		userClient: userClient,
	}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	t, err := h.svc.CreateTask(&task.Task{UserID: uint(req.UserId), Task: req.Title})
	if err != nil {
		return nil, err
	}

	return &taskpb.CreateTaskResponse{Task: &taskpb.Task{Id: uint32(t.ID), UserId: uint32(t.UserID), Title: t.Task, IsDone: t.Is_done}}, nil
}

func (h *Handler) GetAllTasks(ctx context.Context, req *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var taskList []*taskpb.Task
	for _, t := range tasks {
		taskList = append(taskList, &taskpb.Task{Id: uint32(t.ID), UserId: uint32(t.UserID), Title: t.Task, IsDone: t.Is_done})
	}

	return &taskpb.ListTasksResponse{Task: taskList}, nil
}

func (h *Handler) GetTaskByID(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	t, err := h.svc.GetTaskByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &taskpb.GetTaskResponse{Task: &taskpb.Task{Id: uint32(t.ID), UserId: uint32(t.UserID), Title: t.Task, IsDone: t.Is_done}}, nil
}

func (h *Handler) GetTasksByUserID(ctx context.Context, req *taskpb.ListTasksByUserRequest) (*taskpb.ListTasksByUserResponse, error) {
	tasks, err := h.svc.GetTasksByUserID(uint(req.UserId))
	if err != nil {
		return nil, err
	}

	var taskList []*taskpb.Task
	for _, t := range tasks {
		taskList = append(taskList, &taskpb.Task{Id: uint32(t.ID), UserId: uint32(t.UserID), Title: t.Task, IsDone: t.Is_done})
	}

	return &taskpb.ListTasksByUserResponse{Task: taskList}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	t, err := h.svc.UpdateTask(&task.Task{
		ID:      uint(req.Task.Id),
		UserID:  uint(req.Task.UserId),
		Task:    req.Task.Title,
		Is_done: req.Task.IsDone,
	})
	if err != nil {
		return nil, err
	}

	return &taskpb.UpdateTaskResponse{Task: &taskpb.Task{Id: uint32(t.ID), UserId: uint32(t.UserID), Title: t.Task, IsDone: t.Is_done}}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	if err := h.svc.DeleteTask(uint(req.Id)); err != nil {
		return nil, err
	}

	return &taskpb.DeleteTaskResponse{Done: true}, nil
}
