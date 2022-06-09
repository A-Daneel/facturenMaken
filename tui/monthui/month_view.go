package monthui

import (
	"fmt"
	"io"

	"github.com/A-Daneel/facturenMaken/tui/constants"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := lipgloss.NewStyle().MarginLeft(2).Render
	if index == m.Index() {
		fn = func(s string) string {
			return lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170")).Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

var cmd tea.Cmd

type BackMsg bool

type Model struct {
	list  list.Model
	items []item
}

func New() tea.Model {
	items := []list.Item{
		item("Januari"),
		item("Februari"),
		item("Maart"),
		item("April"),
		item("Mei"),
		item("Juni"),
		item("July"),
		item("Augustus"),
		item("September"),
		item("Oktober"),
		item("November"),
		item("December"),
	}

	l := list.New(items, itemDelegate{}, 0, 0)
	l.Title = "Main menu"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = lipgloss.NewStyle().MarginLeft(2)
	l.Styles.PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	l.Styles.HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)

	m := Model{
		list:  l,
		items: []item{},
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		top, right, bottom, left := constants.DocStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			// so yeah.... not lovely.. but it works
			switch selection := m.list.Index(); selection {
			case 0:
				print("Januari")
			case 1:
				print("Februari")
			case 2:
				print("Maart")
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return "\n" + m.list.View()
}
