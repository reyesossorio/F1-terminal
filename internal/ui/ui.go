package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/reyesossorio/f1-terminal/internal/domain"
	"github.com/reyesossorio/f1-terminal/internal/service"
)

type Model struct {
	service *service.RaceService
	drivers []*domain.DriverInfo
	session *domain.SessionResult
	loading bool
	err     error
}

// Constructor: recibe la instancia de service
func New(service *service.RaceService) Model {
	return Model{
		service: service,
		loading: true,
	}
}

// Mensajes internos
type (
	loadedMsg struct{}
	errMsg    struct{ err error }
)

// Init se ejecuta al iniciar el programa
func (m Model) Init() tea.Cmd {
	// Cargar datos de manera asíncrona
	return func() tea.Msg {
		err := m.service.SaveLatestSession()
		if err != nil {
			return errMsg{err}
		}

		err = m.service.LazyDriversRaceResults(10, false)
		if err != nil {
			return errMsg{err}
		}
		driversNumbers := m.service.GetDriversNumbersFromLastSession(10, false)
		err = m.service.LazyDriversInfo(driversNumbers)
		if err != nil {
			return errMsg{err}
		}
		return loadedMsg{}
	}
}

// Update maneja los eventos
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case loadedMsg:
		m.drivers = m.service.GetDriversInSession()
		m.session = m.service.GetSessionResult()
		m.loading = false

	case errMsg:
		m.err = msg.err
		m.loading = false
	}

	return m, nil
}

// View define qué se muestra en la terminal
func (m Model) View() string {
	if m.loading {
		return "⏳Loading last session info...\n"
	}

	if m.err != nil {
		return fmt.Sprintf("❌ Error: %v\nPress 'q' to exit.\n", m.err)
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("🏎️ Latest Session Results: %s - %s\n\n", m.session.Circuit, m.session.SessionName))
	b.WriteString("POS	DRIVER	TEAM\n")

	for _, d := range m.drivers {
		b.WriteString(fmt.Sprintf("%d	%s	%s\n", d.DriverNumber,  d.Name, d.TeamName))
	}

	b.WriteString("\nPress 'q' to exit.\n")

	return b.String()
}
