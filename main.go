package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/PuerkitoBio/goquery"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"net/http"
	"time"
)


func getTranslation(lang string, query string) string {
	// Request the HTML page.
	fmt.Println("https://slovnik.seznam.cz/preklad/rusky_cesky/" + query)
	res, err := http.Get("https://slovnik.seznam.cz/preklad/rusky_cesky/" + query)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".Box-content-line").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		translation := s.Find("a").Text()
		fmt.Println(translation)
	})

	return ""
}


func main() {

	var botKey string
	flag.StringVar(&botKey, "key", "", "Bot API key")
	flag.Parse()
	fmt.Println(botKey)
	b, err := tb.NewBot(tb.Settings{
		Token: botKey,
		// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
		URL:    "https://api.telegram.org",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/cz", func(m *tb.Message) {
		getTranslation("cz", m.Payload)
		//b.Send(m.Sender, "Echo:" + m.Payload)
	})

	b.Start()
}
