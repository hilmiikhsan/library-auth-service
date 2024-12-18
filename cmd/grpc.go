package cmd

import (
	"net"

	"github.com/hilmiikhsan/library-auth-service/cmd/proto/tokenvalidation"
	"github.com/hilmiikhsan/library-auth-service/helpers"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	dependency := dependencyInject()

	server := grpc.NewServer()

	tokenvalidation.RegisterTokenValidationServer(server, dependency.TokenValidationAPI)

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "6000"))
	if err != nil {
		helpers.Logger.Fatal("failed to listen grpc port: ", err)
	}

	helpers.Logger.Info("start listening grpc on port:" + helpers.GetEnv("GRPC_PORT", "6000"))
	if err := server.Serve(lis); err != nil {
		helpers.Logger.Fatal("failed to serve grpc port: ", err)
	}
}
