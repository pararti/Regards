// main
package main

import (
	"net"

	pb "github.com/pararti/Regards/api/golang"
	db "github.com/pararti/Regards/pkg/postgresdb"
	"google.golang.org/grpc"
)

func main() {
	err := loadConfig("../config")
	if err != nil {
		log3.Error.Fatalf("failed to load config: %v", err)
	}

	psqldb, err := db.NewDataBase(&psqlConf)
	if err != nil {
		log3.Error.Fatalf("failed to create db: %v", err)
	}
	err = psqldb.DB.Ping()
	if err != nil {
		log3.Error.Fatalf("failed to create db: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMediaAndSessionServer(s, &server{})

	lis, err := net.Listen("tcp", connConf.Address)
	if err != nil {
		log3.Error.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log3.Error.Fatalf("failed to serve: %v", err)
	}

}
