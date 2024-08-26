package console

import (
	"github.com/kodinggo/product-service-gb1/db"
	"github.com/kodinggo/product-service-gb1/internal/delivery/http"
	"github.com/kodinggo/product-service-gb1/internal/repository"
	"github.com/kodinggo/product-service-gb1/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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

	e := echo.New()

	categoryRepo := repository.NewCategoryRepository(mysql)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	handler := http.NewHTTPHandler()
	handler.RegisterCategoryUsecase(categoryUsecase)

	handler.Routes(e)

	err = e.Start(":3232")
	continueOrFatal(err)
}

func continueOrFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
