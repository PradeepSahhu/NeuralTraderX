package service

import (
	"fmt"

	news "github.com/PradeepSahhu/NeuralTraderX/scripts/Internal/News"
	StockNewsRepository "github.com/PradeepSahhu/NeuralTraderX/scripts/Internal/Repository/stocknews"
	repository "github.com/PradeepSahhu/NeuralTraderX/scripts/Internal/Repository/stocks"
	"gorm.io/gorm"
)

func NewsProcessing(db *gorm.DB) {

	stockRepo := repository.New(db)
	stockNewsRepo := StockNewsRepository.New(db)
	stocks, err := stockRepo.GetAllStocks()

	if err != nil {
		fmt.Println("Failed to get all the stocks")
		return
	}

	for _, stock := range stocks {

		stockNews, err := news.GoogleRssFeed(stock.Id, stock.Company_name, stock.Symbol)
		if err != nil {
			fmt.Println("Something went wrong skipping the " + stock.Symbol)
			continue

		}

		if len(stockNews) == 0 {
			fmt.Println("No recent news found for", stock.Symbol)
			continue
		}

		success, err := stockNewsRepo.InsertBatchStockNews(stockNews)

		if err != nil {
			fmt.Println("something went wrong")
			return
		}

		if success {
			fmt.Println("Successfully inserted batch of", stock.Symbol)
		}

	}

}
