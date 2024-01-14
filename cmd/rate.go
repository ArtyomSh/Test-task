package cmd

import (
	"TestTask/internal/configs"
	"TestTask/internal/models"
	"TestTask/pkg/loggers"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

var rateCmd = &cobra.Command{
	Use:   "rate",
	Short: "Get rate command",
	Long:  `Returns the rate on the pairs passed in the "pairs" parameter`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := configs.GetConfig()

		logger := loggers.New(cfg.Logging.Verbosity)
		logger.Info("logger created")

		pairs, _ := cmd.Flags().GetString("pairs")
		if pairs == "" {
			logger.Warn("Empty flag")
			return
		}
		getResponse := fmt.Sprintf("http://localhost:%s/api/v1/rates?pairs=%s", cfg.Listen.Port, pairs)
		resp, err := http.Get(getResponse)
		if err != nil {
			logger.Error(err.Error())
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error(err.Error())
			return
		}

		var Response models.Response
		err = json.Unmarshal(body, &Response)
		if err != nil {
			logger.Error(err.Error())
			return
		}

		if e := Response.Error; e != "" {
			logger.Error(e)
			return
		}

		fmt.Println(Response.Message)
		logger.Info("rate called")
	},
}

func init() {
	rootCmd.AddCommand(rateCmd)
	rateCmd.Flags().String("pairs", "", "String that contain pairs")
}
