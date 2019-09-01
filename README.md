# Kindling

[![Build Status](https://travis-ci.org/nchaloult/kindling.svg?branch=master)](https://travis-ci.org/nchaloult/kindling)

## Getting Up and Running

1. `docker-compose up -d`
1. `go run main.go`

## Testing

1. `go test -coverprofile cover.out ./...`
1. `go tool cover -html=cover.out -o cover.html`
1. `open cover.html`
