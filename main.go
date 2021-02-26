package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kalpg69/database_service/api/v1/carrierpb"
	"github.com/kalpg69/database_service/api/v1/customerpb"
	"github.com/kalpg69/database_service/server/carriersrv"
	"github.com/kalpg69/database_service/server/customersrv"
	"github.com/kalpg69/database_service/sqlclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Starting Database Service!")
	db, err := sqlclient.NewSQLClient("sa", "itas@123", "localhost", "blackduck", "30")
	if err != nil {
		return
	}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("database service failed to listen: %v", err)
		return
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(grpcServer, customersrv.NewCustomerServer(db))
	carrierpb.RegisterCarrierServiceServer(grpcServer, carriersrv.NewCarrierServer(db))

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
