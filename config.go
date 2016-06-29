package main

// Config for the app
type Config struct {
	SlackAPIKey    string `yaml:"slack_api_key"`
	SlackChannelID string `yaml:"slack_channel_id"`
	Username       string `yaml:"username"`
	IconURL        string `yaml:"icon_url"`
	IconEmoji      string `yaml:"icon_emoji"`
	Message        string `yaml:"message"`
	PackageName    string
	PlayStoreURL   string
}
