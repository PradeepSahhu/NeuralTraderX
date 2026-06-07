package news

import (
	"fmt"
	"net/url"

	"github.com/mmcdole/gofeed"
)

func GoogleRssFeed(stock, symbol string) (string, error) {
	fp := gofeed.NewParser()

	query := url.QueryEscape(stock + " stocks or " + symbol + " stock symbol")

	rssUrl := fmt.Sprintf("https://news.google.com/rss/search?q=%s&hl=en-IN&gl=IN&ceid=IN:en", query)

	feeds, err := fp.ParseURL(rssUrl)

	if err != nil {
		fmt.Println("Something went wrong")
	}

	// fmt.Printf("%+v\n", *feeds.Items[0])

	// for _, f := range feeds.Items {
	// 	fmt.Print(f.)
	// 	fmt.Println(f.Title)
	// 	fmt.Println(f.Description)
	// 	fmt.Println(f.Link)
	// 	fmt.Println("-------------------------------")
	// }

	fmt.Println(feeds)

	return "", nil

}
