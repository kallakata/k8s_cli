package pretty_clusters

import (
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

const (
	columnKeyNodepool    = "nodepool"
	columnKeyStatus 	 = "status"
	columnKeyNpVersion   = "version"
	columnKeyMinNode     = "minNode"
	columnKeyMaxNode     = "maxNode"
	columnKeyAutoscaling = "autoscaling"
)

type Model struct {
	table table.Model
	selectedCluster  model.Cluster
	clusters 		 []model.Cluster
	selectedIndex    int
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

func (m Model) createNodePoolsRows(nodePools []model.Nodepool) []table.Row {
    var rows []table.Row
    for _, np := range nodePools {
        rowData := table.RowData{
            columnKeyNodepool:       np.Nodepool,
            columnKeyClusterStatus: np.Status,
            columnKeyVersion:       np.Version,
        }
        row := table.NewRow(rowData)
        rows = append(rows, row)
    }
    return rows
}

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
		clusters: items,
		selectedIndex: -1,
	}
}

var zone string
var projectID string

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
            // selectedIdx := m.selectedIndex
            // if selectedIdx >= 0 && selectedIdx < len(m.clusters) {
			// 	fmt.Println("Key pressed")
            //     selectedCluster := m.clusters[selectedIdx]
			// 	fmt.Println(selectedCluster)
            //     Fetch node pools for the selected cluster
            //     nodePools, err := fetcher.ListNodepools("devops-internal-t", "europe-west1-b", selectedCluster.Cluster)
            //     if err != nil {
            //         Handle the error
            //         fmt.Println("Error fetching node pools:", err)
            //     } else {
            //         Create node pools rows and update the model
            //         nodePoolsRows := m.createNodePoolsRows(nodePools)
            //         m.table = m.table.WithRows(nodePoolsRows)
            //     }
            // }
        }
    }

    return m, tea.Batch(cmds...)
}

func (m Model) View() string {
    body := strings.Builder{}

    if m.selectedCluster.Cluster != "" {
        // body.WriteString("List of Node Pools in cluster " + m.selectedCluster.Cluster + "\n\n")

        // Create a table for node pools similar to the cluster table
        // nodePoolsColumns := []table.Column{
        //     // Define your node pools table columns here
        //     // e.g., table.NewColumn("columnKeyNodepool", "Node Pool", 20),
        //     //       table.NewColumn("columnKeyStatus", "Status", 15),
        //     //       ...
        // }

        // Fetch node pools for the selected cluster
        // nodePools, err := fetcher.ListNodepools("devops-internal-t", "europe-west1-b", m.selectedCluster.Cluster)
    //     if err != nil {
    //         // Handle the error
    //         body.WriteString("Error fetching node pools: " + err.Error())
    //     } else {
    //         // Create rows for node pools
    //         nodePoolsRows := m.createNodePoolsRows(nodePools)

    //         // Create the node pools table
    //         nodePoolsTable := table.New(nodePoolsColumns).
    //             Filtered(true).
    //             Focused(true).
    //             WithPageSize(10).
    //             WithRows(nodePoolsRows)

    //         // Render the node pools table view
    //         body.WriteString(nodePoolsTable.View())
    //     }
    // } else {
        // Render the clusters table view
        body.WriteString("List of Clusters in project and zone.\n\n" +
            "Status:\n 0 = Unknown\n 1 = Provisioning\n 2 = Running\n 4 = Stopping\n 5 = Error" +
            "\n\n| Currently filter by Cluster, Status and Version, press / + letters to start filtering, and escape to clear filter. |\n| Press q or ctrl+c to quit | \n\n")
        body.WriteString(m.table.View())
    }

    return body.String()
}
