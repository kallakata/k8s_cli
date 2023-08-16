package pretty

import (
	// "log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
	"github.com/kallakata/k8s_cli/model"
	"github.com/charmbracelet/lipgloss"
)

const (
	columnKeyNamespace       = "namespace"
	columnKeyContext = "context"
)

type Model struct {
	table table.Model
}

func NewNsModel(items []model.Ns, ctx string) Model {
	columns := []table.Column{
		table.NewColumn(columnKeyNamespace, "Namespace", 70).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff0")).
			Align(lipgloss.Center)),
		table.NewColumn(columnKeyContext, "Context", 70).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
			Faint(true).
			Align(lipgloss.Center)),
	}

	var rows []table.Row

	for _, item := range items {
		rowData := table.RowData{
			columnKeyNamespace:  item.Namespace,
			columnKeyContext: ctx,
		}
		row := table.NewRow(rowData)
		rows = append(rows, row)
	}

	return Model{
		table: table.
			New(columns).
			Filtered(true).
			Focused(true).
			WithPageSize(10).
			WithRows(rows),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			cmds = append(cmds, tea.Quit)
		}

	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString("List of namespaces in context.\n\n" +
		"| Currently filter by Namespace, press / + letters to start filtering, and escape to clear filter. |\n| Press q or ctrl+c to quit | \n\n")

	body.WriteString(m.table.View())

	return body.String()
}
