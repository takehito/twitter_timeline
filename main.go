package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dghubble/oauth1"
)

type UserInfo struct {
	Name string `json:"name"`
}

type Tweet struct {
	User UserInfo `json:"user"`
	Text string   `json:"text"`
	When string   `json:"created_at"`
}

type tweets []Tweet

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	count := flag.Int("count", 50, "表示したいツイート数を入力して下さい")
	flag.Parse()

	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	path := fmt.Sprintf("https://api.twitter.com/1.1/statuses/home_timeline.json?count=%d", *count)
	resp, err := httpClient.Get(path)
	errCheck(err)
	defer resp.Body.Close()

	var tweets tweets
	reader := resp.Body
	err = json.NewDecoder(reader).Decode(&tweets)
	errCheck(err)

	for i := len(tweets); i > 0; i-- {
		tweet := tweets[i-1]
		when, err := time.Parse(time.RubyDate, tweet.When)
		errCheck(err)

		before := time.Since(when)
		var fmtBefore string
		switch {
		case before.Hours() >= 1:
			fmtBefore = fmt.Sprintf("%d時間前", int(before.Hours()))
		case before.Minutes() >= 1:
			fmtBefore = fmt.Sprintf("%d分前", int(before.Minutes()))
		default:
			fmtBefore = fmt.Sprintf("%d秒前", int(before.Seconds()))
		}

		fmt.Printf("\x1b[33m%s\x1b[0m: %s\n", tweet.User.Name, fmtBefore)
		fmt.Printf("%s\n\n", tweet.Text)
	}
}
