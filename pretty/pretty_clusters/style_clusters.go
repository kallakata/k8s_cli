package pretty_clusters

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/kallakata/k8s_cli/model"
	"strings"
)

const (
	columnKeyCluster       = "cluster"
	columnKeyClusterStatus = "status"
	columnKeyVersion       = "version"
	columnKeyEndpoint      = "endpoint"
)

// const (
// 	columnKeyNodepool    = "nodepool"
// 	columnKeyStatus      = "status"
// 	columnKeyNpVersion   = "version"
// 	columnKeyMinNode     = "minNode"
// 	columnKeyMaxNode     = "maxNode"
// 	columnKeyAutoscaling = "autoscaling"
// )

type Model struct {
	table           table.Model
	selectedCluster model.Cluster
	clusters        []model.Cluster
	selectedIndex   int
}

func (m Model) CreateClusterRows(items []model.Cluster) []table.Row {
	var rows []table.Row
	for _, item := range m.clusters {
		rowData := table.RowData{
			columnKeyCluster:       item.Cluster,
			columnKeyClusterStatus: item.Status,
			columnKeyVersion:       item.Version,
			columnKeyEndpoint:      item.Endpoint,
		}
		row := table.NewRow(rowData)
		rows = append(rows, row)
	}
	return rows
}

// func (m Model) createNodePoolsRows(nodePools []model.Nodepool) []table.Row {
// 	var rows []table.Row
// 	for _, np := range nodePools {
// 		rowData := table.RowData{
// 			columnKeyNodepool:      np.Nodepool,
// 			columnKeyClusterStatus: np.Status,
// 			columnKeyVersion:       np.Version,
// 		}
// 		row := table.NewRow(rowData)
// 		rows = append(rows, row)
// 	}
// 	return rows
// }

var rows []table.Row

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
				Foreground(lipgloss.Color("#00FFFF")).
				Align(lipgloss.Center)),
	}

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
		selectedCluster: model.Cluster{},
		clusters:        items,
		selectedIndex:   -1,
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
		case "b":
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}

	if m.selectedCluster.Cluster != "" {
		// Render the clusters table view
		body.WriteString("List of Clusters in project and zone.\n\n" +
			"Status:\n 0 = Unknown\n 1 = Provisioning\n 2 = Running\n 4 = Stopping\n 5 = Error" +
			"\n\n| Currently filter by Cluster, Status and Version, press / + letters to start filtering, and escape to clear filter. |\n| Press q or ctrl+c to quit | \n\n")
		body.WriteString(m.table.View())
	}

	return body.String()
}
