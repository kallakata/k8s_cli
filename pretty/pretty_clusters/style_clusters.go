package pretty_clusters

import (
	// "log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/kallakata/k8s_cli/model"
)

const (
	columnKeyCluster       = "cluster"
	columnKeyClusterStatus = "status"
	columnKeyVersion       = "version"
	columnKeyEndpoint      = "endpoint"
)

type Model struct {
	table table.Model
}

func NewModel(items []model.Cluster) Model {

	columns := []table.Column{
		table.NewColumn(columnKeyCluster, "Cluster", 35).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyClusterStatus, "Status", 15).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#8c8")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyVersion, "Version", 25).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyEndpoint, "Endpoint", 15).
			WithFiltered(false).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
	}

	var rows []table.Row

	for _, item := range items {
		rowData := table.RowData{
			columnKeyCluster:       item.Cluster,
			columnKeyClusterStatus: item.Status,
			columnKeyVersion:       item.Version,
			columnKeyEndpoint:      item.Endpoint,
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
