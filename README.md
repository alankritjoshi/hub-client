# hub-client

## Simple TCP-based chat client. Server can be found: https://github.com/alankritjoshi/hub-server

## Setup
1. `go run client.go 127.0.0.1:1234`
2. Use input:
    1. "whoami" -> returns clientID of the requesting client
    2. "whoelse" -> returns space-separated clientIDs connected to server apart from requesting client
    3. "send [message] [clientID 1] [clientID 2] ... [clientID 3]" -> sends [message] to specified space-separated clientIDs
