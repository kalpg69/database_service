package customer

import (
	"context"
	"database/sql"
	"log"

	"github.com/kalpg69/database_service/api/v1/customerpb"
	"github.com/kalpg69/database_service/api/v1/databasepb"
)

const (
	queryCreateCustomer = "CreateCustomer" //Name of Stored Procedure
)

func NewCustomerServer(db *sql.DB) databasepb.CustomerServiceServer {
	log.Println("starting customer server")
	return &customerServer{
		db: db,
	}
}

type customerServer struct {
	databasepb.UnimplementedCustomerServiceServer
	db *sql.DB
}

func (cs *customerServer) CreateCustomer(ctx context.Context, req *customerpb.CreateCustomerRequest) (*customerpb.CreateCustomerResponse, error) {
	var id int64
	newCustomer := req.GetCustomer()

	_, err := cs.db.ExecContext(ctx, queryCreateCustomer,
		newCustomer.GetCustomerCode(),
		newCustomer.GetCustomerName(),
		newCustomer.GetCustomerAddress(),
		newCustomer.GetCustomerEmail(),
		newCustomer.GetCustomerPhone(),
		sql.Named("LastInsertedId", sql.Out{Dest: &id}))
	if err != nil {
		log.Fatalf("error while creating customer: %v", err)
		return nil, err
	}

	return &customerpb.CreateCustomerResponse{
		Customer: &customerpb.Customer{
			Id: int32(id),
		},
	}, nil

}
