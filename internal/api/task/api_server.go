package task

import (
	"artemzinchenko.com/auth/pb"
	"context"
	"google.golang.org/grpc"
)

type ApiServer struct {
	pb.UnimplementedTaskApiServer
}

func NewApiServer() (*ApiServer, error) {
	return &ApiServer{}, nil
}

func (a ApiServer) RegisterService(g grpc.ServiceRegistrar) {
	pb.RegisterTaskApiServer(g, a)
}

func (a ApiServer) Add(ctx context.Context, req *pb.AddTaskRequest) (*pb.TaskResponse, error) {
	return &pb.TaskResponse{}, nil
}

func (a ApiServer) Archive(ctx context.Context, req *pb.ArchiveRequest) (*pb.TaskResponse, error) {
	return &pb.TaskResponse{}, nil
}

func (a ApiServer) MarkDone(ctx context.Context, req *pb.MarkDoneRequest) (*pb.TaskResponse, error) {
	return &pb.TaskResponse{}, nil
}
