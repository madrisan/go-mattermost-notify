![go-mattermost-notify logo][logo]

![Release Status](https://img.shields.io/badge/status-stable-brightgreen.svg)
[![License](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](https://spdx.org/licenses/Apache-2.0.html)
[![Coverage](https://img.shields.io/badge/Go%20Coverage-50.9%25-green.svg?longCache=true&style=flat)](https://github.com/jpoles1/gopherbadger)
[![Go Report Card](https://goreportcard.com/badge/github.com/madrisan/go-mattermost-notify)](https://goreportcard.com/report/github.com/madrisan/go-mattermost-notify)

# A simple Mattermost notifier written in Go [![Go Reference](https://pkg.go.dev/badge/go-mattermost-notify.svg)](https://pkg.go.dev/github.com/madrisan/go-mattermost-notify)

A Go (golang) simple client for sending [Mattermost](https://mattermost.com/) posts via its REST API v4.
This program makes use of the Go libraries `http` and `url` for interacting with a Mattermost server and
[Cobra](https://cobra.dev/) coupled with [Viper](https://github.com/spf13/viper) to implement the CLI interface.

## Official Docker Image

A docker image is available in Docker Hub: [madrisan/mattermost-notify](https://hub.docker.com/r/madrisan/mattermost-notify)

### How to pull the Docker Image
```
docker pull madrisan/mattermost-notify
```

## Build

### Using go get

```
export GO111MODULE=on
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

### Create a container
```
podman build -t go-mattermost-notify:latest -f deploy/Containerfile .
```
or if you prefer, use `docker` or `nerdctl` instead of *podman*.

If you need to add an extra certificate that is signed by a custom CA (to fix the error message `x509: certificate signed by unknown authority`), do create a file named `additional-ca-cert-bundle.crt` at the root of the project sources and choose the docker file `deploy/Containerfile.additional_ca` instead.
```
podman build -t go-mattermost-notify:latest -f deploy/Containerfile.additional_ca .
```

Now you can load the image created by the previous command with the following command.
```
podman run --rm -it localhost/go-mattermost-notify:latest [add-the-required-options]
```
If you need to debug some issues, you can overwrite the entrypoint with the busybox shell:
```
podman run --entrypoint=sh --rm -it localhost/go-mattermost-notify:latest
```

## Usage

### Post Command

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
  -l, --level string     criticity level. Can be info, success, warning, or critical (default "info")
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

#### Output in Mattermost

As an example we show a Mattermost message using some markdown features (text modifiers, emoticons, and a clickable URL):

![notifications example in Mattermost][example_message]

### Get Command

The `get` command of `go-mattermost-notify` is mainly intended for debugging or for getting Mattemost configuration information.
```
$ go-mattermost-notify get --help
Send a Get query to Mattermost using its REST APIv4 interface.

See the Mattermost API documentation:
  https://api.mattermost.com/

Example:
  get /bots
  get /channels
  get /users/me

Usage:
  go-mattermost-notify get [flags]

Flags:
  -h, --help   help for get

Global Flags:
  -a, --access-token string   Mattermost Access Token. The command-line value has precedence over the MATTERMOST_ACCESS_TOKEN environment variable.
      --config string         config file (default is $HOME/.go-mattermost-notify.yaml)
  -q, --quiet                 quiet mode
  -u, --url string            Mattermost URL. The command-line value has precedence over the MATTERMOST_URL environment variable.
```

## Developers' corner

Some extra actions that may be usefull to project developers.

#### Run the Test Suite

Just run in the top source folder:
```
make test
```
or, if you like a more verbose output:
```
TESTARGS="-test.v" make test
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

[logo]: images/go-mattermost-notify-logo.png?raw=true
[example_message]: images/mattermost_post_example.png?raw=true
