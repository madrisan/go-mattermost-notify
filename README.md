![](images/go-mattermost-notify-logo.png?raw=true)

![Release Status](https://img.shields.io/badge/status-stable-brightgreen.svg)
[![License](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](https://spdx.org/licenses/Apache-2.0.html)
[![Coverage](https://img.shields.io/badge/Go%20Coverage-46.7%25-green.svg?longCache=true&style=flat)](https://github.com/jpoles1/gopherbadger)
[![Go Report Card](https://goreportcard.com/badge/github.com/madrisan/go-mattermost-notify)](https://goreportcard.com/report/github.com/madrisan/go-mattermost-notify)

# A simple Mattermost notifier written in Go [![Go Reference](https://pkg.go.dev/badge/go-mattermost-notify.svg)](https://pkg.go.dev/github.com/madrisan/go-mattermost-notify)

A Go (golang) simple client for sending [Mattermost](https://mattermost.com/) posts via its REST API v4.
This program makes use of the Go libraries `http` and `url` for interacting with a Mattermost server and
[Cobra](https://cobra.dev/) coupled with [Viper](https://github.com/spf13/viper) to implement the CLI interface.

## Build

### Using go get

```
[ "$GOPATH" ] || export GOPATH="$HOME/go"
go get -u github.com/madrisan/go-mattermost-notify

export PATH="$PATH:$GOPATH/bin"
$GOPATH/bin/go-mattermost-notify version
```

### From the source code

```
git clone https://github.com/madrisan/go-mattermost-notify
cd go-mattermost-notify

make dev     # creates the binary for testing the application locally
make bin     # creates the binaries for all the supported OS and architectures
```

## Usage

Use the `post` command of `go-mattermost-notify` to send a message to Mattermost.
```
$ go-mattermost-notify post --help
Post a message to a Mattermost channel or user using its REST APIv4 interface.

Example:
  post -c rybfbdi9ojy8xxxjjxc88kh3me -A CI -t "Job Status" -m "The job \#BEEF has failed :bug:" -l critical
  post -c @alice -A CI -t "Job Status" -m "The job \#BEEF ended successfully :tada:" -l success

Usage:
  go-mattermost-notify post [flags]

Flags:
  -A, --author string    author of the message
  -c, --channel string   Mattermost channel ID or username. Example: rybfbdi9ojy8xxxjjxc88kh3me or @alice
  -h, --help             help for post
  -l, --level string     criticity level. Can be info (default), success, warning, or critical (default "info")
  -m, --message string   the (markdown-formatted) message to send to the Mattermost channel
  -T, --team string      the Mattermost team
  -t, --title string     the title that will precede the text message

Global Flags:
  -a, --access-token string   Mattermost Access Token. The command-line value has precedence over the MATTERMOST_ACCESS_TOKEN environment variable.
      --config string         config file (default is $HOME/.go-mattermost-notify.yaml)
  -q, --quiet                 quiet mode
  -u, --url string            Mattermost URL. The command-line value has precedence over the MATTERMOST_URL environment variable.
```

The *access token* and the *url* can be set using different methods:
 * At command-line (`--access-token` and `--url` respectively)
 * By setting the environment variables `MATTERMOST_ACCESS_TOKEN` and `MATTERMOST_URL`
 * By creating a yaml configuration file (by default `~/.go-mattermost-notify.yaml`) containing the lines:
```
mattermost:
  access-token: <access-token>
  url: <base URL of the Mattermost server>
```

The precedence order is: **flags > environment variables > configuration file**.

### Example of the output

![](images/mattermost_posts_example.png?raw=true)

## Developers' corner

Some extra actions that may be usefull to project developers.

#### Run the Test Suite

Just run in the top source folder:
```
make test
```

#### Generate Test Coverage Statistics

Go to the top source folder and enter the command:
```
make cover
```

#### Style and Static Code Analyzers

##### GolangCI-Lint

Run the `GolangCI-Lint` linters aggregator:
```
make lint
```

##### Go Vet

Run the Go source code static analysis tool `vet` to find any common errors:
```
make vet
```
