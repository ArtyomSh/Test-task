package repositories

import "TestTask/internal/models"

type RateRepo interface {
	SetRate(rate models.Rate) error
	GetRate(symbol string) (models.Rate, bool)
}
