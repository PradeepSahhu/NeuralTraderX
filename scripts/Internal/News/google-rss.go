package news

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/PradeepSahhu/NeuralTraderX/scripts/Internal/models"
	"github.com/mmcdole/gofeed"
)

func GoogleRssFeed(id uint, stock, symbol string) ([]*models.StockNews, error) {

	fp := gofeed.NewParser()

	query := url.QueryEscape(stock + " stocks or " + symbol + " stock symbol")

	rssUrl := fmt.Sprintf("https://news.google.com/rss/search?q=%s&hl=en-IN&gl=IN&ceid=IN:en", query)

	feeds, err := fp.ParseURL(rssUrl)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", *feeds.Items[0])

	// stockNews := mapToStockNews(feeds.Items, id)
	// 	fmt.Print(f.)
	// 	fmt.Println(f.Title)
	// 	fmt.Println(f.Description)
	// 	fmt.Println(f.Link)
	// 	fmt.Println("-------------------------------")

	// fmt.Println(feeds)

	// return stockNews, nil
	return nil, nil

}

func mapToStockNews(feeds []*gofeed.Item, stockId uint) []*models.StockNews {
	var stockNews []*models.StockNews

	oneWeekBefore := time.Now().AddDate(0, 0, -7)
	for _, feed := range feeds {
		if feed.PublishedParsed == nil || feed.PublishedParsed.Before(oneWeekBefore) {
			continue
		}
		source := feed.Title[strings.LastIndex(feed.Title, "-")+1:]

		news := &models.StockNews{
			Title:       feed.Title,
			Description: feed.Description,
			Link:        feed.Link,
			Source:      source,
			StockId:     stockId,
			StoredAt:    time.Now(),
		}

		if feed.PublishedParsed != nil {
			news.PublishedAt = *feed.PublishedParsed
		} else {
			news.PublishedAt = time.Now()
		}

		stockNews = append(stockNews, news)

	}

	return stockNews
}
