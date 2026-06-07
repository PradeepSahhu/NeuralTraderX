package news

import "net/http"

func GoogleRssFeed(name string) {
	resp, err := http.Get("https://news.google.com/rss/search?q=RELIANCE+stocks&hl=en-IN&gl=IN&ceid=IN:en")

}
