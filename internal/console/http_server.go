package console

import (
	"github.com/kodinggo/product-service-gb1/db"
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

	v1 := e.Group("/api/v1")
	v1.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	err = e.Start(":3232")
	continueOrFatal(err)
}

func continueOrFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
