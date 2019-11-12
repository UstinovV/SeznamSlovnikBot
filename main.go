package main

import (
	"log"
	"time"
	"flag"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

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

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Start()
}
