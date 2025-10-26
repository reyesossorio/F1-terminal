# F1 Terminal Results Viewer

A simple terminal-based application written in Go that fetches and displays the latest Formula 1 session results.  
Get a quick view of race, qualifying, or practice session standings directly in your terminal.

## Features

- Fetch latest F1 session results (race, qualifying, practice) using the OpenF1 API.
- Display driver positions, teams, and times in a clean terminal table.
- Highlight gaps or fastest laps (optional enhancement).
- Easy-to-use terminal interface with minimal setup.

## Example Output
```

Latest Session Results: Bahrain GP - Race

POS DRIVER TEAM TIME
1 VER Red Bull 1:32:01.123
2 HAM Mercedes +1.256s
3 LEC Ferrari +5.432s
4 SAI Ferrari +8.123s
...

````

## Getting Started

### Prerequisites

- Go 1.20+ installed: [Download Go](https://golang.org/dl/)
- Terminal that supports ANSI escape codes (most modern terminals do)

### Installation

```bash
git clone https://github.com/yourusername/f1-terminal.git
cd f1-terminal
go mod tidy
go run ./cmd/f1-terminal
````

### Configuration

If needed, configure API endpoints or refresh intervals in `internal/config/config.go`.

---

## Project Structure

```
f1-terminal/
├── cmd/               # Main application entry point
│   └── main.go
├── internal/
│   ├── api/           # API client 
│   ├── models/        # Models
│   ├── ui/            # Terminal rendering
│   └── config/        # Configurations
├── go.mod
└── README.md
```
