package transportgrpc

import (
	"net"

	taskpb "github.com/Guglahai/project-protos/proto/task"
	userpb "github.com/Guglahai/project-protos/proto/user"
	"github.com/Guglahai/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc task.Service, userClient userpb.UserServiceClient) error {
	grpcServer := grpc.NewServer()
	handler := NewHandler(svc, userClient)
	taskpb.RegisterTaskServiceServer(grpcServer, handler)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}
	return grpcServer.Serve(listener)
}
