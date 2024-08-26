package console

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kodinggo",
	Short: "Kodinggo App",
	Long:  `Kodinggo is a CLI application for managing your story.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setupLogger()
	loadEnv()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found")
	}
}

func setupLogger() {
	log := logrus.New()

	// Set JSON formatter
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}
