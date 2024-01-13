package Ticker

import (
	"TestTask/internal/configs"
	"TestTask/internal/models"
	"TestTask/internal/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type updater struct {
	Ticker *time.Ticker
	Done   chan bool
}

func New() updater {
	ticker := time.NewTicker(5000 * time.Millisecond)
	done := make(chan bool)
	return updater{Ticker: ticker, Done: done}
}

func UpdateRates(cfg configs.Config) []models.Rate {
	response, err := http.Get(cfg.Binance.URL)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	var data []models.Rate
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
		return nil
	}
	return data
}

func (u *updater) RunUpdate(cfg configs.Config, repo repositories.RateRepo) {
	go func() {
		for {
			select {
			case <-u.Done:
				u.Ticker.Stop()
				return
			case <-u.Ticker.C:
				//fmt.Println("Update rates", t)
				for _, rate := range UpdateRates(cfg) {
					err := repo.SetRate(rate)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}()
}
