package prompt_pods

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
	namespace string
	err       error
}

type NamespaceMsg string

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "namespace"
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
			namespace := m.textInput.Value()
			return model{
				textInput: m.textInput,
				namespace: namespace,
				err:       nil,
			}, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			os.Exit(1)
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Which namespace do you want to list in?\n\n%s\n\nPress Enter for all namespaces\n%s",
		m.textInput.View(),
		"(esc to quit)\n",
	)
}

func (m model) GetNamespace() string {
	return m.namespace
}
