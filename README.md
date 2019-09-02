# Kindling

[![Build Status](https://travis-ci.org/nchaloult/kindling.svg?branch=master)](https://travis-ci.org/nchaloult/kindling)

## Getting Up and Running

1. `docker-compose up -d`
1. `go run main.go`

## Testing

1. `go test -coverprofile cover.out ./...`
1. `go tool cover -html=cover.out -o cover.html`
1. `open cover.html`

# To-Do List

## Refactoring/Restructuring

I have a lot to learn about structuring an app like this.

How I've chosen to organize things is most definitely overkill for a project of this size, but I wanted to practice building something that could go large without having to tweak the structure. As far as I can tell, there isn't a small set of widely-adopted conventions for Go projects, which has made learning Go more challenging (a fun challenge, though 😃).

I learned about dependency injection while working on this project in its initial stages. I'd like to use [wire](https://github.com/google/wire) in this project one day.

[This article](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091) has inspired me to reconsider particular aspects of this project's organization, especially how I hit the database, and how I mock results from functions in `*_repository` files.
