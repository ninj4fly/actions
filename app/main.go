package main

import (
	"fmt"
	"net/http"

	"github.com/oklahomer/go-sarah/v4"
	"github.com/oklahomer/go-sarah/v4/slack"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func setupSlack() {
	// Set up slack adapter.
	slackConfig := slack.NewConfig()
	slackConfig.Token = "REPLACE THIS"
	adapter, err := slack.NewAdapter(slackConfig, slack.WithRTMPayloadHandler(slack.DefaultRTMPayloadHandler))
	if err != nil {
		panic(fmt.Errorf("faileld to setup Slack Adapter: %s", err.Error()))
	}

	// Set up an optional storage so conversational contexts can be stored.
	cacheConfig := sarah.NewCacheConfig()
	storage := sarah.NewUserContextStorage(cacheConfig)

	// Set up a bot with Slack adapter and a default storage.
	bot := sarah.NewBot(adapter, sarah.BotWithStorage(storage))

	sarah.RegisterBot(bot)
}
