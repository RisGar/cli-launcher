package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type (
	item struct {
		title, desc string
	}
	model struct {
		list list.Model
	}
)

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }
func (m model) Init() tea.Cmd      { return nil }

// Update updates the model based on the given message.
// Returns the updated model and a command to execute.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		// --- Keybinds ---
		case "q", "esc", "ctrl+c":
			// Quit the program
			return m, tea.Quit

		case "enter":
			// Open the selected program if the user presses enter.
			i, ok := m.list.SelectedItem().(item)
			if ok {
				return m, openTUI(i.desc)
			}
			// Quit the program if the user didn't select anything.
			return m, tea.Quit
		}
	}

	// Update the list and return the updated model and command.
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View returns a styled string representation of the model's list.
// The styling is applied using the docStyle renderer.
func (m model) View() string {
	return docStyle.Render(m.list.View())
}

// openTUI runs a command in the current terminal and returns a tea.Cmd.
// The cmd argument is the command to be executed.
func openTUI(cmd string) tea.Cmd {
	c := exec.Command(cmd)

	// Use tea.ExecProcess to create a new tea.Cmd from the exec.Command.
	// If there is an error, it will be returned as a tea.Msg.
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return err
	})
}

func main() {
	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Choose a Terminal Application"

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		// If an error occurs, print it and exit with a non-zero status code.
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
