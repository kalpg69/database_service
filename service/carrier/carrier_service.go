package carrier

import (
	"context"
	"database/sql"
	"log"

	"github.com/kalpg69/database_service/api/v1/carrierpb"
	"github.com/kalpg69/database_service/api/v1/databasepb"
)

const (
	queryCreateCarrier = "CreateCarrier" //Name of Stored Procedure
)

func NewCarrierServer(db *sql.DB) databasepb.CarrierServiceServer {
	log.Println("starting carrier server")
	return &carrierServer{
		db: db,
	}
}

type carrierServer struct {
	databasepb.UnimplementedCarrierServiceServer
	db *sql.DB
}

func (cs *carrierServer) CreateCarrier(ctx context.Context, req *carrierpb.CreateCarrierRequest) (*carrierpb.CreateCarrierResponse, error) {
	var id int64
	newCarrier := req.GetCarrier()

	_, err := cs.db.ExecContext(ctx, queryCreateCarrier,
		newCarrier.GetCarrierCode(),
		newCarrier.GetCarrierName(),
		newCarrier.GetCarrierAddress(),
		newCarrier.GetCarrierEmail(),
		newCarrier.GetCarrierPhone(),
		sql.Named("LastInsertedId", sql.Out{Dest: &id}))
	if err != nil {
		log.Fatalf("error while creating carrier: %v", err)
		return nil, err
	}

	return &carrierpb.CreateCarrierResponse{
		Carrier: &carrierpb.Carrier{
			Id: int32(id),
		},
	}, nil

}
