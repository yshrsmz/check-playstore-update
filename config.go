package main

// Config for the app
type Config struct {
	SlackAPIKey    string `yaml:"slack_api_key"`
	SlackChannelID string `yaml:"slack_channel_id"`
	Username       string `yaml:"username"`
	IconURL        string `yaml:"icon_url"`
	IconEmoji      string `yaml:"icon_emoji"`
	StartMessage   string `yaml:"start_message"`
	SuccessMessage string `yaml:"success_message"`
	ErrorMessage   string `yaml:"error_message"`
	SleepTime      int    `yaml:"sleep_time"`
}

// AppConfig include Config and cli's local config
type AppConfig struct {
	Params       Config
	PackageName  string
	PlayStoreURL string
}
