package service

import (
	"internshipApplicationTemplate/pkg/db"
)

type ChartaService struct {
	DB db.Charta
}

func NewChartaService(db db.Charta) *ChartaService {
	return &ChartaService{DB: db}
}
