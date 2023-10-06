package pretty_pods

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/kallakata/k8s_cli/model"
	"github.com/charmbracelet/bubbles/spinner"
	"fmt"
	"time"
)

var (
	styleSubtle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5733"))

	styleBase = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff5733")).
			BorderForeground(lipgloss.Color("#ff5733")).
			Align(lipgloss.Right)
)

const (
	columnKeyPod    = "pod"
	columnKeyStatus = "status"
	columnKeyNs     = "namespace"
	columnKeyCtx    = "context"
	columnKeyCPUreq = "CPU requests"
	columnKeyCPUlim = "CPU limits"
	columnKeyMemReq = "Mem requests"
	columnKeyMemLim = "Mem limits"
	columnKeyImage 	= "Image"
)

type Model struct {
	table table.Model
	spinner spinner.Model
}

func NewModel(items []model.Pod, ctx string, ns string) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	columns := []table.Column{
		table.NewColumn(columnKeyPod, "Pod", 55).
			WithFiltered(true).
			WithStyle(lipgloss.NewStyle().
				Foreground(lipgloss.Color("#dd77d5")).
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
	}

	var rows []table.Row

	for _, item := range items {
		rowData := table.RowData{
			columnKeyPod:    item.Pod,
			columnKeyStatus: item.Status,
			columnKeyNs:     item.Namespace,
			columnKeyCPUreq: item.CPUReq,
			columnKeyCPUlim: item.CPULim,
			columnKeyMemReq: item.MemReq,
			columnKeyMemLim: item.MemLim,
			columnKeyImage: item.Image,
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
			fmt.Printf("\nExiting...\n\n")
			time.Sleep(1 * time.Second)
			cmds = append(cmds, tea.Quit)
		case "enter":
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, m.spinner.Tick)
			m.spinner.View()
			return m, tea.Batch(cmds...)
		}

	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	// body := strings.Builder{}

	// body.WriteString("List of Pods in namespace and context.\n\n" +
	// 	"| Currently filter by Pod, Status and Namespace, press / + letters to start filtering, and escape to clear filter. |\n| Press q or ctrl+c to quit | \n\n")

	// body.WriteString(m.table.View())

	// return body.String()
	highlightedRow := m.table.HighlightedRow()

	footerText := fmt.Sprintf(
		"Pg. %d/%d \n Currently looking at context: %s",
		m.table.CurrentPage(),
		m.table.MaxPages(),
		highlightedRow.Data[columnKeyCtx],
	)
	m.table = m.table.WithStaticFooter(footerText)
	selected := m.table.HighlightedRow().Data[columnKeyPod].(string)
	rq_cpu := m.table.HighlightedRow().Data[columnKeyCPUreq].(string)
	rq_mem := m.table.HighlightedRow().Data[columnKeyMemReq].(string)
	lim_cpu := m.table.HighlightedRow().Data[columnKeyCPUlim].(string)
	lim_mem := m.table.HighlightedRow().Data[columnKeyMemLim].(string)
	image := m.table.HighlightedRow().Data[columnKeyImage].(string)
	view := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#03a1d3")).Render("| Currently filter by Pod, Status and Namespace, press / + letters to start filtering, and escape to clear filter. | \n"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#42d303")).Render("Press q or ctrl+c to quit\n\n"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#77d5dd")).Render("Pod: "+selected),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#77d5dd")).Render("Image: "+image),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5733")).Render("Requests CPU: "+rq_cpu, "/", "Limits CPU: "+lim_cpu),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5733")).Render("Requests Mem: "+rq_mem, "/", "Limits Mem: "+lim_mem),
		m.table.View(),
	) + "\n"

	return lipgloss.NewStyle().MarginLeft(1).Render(view)
}
