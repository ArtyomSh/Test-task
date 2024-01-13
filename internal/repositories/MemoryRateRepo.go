package repositories

import (
	"TestTask/internal/models"
	"sync"
)

type MemoryRateRepo struct {
	Rates map[string]string
	mx    sync.Mutex
}

func NewMemoryRepo() MemoryRateRepo {
	return MemoryRateRepo{Rates: make(map[string]string)}
}

func (m *MemoryRateRepo) SetRate(rate models.Rate) error {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.Rates[rate.Symbol] = rate.Price

	return nil
}

func (m *MemoryRateRepo) GetRate(symbol string) (models.Rate, bool) {
	m.mx.Lock()
	defer m.mx.Unlock()

	if val, ok := m.Rates[symbol]; ok {
		return models.Rate{
			Symbol: symbol,
			Price:  val,
		}, true
	}
	return models.Rate{}, false
}
