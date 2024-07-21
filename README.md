_This project is replaced by https://github.com/joakim-ribier/gmocky-v2_

# GMOCKY - Generate custom HTTP responses

[![Go Report Card](https://goreportcard.com/badge/github.com/joakim-ribier/gmocky)](https://goreportcard.com/report/github.com/joakim-ribier/gmocky)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-lint.svg)](https://golangci.com)
[![Codeship Status for joakim-ribier/gttp](https://app.codeship.com/projects/b1480cf0-7693-0137-6fa9-3263308262a3/status?branch=master)](https://app.codeship.com/projects/349985)

![gmocky](/resources/gmocky.gif)

## TOC

* [Description](#description)
* [How to use](#how-to-use)
* [Docker](#docker)
* [Troubleshooting](#troubleshooting)

## Description

GMOCKY is a fork of [mocky.io](https://www.mocky.io/) but with a [Go](https://golang.org/) server.

Thanks to [Julien Lafont](https://github.com/julien-lafont/Mocky) for mocky.io :p

**Mock your HTTP responses to test your REST API**

## How to use

1. Build the app
```bash
# Download the Go project
$ go get github.com/joakim-ribier/gmocky

# Build
$ cd $HOME/go/src/gmocky
$ go build
```

2. Start the server
```bash
$ export GMOCKY_PORT=9595
$ ./gmocky start
# Server waiting on :9595...
```
The default port is 8080.

3. Use the helper to generate URL

* Modify the [config.gmocky.json](/resources/config.gmocky.json) file to customize the mocked response.

| Field        | Description |
| ------------ | ----------- |
| status       | status (200, 201, 400, 404...) |
| contentType  | content-type (application/json, text/plain...) |
| charset      | charset of the content-type (utf-8) |
| headers      | custom headers (ETag, If-None-Match, Expires, Last-Modified, Server, X-Cache, Cache-Control...) |
| delay        | delay to simulate a response time ("300ms", "3s") |


* Modify the [content.gmocky.txt](/resources/content.gmocky.txt) file to customize body content of the mocked response.

| Field    | Description |
| -------- | ----------- |
| no-field | plain text can be json, text, xml... |

```bash
# Generate the mocked URL from the 'config.gmocky.json' & 'content.gmocky.txt' files
$ ./gmocky
# Generate mocked URL with encoded params
:8080/?gmocky-content=%7B%0A++%22error%22%3A+%22UNAUTHORIZED%22%2C%0A++%22message%22%3A+%22Wrong+%27X-AUTH-ISSUER%27...%40see+documentation+for+more+details%21%22%0A%7D&gmocky-delay=3s&gmocky-header=%7B%22status%22%3A403%2C%22contentType%22%3A%22application%2Fjson%22%2C%22charset%22%3A%22utf-8%22%2C%22headers%22%3A%7B%22X-AUTH-ISSUER%22%3A%22GMOCKY%22%2C%22X-AUTH-TOKEN%22%3A%22%7Btoken%7D%22%7D%7D

# Generate mocked URL with only encoded 'body' param
:8080/?gmocky-header={"status":403,"contentType":"application/json","charset":"utf-8","headers":{"X-AUTH-ISSUER":"GMOCKY","X-AUTH-TOKEN":"{token}"}}&gmocky-delay=3s&gmocky-content=%7B%0A++%22error%22%3A+%22UNAUTHORIZED%22%2C%0A++%22message%22%3A+%22Wrong+%27X-AUTH-ISSUER%27...%40see+documentation+for+more+details%21%22%0A%7D
```

## Docker

1. Use the remote container

The container is free on [Docker Hub](https://hub.docker.com/r/joakimribier/gmocky).

```bash
# Pull the container
$ docker pull joakimribier/gmocky

# Run the container
docker run -d -p 8585:8585 joakimribier/gmocky
```

2. Build & run local container

```bash
# Build the docker image
$ cd $HOME/go/src/gmocky
$ docker build -t gmocky .

# Run
$ docker run -d -p 8585:8585 gmocky
```

& test the container on
http://localhost:8585/?gmocky-content...

## Troubleshooting

![](https://media0.giphy.com/media/1XgIXQEzBu6ZWappVu/giphy.gif)
