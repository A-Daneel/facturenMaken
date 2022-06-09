package tui

import (
	"fmt"
	"log"
	"os"

	"github.com/A-Daneel/facturenMaken/tui/monthui"
	"github.com/A-Daneel/facturenMaken/tui/projectui"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	p *tea.Program
)

type sessionState int

const (
	projectView sessionState = iota
	monthView
)

type MainModel struct {
	state           sessionState
	project         tea.Model
	activeSelection uint
}

func StartTea() {
	if os.Getenv("HELP_DEBUG") != "" {
		if f, err := tea.LogToFile("debug.log", "help"); err != nil {
			fmt.Println("Could't open a file for logging:", err)
			os.Exit(0)
		} else {
			defer func() {
				err = f.Close()
				if err != nil {
					log.Fatal(err)
				}
			}()
		}
	}
	m := New()
	p = tea.NewProgram(m)
	p.EnterAltScreen()
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func New() MainModel {
	return MainModel{
		state:   projectView,
		project: projectui.New(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg.(type) {
	//switching views, hopefully
	case projectui.SelectMsg:
		print("ello?")
		m.state = monthView
	}

	switch m.state {
	case monthView:
		mewMonth, newCmd := m.project.Update(msg)
		projectModel, ok := mewMonth.(monthui.Model)
		if !ok {
			panic("could not perform assertion on projectui model")
		}
		m.project = projectModel
		cmd = newCmd
	case projectView:
		newProject, newCmd := m.project.Update(msg)
		projectModel, ok := newProject.(projectui.Model)
		if !ok {
			panic("could not perform assertion on projectui model")
		}
		m.project = projectModel
		cmd = newCmd
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	switch m.state {
	case monthView:
		print("we should goooooo")
		return m.project.View()
	default:
		return m.project.View()
	}
}
