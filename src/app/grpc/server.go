package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/iago-f-s-e/pix-code-go/src/app/grpc/pb"
	"github.com/iago-f-s-e/pix-code-go/src/app/useCases"
	"github.com/iago-f-s-e/pix-code-go/src/infra/repositories"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repositories.PixKeyRepositoryDb{Db: database}
	pixUseCase := useCases.PixUseCases{PixKeyRepository: pixRepository}

	pixService := NewPixGrpcService(pixUseCase)

	pb.RegisterPixServiceServer(grpcServer, pixService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

}
