# TermTalk
A command-line chat server built with Go. TermTalk gets you chatting fast -- with no external dependencies and minimal setup.

## Why TermTalk?
Sometimes you just want to chat — no accounts, no web browser, no cloud, no clutter. TermTalk is a clean, command-line chat server that runs locally and lets you and others on your LAN exchange messages instantly. No setup, no tracking, no third-party services.

That’s why I built TermTalk: to have a fast, private, and minimal way to chat — right from the terminal.

## Quick Start

### Install TermTalk using the Go toolchain
```bash
go install github.com/sternerr/TermTalk@latest
```

## Run TermTalk as a server or client
```bash
# Running TermTalk as a server
termtalk host --port <port>

# Running a client
termtalk client --port <port> 
```
and start chatting

