package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

var api *slack.Client

func reportStart(config AppConfig) {
	api = getAPI(config)

	params := createDefaultParameters(config)
	message := "Start polling update..."
	if config.Params.StartMessage != "" {
		message = config.Params.StartMessage
	}

	resChannel, timestamp, err := api.PostMessage(config.Params.SlackChannelID, message, params)
	if err != nil {
		log.Fatalf("error during post message: %v", err)
	}
	log.Printf("posted %v, %v", resChannel, timestamp)
}

func reportSuccess(config AppConfig) {
	api = getAPI(config)

	params := createDefaultParameters(config)
	message := "App update has been deployed!"
	if config.Params.SuccessMessage != "" {
		message = config.Params.SuccessMessage
	}

	resChannel, timestamp, err := api.PostMessage(config.Params.SlackChannelID, message, params)
	if err != nil {
		log.Fatalf("error during post message: %v", err)
	}

	log.Printf("posted %v, %v", resChannel, timestamp)
}

func reportError(config AppConfig) {
	api = getAPI(config)

	params := createDefaultParameters(config)
	message := "polling failed somehow :("
	if config.Params.ErrorMessage != "" {
		message = config.Params.ErrorMessage
	}

	resChannel, timestamp, err := api.PostMessage(config.Params.SlackChannelID, message, params)
	if err != nil {
		log.Fatalf("error during post message: %v", err)
	}
	log.Printf("posted %v, %v", resChannel, timestamp)
}

func createDefaultParameters(config AppConfig) slack.PostMessageParameters {
	params := slack.PostMessageParameters{
		Username: "Playstore update checker",
		Attachments: []slack.Attachment{
			slack.Attachment{
				Title:     config.PackageName,
				TitleLink: config.PlayStoreURL,
			},
		},
	}

	if config.Params.Username != "" {
		params.Username = config.Params.Username
	}

	if config.Params.IconURL != "" {
		params.IconURL = config.Params.IconURL
	} else if config.Params.IconEmoji != "" {
		params.IconEmoji = config.Params.IconEmoji
	} else {
		params.IconEmoji = ":android:"
	}

	return params
}

func getAPI(config AppConfig) *slack.Client {
	if api == nil {
		log.Println("create new api instance")
		api = slack.New(config.Params.SlackAPIKey)

		logger := log.New(os.Stdout, "slack: ", log.Lshortfile|log.LstdFlags)
		slack.SetLogger(logger)
		api.SetDebug(true)
	}
	return api
}
