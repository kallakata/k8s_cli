package pretty_nodepools

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/kallakata/k8s_cli/model"
	"strings"
)

const (
	columnKeyNodepool       = "nodepool"
	columnKeyNodepoolStatus = "status"
	columnKeyNpVersion      = "version"
	columnKeyMinNode        = "min_node"
	columnKeyMaxNode        = "max_node"
	columnKeyAutoscaling    = "autoscaling"
	columnKeyNodeCount      = "node_count"
)

type Model struct {
	table table.Model
}

func NewModel(nodepools []model.Nodepool) Model {

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
				Foreground(lipgloss.Color("#cd03d3")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyNodeCount, "NodeCount", 25).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#00FFFF")).
				Align(lipgloss.Center)),
	}

	var rows []table.Row

	for _, nodepool := range nodepools {
		rowData := table.RowData{
			columnKeyNodepool:       nodepool.Nodepool,
			columnKeyNodepoolStatus: nodepool.Status,
			columnKeyNpVersion:      nodepool.Version,
			columnKeyNodeCount:      nodepool.NodeCount,
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

	body.WriteString("List of Nodepools in project, zone and cluster.\n\n" +
		"| Currently filter by Nodepool, press / + letters to start filtering, and escape to clear filter. |\n| Press q or ctrl+c to quit | \n\n")

	body.WriteString(m.table.View())

	return body.String()
}
