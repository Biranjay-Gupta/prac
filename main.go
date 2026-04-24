package main

import (
	"fmt"
	"net/http"
	"smallcase/handler"
	"smallcase/repo"
	"smallcase/service"
)

func main() {
	repoLayer := repo.NewRepo()
	serviceLayer := service.NewService(repoLayer)
	handlerLayer := handler.NewHandler(serviceLayer)

	http.HandleFunc("/buy", handlerLayer.Buy)
	http.HandleFunc("/Sell", handlerLayer.Sell)
	http.HandleFunc("/fetch", handlerLayer.FetchingHoldings)
	http.HandleFunc("/return", handlerLayer.FetchingReturns)
	fmt.Println("Server is Runnning on port 8086!")
	http.ListenAndServe(":8086", nil)
}

/*
{
    "TickerSymbol" : "APPL",
    "Quantity":5,
    "AveragePrice":100,
    "StockName":"Apple"
}
*/
