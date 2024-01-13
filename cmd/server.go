package cmd

import (
	"TestTask/internal/configs"
	"TestTask/internal/handlers"
	"TestTask/internal/repositories"
	"TestTask/pkg/Ticker"
	"TestTask/pkg/loggers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Long:  `Run server`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := loggers.New(true)
		logger.Info("logger created")

		cfg := configs.GetConfig(logger)
		logger.Info("config received")

		var repo repositories.RateRepo
		switch cfg.Repository.Repo {
		case "redis":
			r := repositories.NewRedisRepo(*cfg)
			repo = &r
			logger.Info("redis rate repository created")
		case "memory":
			r := repositories.NewMemoryRepo()
			repo = &r
			logger.Info("memory rate repository created")
		default:
			log.Fatalln("Unsupported rate repository")
		}

		h := handlers.New(repo)

		logger.Info("router created")
		router := mux.NewRouter()

		router.HandleFunc("/api/v1/rates", h.GETPair).Methods(http.MethodGet)
		router.HandleFunc("/api/v1/rates", h.POSTPair).Methods(http.MethodPost)
		logger.Info("API is running!")

		updater := Ticker.New()
		updater.RunUpdate(*cfg, repo)
		logger.Info("Rate update started")

		port := cfg.Listen.Port
		fmt.Printf("Listening at port %s\n", cfg.Listen.Port)
		http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
