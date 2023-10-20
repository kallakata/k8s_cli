package pretty_azs_clusters

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/fatih/color"
	"github.com/kallakata/k8s_cli/model"
)

const (
	columnKeyCluster       = "cluster"
	columnKeyClusterStatus = "status"
	columnKeyVersion       = "version"
	columnKeyLocation      = "location"
	columnKeyIdentity      = "identity"
)

type Model struct {
	table   table.Model
	spinner spinner.Model
}

func NewModel(items []model.AzureCluster) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	columns := []table.Column{
		table.NewColumn(columnKeyCluster, "Cluster", 40).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#dd77d5")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyClusterStatus, "Status", 15).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#8c8")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyVersion, "Version", 15).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyVersion, "Location", 15).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff0")).
				Align(lipgloss.Center)),
		table.NewColumn(columnKeyVersion, "Identity", 15).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#dd77d5")).
				Align(lipgloss.Center)),
	}

	var rows []table.Row

	for _, item := range items {
		rowData := table.RowData{
			columnKeyCluster:       item.Cluster,
			columnKeyClusterStatus: item.Status,
			columnKeyVersion:       item.Version,
			columnKeyLocation:      item.Location,
			columnKeyIdentity:      item.Identity,
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
		spinner: s,
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
	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			color.Magenta("\nExiting...\n\n")
			time.Sleep(1 * time.Second)
			cmds = append(cmds, tea.Quit)
		case "enter":
			// m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, m.spinner.Tick)
			m.spinner.View()
			return m, tea.Batch(cmds...)
		}

	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {

	view := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#03a1d3")).Render("| Press / + letters to start filtering by Cluster, and escape to clear filter. | \n"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#42d303")).Render("Press q or ctrl+c to quit\n\n"),
		m.table.View(),
	) + "\n"
	viewH := lipgloss.JoinVertical(
		lipgloss.Right,
	) + "\n"

	return lipgloss.NewStyle().MarginLeft(1).Render(view, viewH)
}
