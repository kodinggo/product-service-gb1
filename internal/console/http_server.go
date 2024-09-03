package console

import (
	"github.com/go-playground/validator"
	"github.com/kodinggo/product-service-gb1/db"
	"github.com/kodinggo/product-service-gb1/internal/delivery/http"
	"github.com/kodinggo/product-service-gb1/internal/repository"
	"github.com/kodinggo/product-service-gb1/internal/usecase"
	"github.com/kodinggo/product-service-gb1/internal/utils"
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

	err = e.Start(":3232")
	continueOrFatal(err)
}

func continueOrFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
