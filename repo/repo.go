package repo

import (
	"fmt"
	"smallcase/models"
)

type Repo interface {
	Buy(userId string, stock *models.Stock) error
	FetchingHoldings(userId string) map[string]*models.Stock
	FetchingReturns(userId string, currentPrice map[string]int) int
	Sell(userId string, stock *models.Stock) error
}

type repo struct {
	holdings map[string]map[string]*models.Stock
}

func NewRepo() Repo {
	return &repo{
		holdings: make(map[string]map[string]*models.Stock),
	}
}

func (r *repo) Buy(userId string, stock *models.Stock) error {
	//check if i have already that stock
	oldStock := r.holdings[userId][stock.TickerSymbol]

	if oldStock == nil {
		r.holdings[userId] = make(map[string]*models.Stock)
		r.holdings[userId][stock.TickerSymbol] = stock
	} else {
		oldAmount := oldStock.AveragePrice * oldStock.Quantity
		newAmount := stock.AveragePrice * stock.Quantity
		newAverage := (oldAmount + newAmount) / (oldStock.Quantity + stock.Quantity)
		oldStock.AveragePrice = newAverage
		oldStock.Quantity += stock.Quantity
		r.holdings[userId][stock.TickerSymbol] = oldStock
	}
	return nil
}

func (r *repo) Sell(userId string, stock *models.Stock) error {
	//check if i have already that stock
	oldStock := r.holdings[userId][stock.TickerSymbol]

	if oldStock == nil {
		return fmt.Errorf("This stock you did not buyed")
	} else {
		oldQty := r.holdings[userId][stock.TickerSymbol].Quantity
		if oldQty < stock.Quantity {
			return fmt.Errorf("This Not Enough stock left")
		}
		// oldQty -= stock.Quantity // here io am supposing selling on current average price
		r.holdings[userId][stock.TickerSymbol].Quantity -= stock.Quantity
	}
	return nil
}
func (r *repo) FetchingHoldings(userId string) map[string]*models.Stock {
	return r.holdings[userId]
}
func (r *repo) FetchingReturns(userId string, currentPrice map[string]int) int {
	// check if tht stock exists or not
	var profit int

	currentHolding := r.holdings[userId]

	for _, stock := range currentHolding {
		// get the curent stock -price * quantiy + currentPrice * quantity
		fmt.Printf("currentPrice: %d, oldPrice: %d \n", currentPrice[stock.TickerSymbol], stock.AveragePrice)
		profit += (currentPrice[stock.TickerSymbol] - stock.AveragePrice) * stock.Quantity
	}
	fmt.Printf("Profit is %d\n", profit)
	return profit
}

// map[1][APPL] {}
