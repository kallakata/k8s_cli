package prompt_nodepools

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	cluster   string
	err       error
}

type ClusterMsg string

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "cluster"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// Capture the entered namespace and return it as a new model
			cluster := m.textInput.Value()
			return model{
				textInput: m.textInput,
				cluster:   cluster,
				err:       nil,
			}, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			os.Exit(1)
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Which cluster do you want to list in?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)\n",
	)
}

func (m model) GetCluster() string {
	return m.cluster
}
