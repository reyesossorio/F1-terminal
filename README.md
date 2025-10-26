# F1 Terminal Tracker

A terminal-based F1 race tracker written in Go.  
Follow your favorite drivers in real-time with live positions, gaps, and lap updates directly in your terminal.

## Features

- Real-time display of all drivers in a race.
- Show current position, lap, and gap to the car ahead.
- Smooth terminal UI with optional color highlights.
- Configurable refresh interval and favorite driver highlighting.

## Demo

```

POS DRIVER      LAP  GAP
1   VER        12   -
2   HAM        12   +1.2s
3   LEC        12   +0.8s
4   SAI        12   +1.1s
...

````

## Getting Started

### Prerequisites

- [Go 1.20+](https://golang.org/dl/) installed
- Terminal that supports ANSI escape codes (most modern terminals do)

### Installation

```bash
git clone https://github.com/reyesossorio/f1-terminal.git
cd f1-terminal
go mod tidy
go run ./cmd/f1-terminal
````

---

## Project Structure

```
f1-terminal/
├── cmd/               # Main application entry point
├── internal/
│   ├── api/           # API client & models
│   ├── race/          # Race state and updater logic
│   ├── ui/            # Terminal rendering
│   └── config/        # Configurations
├── go.mod
└── README.md
```
