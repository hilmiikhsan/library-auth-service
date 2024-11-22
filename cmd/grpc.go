package cmd

import (
	"log"
	"net"

	"github.com/hilmiikhsan/library-auth-service/helpers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	server := grpc.NewServer()

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "7000"))
	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}

	logrus.Info("start listening grpc on port:" + helpers.GetEnv("GRPC_PORT", "7000"))
	if err := server.Serve(lis); err != nil {
		log.Fatal("failed to serve grpc port: ", err)
	}
}
