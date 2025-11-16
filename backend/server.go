package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    pb "example/proto/echo"
)

type echoServer struct {
    pb.UnimplementedEchoServiceServer
}

func (s *echoServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoReply, error) {
    return &pb.EchoReply{Message: "Hello! This is Backend" + req.Message}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":6565")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()

    // Register gRPC service
    pb.RegisterEchoServiceServer(s, &echoServer{})

    // 👉 Enable reflection
    reflection.Register(s)

    log.Println("gRPC server listening on :6565")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
