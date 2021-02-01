# A simple Mattermost notifier written in Go 

[![License](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](https://spdx.org/licenses/Apache-2.0.html)

A Go (golang) simple client for sending [Mattermost](https://mattermost.com/) posts via its REST API v4.

![](images/mattermost_logo.png?raw=true)

## Build

```
[ "$GOPATH" ] || export GOPATH="$HOME/go"
go get -u github.com/madrisan/go-mattermost-notify

export PATH="$PATH:$GOPATH/bin"
$GOPATH/bin/go-mattermost-notify version
```

## Usage

```
$ go-mattermost-notify post --help
Post a message to a Mattermost channel or user using its REST APIv4 interface.

Example:
  go-mattermost-notify post rybfbdi9ojy8xxxjjxc88kh3me --author CI --title "Job Status" --message "The job \#BEEF has failed :bug:" --level=critical

Usage:
  go-mattermost-notify post [flags]

Flags:
  -A, --author string    author of the message
  -c, --channel string   mattermost channel ID or username. Example: rybfbdi9ojy8xxxjjxc88kh3me or @alice
  -h, --help             help for post
  -l, --level string     criticity level. Can be info (default), success, warning, or critical (default "info")
  -m, --message string   the (markdown-formatted) message to send to the Mattermost channel
  -T, --team string      the mattermot team
  -t, --title string     the title that will precede the text message

Global Flags:
  -a, --access-token string   mattermost Access Token. The command-line value has precedence over the MATTERMOST_ACCESS_TOKEN environment variable.
      --config string         config file (default is $HOME/.go-mattermost-notify.yaml)
  -u, --url string            mattermost URL. The command-line value has precedence over the MATTERMOST_URL environment variable.
```
