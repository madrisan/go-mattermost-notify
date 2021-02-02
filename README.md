![](images/mattermost_logo.png?raw=true)

![Release Status](https://img.shields.io/badge/status-beta-yellow.svg)
[![License](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](https://spdx.org/licenses/Apache-2.0.html)

# A simple Mattermost notifier written in Go

A Go (golang) simple client for sending [Mattermost](https://mattermost.com/) posts via its REST API v4.

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

```
$ go-mattermost-notify post --help
Post a message to a Mattermost channel or user using its REST APIv4 interface.

Example:
  go-mattermost-notify post rybfbdi9ojy8xxxjjxc88kh3me --author CI --title "Job Status" --message "The job \#BEEF has failed :bug:" --level=critical

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

## Dockerfile

A simple Dockerfile for creating a container providing `go-mattermost-notify` follows.

```
FROM golang:1.15.6-buster as compiler

WORKDIR /usr/local/src

ENV GOPATH /go
RUN mkdir -p $GOPATH/src $GOPATH/bin \
    && chmod -R 755 $GOPATH
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY go-mattermost-notify go-mattermost-notify
RUN make -C go-mattermost-notify dev

FROM alpine:3.13

WORKDIR /usr/local/bin
COPY --from=compiler /go/bin/go-mattermost-notify go-mattermost-notify

CMD ["./go-mattermost-notify"]
```
