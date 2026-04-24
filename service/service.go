package service

import (
	"smallcase/models"
	"smallcase/repo"
)

type Service interface {
	Buy(userId string, stock *models.Stock) error
	FetchingHoldings(userId string) map[string]*models.Stock
	FetchingReturns(userId string, currentPrice map[string]int) int
	Sell(userId string, stock *models.Stock) error
}

type service struct {
	repo repo.Repo
}

func NewService(r repo.Repo) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Buy(userId string, stock *models.Stock) error {
	return s.repo.Buy(userId, stock)
}
func (s *service) FetchingHoldings(userId string) map[string]*models.Stock {
	return s.repo.FetchingHoldings(userId)
}
func (s *service) FetchingReturns(userId string, currentPrice map[string]int) int {
	return s.repo.FetchingReturns(userId, currentPrice)
}

func (s *service) Sell(userId string, stock *models.Stock) error {
	return s.repo.Sell(userId, stock)
}
