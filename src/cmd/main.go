package main

import (
	"os"

	"github.com/iago-f-s-e/pix-code-go/src/app/grpc"
	"github.com/iago-f-s-e/pix-code-go/src/infra/db"
	"gorm.io/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)
}
