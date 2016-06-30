Check Playstore update
===

cli tool to check Android app's update and report to Slack.

## Usage

```shell
$ bin/check-playstore-update-linux-386 -h
NAME:
   check-playstore-update - Check if android app update is deployed to PlayStore

USAGE:
   check-playstore-update-linux-386 [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR(S):
   yshrsmz <the.phantom.bane@gmail.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name value, -n value    app package name to check.
   --config value, -c value  specify config file. (default: "config.yml")
   --help, -h                show help
   --version, -v             print the version
```

## Configuration

You can specify these values by config.yml.
checkout `config.yml.template` for the detail.

```yaml
# slack api config
slack_api_key: "your slack token"
slack_channel_id: "slack channel id to post"

# bot user config
icon_url: "icon url, empty if not needed(icon_emoji will be used)."
icon_emoji: "icon emoji, default to :android:"
username: "default to Playstore update checker"
start_message: "message when the polling starts"
success_error: "message when the update found"
error_message: "message when failed to find update"

# cli config
sleep_time: 60 # time in seconds. defaults to 60
```

To try this tool, you can generate your test slack token [here](https://api.slack.com/web)
And you can get channel id from [here](https://api.slack.com/methods/channels.list/test)


## License

    Copyright 2016 Shimizu Yasuhiro (yshrsmz)

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
