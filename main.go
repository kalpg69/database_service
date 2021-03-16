package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kalpg69/database_service/api/v1/databasepb"
	"github.com/kalpg69/database_service/service/carrier"
	"github.com/kalpg69/database_service/service/customer"
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

	databasepb.RegisterCarrierServiceServer(grpcServer, carrier.NewCarrierServer(db))

	databasepb.RegisterCustomerServiceServer(grpcServer, customer.NewCustomerServer(db))

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
