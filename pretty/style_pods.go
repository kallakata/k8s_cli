package pretty

import (
	// "log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/kallakata/k8s_cli/model"
)

const (
	columnKeyPod    = "pod"
	columnKeyStatus = "status"
	columnKeyNs     = "namespace"
	columnKeyCtx    = "context"
)

func NewPodsModel(items []model.Pod, ctx string, ns string) Model {

	columns := []table.Column{
		table.NewColumn(columnKeyPod, "Pod", 40).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff0")).
			Align(lipgloss.Center)),
		table.NewColumn(columnKeyStatus, "Status", 15).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8c8")).
			Align(lipgloss.Center)),
		table.NewColumn(columnKeyNs, "Namespace", 15).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff0")).
			Align(lipgloss.Center)),
		table.NewColumn(columnKeyCtx, "Context", 30).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
			Faint(true).
			Align(lipgloss.Center)),
	}

	var rows []table.Row

	for _, item := range items {
		rowData := table.RowData{
			columnKeyPod:    item.Pod,
			columnKeyStatus: item.Status,
			columnKeyNs:     ns,
			columnKeyCtx:    ctx,
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

func (m Model) PodsInit() tea.Cmd {
	return nil
}

func (m Model) PodsUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m Model) PodsView() string {
	body := strings.Builder{}

	body.WriteString("List of Pods in namespace and context.\n\n" +
		"| Currently filter by Pod, Status and Namespace, press / + letters to start filtering, and escape to clear filter. |\n| Press q or ctrl+c to quit | \n\n")

	body.WriteString(m.table.View())

	return body.String()
}
