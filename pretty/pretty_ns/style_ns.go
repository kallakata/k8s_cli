package pretty_ns

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/fatih/color"
	"github.com/kallakata/k8s_cli/model"
	"time"
)

const (
	columnKeyNamespace = "namespace"
	columnKeyContext   = "context"
	columnKeyPods      = "pods"
)

type Model struct {
	table table.Model
}

func NewModel(items []model.Ns, ctx string) Model {
	columns := []table.Column{
		table.NewColumn(columnKeyNamespace, "Namespace", 40).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyPods, "Pods", 40).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#dd77d5")).
				Align(lipgloss.Center)),
	}

	var rows []table.Row

	for _, item := range items {
		rowData := table.RowData{
			columnKeyNamespace: item.Namespace,
			columnKeyContext:   ctx,
			columnKeyPods:      item.Pods,
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
			color.Magenta("\nExiting...\n\n")
			time.Sleep(1 * time.Second)
			cmds = append(cmds, tea.Quit)
		}

	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	footer := m.table.HighlightedRow().Data[columnKeyContext].(string)

	view := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#03a1d3")).Render("| Press / + letters to start filtering by Namespace, and escape to clear filter. | \n"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#42d303")).Render("Press q or ctrl+c to quit\n\n"),
		m.table.View(),
	) + "\n"
	viewH := lipgloss.JoinVertical(
		lipgloss.Right,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#42d303")).Render("Looking at context: "+footer),
	) + "\n"

	return lipgloss.NewStyle().MarginLeft(1).Render(view, viewH)
}
