package service

import (
	"fmt"
	"io"
	"net/http"

	"gorm.io/gorm"
)

func NewsProcessing(db *gorm.DB) {

	// stockRepo := repository.New(db)
	// stockNewsRepo := StockNewsRepository.New(db)
	// stocks, err := stockRepo.GetAllStocks()

	// if err != nil {
	// 	fmt.Println("Failed to get all the stocks")
	// 	return
	// }

	// for _, stock := range stocks {

	// 	stockNews, err := news.GoogleRssFeed(stock.Id, stock.Company_name, stock.Symbol)
	// 	if err != nil {
	// 		fmt.Println("Something went wrong skipping the " + stock.Symbol)
	// 		continue

	// 	}

	// 	if len(stockNews) == 0 {
	// 		fmt.Println("No recent news found for", stock.Symbol)
	// 		continue
	// 	}

	// 	success, err := stockNewsRepo.InsertBatchStockNews(stockNews)

	// 	if err != nil {
	// 		fmt.Println("something went wrong")
	// 		return
	// 	}

	// 	if success {
	// 		fmt.Println("Successfully inserted batch of", stock.Symbol)
	// 	}

	// }

	//Links are inside the stockNews, get it by stock_id (batch) and store the consolidated news (description) in the stocksentiment
	// and for now use the llm for the sentiment score based on description.
	//TODO: create multiple news sources for better pridiction, and also create a sentiment model

	// and then again use the llm to get the top 10 among them and get them and calculate the final score

	// for _, stock := range stocks {

	// latestStockNews, err := stockNewsRepo.GetBatchNews()

	// if err != nil {
	// 	fmt.Println("Failed to get the Latest Stock News from the db", err)
	// 	return
	// }

	// if len(latestStockNews) == 0 {
	// 	fmt.Println("There is no latest news of this stock")
	// 	return
	// }

	url := "https://simplywall.st/stocks/in/healthcare/nse-apollohosp/apollo-hospitals-enterprise-shares/news/earnings-update-apollo-hospitals-enterprise-limited-nseapoll"

	client := &http.Client{}
	req, errs := http.NewRequest("GET", url, nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 Chrome/137.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	// resp, err := http.Get("https://simplywall.st/stocks/in/healthcare/nse-apollohosp/apollo-hospitals-enterprise-shares/news/earnings-update-apollo-hospitals-enterprise-limited-nseapoll")

	// for _, s := range latestStockNews {
	// 	resp, err := http.Get(s.Link)

	if err != nil {
		fmt.Println("unable to get the data")

	}

	defer resp.Body.Close()

	body, errs := io.ReadAll(resp.Body)
	if errs != nil {
		fmt.Println("unable to convert")
	}

	content := string(body)
	fmt.Println(content)
	fmt.Println(resp.Request.URL.String())

	// 	break
	// }
	// }

}

// Biggest hurdle is to decode the original url from the google news feed Link and store it in the DB.
