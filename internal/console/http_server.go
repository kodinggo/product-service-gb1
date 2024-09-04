package console

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-playground/validator"
	"github.com/kodinggo/product-service-gb1/db"
	"github.com/kodinggo/product-service-gb1/internal/delivery/grpcservice"
	"github.com/kodinggo/product-service-gb1/internal/delivery/http"
	"github.com/kodinggo/product-service-gb1/internal/repository"
	"github.com/kodinggo/product-service-gb1/internal/usecase"
	"github.com/kodinggo/product-service-gb1/internal/utils"
	pb "github.com/kodinggo/product-service-gb1/pb/product"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authPb "github.com/kodinggo/user-service-gb1/pb/auth"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "httpsrv",
	Short: "Start the HTTP server",
	Run:   httpServer,
}

func httpServer(cmd *cobra.Command, args []string) {
	mysql := db.NewMysql()

	db, err := mysql.DB()
	continueOrFatal(err)

	defer db.Close()

	conn, err := grpc.NewClient("localhost:4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	continueOrFatal(err)

	defer conn.Close()

	auth := authPb.NewJWTValidatorClient(conn)
	continueOrFatal(err)

	e := echo.New()
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}

	categoryRepo := repository.NewCategoryRepository(mysql)
	productRepo := repository.NewProductRepository(mysql)

	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	productUsecase := usecase.NewProductUsecase(productRepo)

	handler := http.NewHTTPHandler()
	handler.RegisterAuthClient(auth)
	handler.RegisterCategoryUsecase(categoryUsecase)
	handler.RegisterProductUsecase(productUsecase)

	authMiddleware := utils.NewJWTMiddleware(auth)

	handler.Routes(e, authMiddleware.ValidateJWT)

	s := grpc.NewServer()

	svc := grpcservice.NewProductService(productUsecase)

	pb.RegisterProductServiceServer(s, svc)

	lis, err := net.Listen("tcp", ":8080")
	continueOrFatal(err)

	// initiate graceful shutdown
	var wg sync.WaitGroup
	errCh := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := e.Start(":3232"); err != nil {
			errCh <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Println("GRPC server is running on port :8080")

		if err := s.Serve(lis); err != nil {
			errCh <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("Shutting down servers...")

		fmt.Println("anjsakjdnk")

		// Gracefully shutdown the Echo server
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			continueOrFatal(fmt.Errorf("Echo server shutdown failed: %v", err))
		}

		// Gracefully stop the gRPC server
		s.GracefulStop()

	case err := <-errCh:
		continueOrFatal(err)
	}

	wg.Wait()
	log.Println("Servers stopped gracefully")
}

func continueOrFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
