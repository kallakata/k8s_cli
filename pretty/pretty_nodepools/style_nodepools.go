package pretty_nodepools

import (
	// "log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/kallakata/k8s_cli/model"
)

const (
	columnKeyNodepool       = "nodepool"
	columnKeyNodepoolStatus = "status"
	columnKeyNpVersion      = "version"
	columnKeyMinNode        = "min_node"
	columnKeyMaxNode        = "max_node"
	columnKeyAutoscaling    = "autoscaling"
)

type Model struct {
	table table.Model
}

func NewModel(items []model.Nodepool) Model {

	columns := []table.Column{
		table.NewColumn(columnKeyNodepool, "Nodepool", 35).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyNodepoolStatus, "Status", 15).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#8c8")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyNpVersion, "Version", 25).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyAutoscaling, "Autoscaling", 10).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyMinNode, "MinNode", 15).
			WithFiltered(false).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyMaxNode, "MaxNode", 15).
			WithFiltered(false).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
	}

	var rows []table.Row

	for _, item := range items {
		rowData := table.RowData{
			columnKeyNodepool:       item.Nodepool,
			columnKeyNodepoolStatus: item.Status,
			columnKeyNpVersion:      item.Version,
			columnKeyMinNode:        item.MinNode,
			columnKeyMaxNode:        item.MaxNode,
			columnKeyAutoscaling:    item.Autoscaling,
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

	body.WriteString("List of Clusters in project and zone.\n\n" +
		"| Currently filter by Cluster, Status and Version, press / + letters to start filtering, and escape to clear filter. |\n| Press q or ctrl+c to quit | \n\n")

	body.WriteString(m.table.View())

	return body.String()
}
