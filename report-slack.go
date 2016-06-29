package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func report(name string, config Config) {
	api := slack.New(config.SlackAPIKey)
	logger := log.New(os.Stdout, "slack: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	// channel, err := api.GetChannelInfo(config.SlackChannelID)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }

	params := slack.PostMessageParameters{
		Username:  "Playstore update checker",
		IconEmoji: ":android:",
		Attachments: []slack.Attachment{
			slack.Attachment{
				Title:     config.PackageName,
				TitleLink: config.PlayStoreURL,
			},
		},
	}

	if config.Username != "" {
		params.Username = config.Username
	}

	if config.IconURL != "" {
		params.IconURL = config.IconURL
	} else if config.IconEmoji != "" {
		params.IconEmoji = config.IconEmoji
	}

	message := "App update has been deployed!"
	if config.Message != "" {
		message = config.Message
	}

	resChannel, timestamp, err := api.PostMessage(config.SlackChannelID, message, params)
	if err != nil {
		log.Fatalf("error during post message: %v", err)
	}

	log.Printf("posted %v, %v", resChannel, timestamp)
}
