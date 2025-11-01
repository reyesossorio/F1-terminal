package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/reyesossorio/f1-terminal/internal/service"
	"github.com/reyesossorio/f1-terminal/internal/storage"
	"github.com/reyesossorio/f1-terminal/internal/ui"
)

func main() {
	drivers := storage.NewDriverStorage()
	sessions := storage.NewSessionStorage()

	service := service.NewRaceService(drivers, sessions)

	model := ui.New(service)
	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("❌ Error ejecutando la UI:", err)
		os.Exit(1)
	}
}
