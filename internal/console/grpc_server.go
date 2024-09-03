package console

import (
	"log"
	"net"

	"github.com/kodinggo/product-service-gb1/db"
	"github.com/kodinggo/product-service-gb1/internal/delivery/grpcservice"
	"github.com/kodinggo/product-service-gb1/internal/repository"
	"github.com/kodinggo/product-service-gb1/internal/usecase"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pb "github.com/kodinggo/product-service-gb1/pb/product"
)

func init() {
	rootCmd.AddCommand(grpcCmd)
}

var grpcCmd = &cobra.Command{
	Use:   "grpcsrv",
	Short: "Start the GRPC server",
	Run:   grpcServer,
}

func grpcServer(cmd *cobra.Command, args []string) {
	mysql := db.NewMysql()

	db, err := mysql.DB()
	continueOrFatal(err)

	defer db.Close()

	s := grpc.NewServer()

	productRepo := repository.NewProductRepository(mysql)
	productUsecase := usecase.NewProductUsecase(productRepo)

	svc := grpcservice.NewProductService(productUsecase)

	pb.RegisterProductServiceServer(s, svc)

	lis, err := net.Listen("tcp", ":8080")
	continueOrFatal(err)

	log.Println("GRPC server is running on port :8080")

	err = s.Serve(lis)
	continueOrFatal(err)
}
