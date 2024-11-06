package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dan.frohlch/tealayout"
)

func main() {
	const (
		variableWidthHeight = 0
		fixedWidth          = 24
		fixedHeight         = 4
	)
	bl := tealayout.NewBorderLayout(
		tealayout.BorderRegionConfigOpt(tealayout.NorthBorderRegion, textModel{text: "North component"}, variableWidthHeight, fixedHeight),
		tealayout.BorderRegionConfigOpt(tealayout.WestBorderRegion, textModel{text: "West component"}, fixedWidth, variableWidthHeight),
		tealayout.BorderRegionConfigOpt(tealayout.CenterBorderRegion,
			textModel{text: "Center component\n  - type [n/s/e/w/c] to show/hide regions"},
			variableWidthHeight, variableWidthHeight),
		tealayout.BorderRegionConfigOpt(tealayout.EastBorderRegion, textModel{text: "East component"}, fixedWidth, variableWidthHeight),
		tealayout.BorderRegionConfigOpt(tealayout.SouthBorderRegion, textModel{text: "South component"}, variableWidthHeight, fixedHeight),
	)

	model := &showcaseModel{mainModel: bl}

	s := lipgloss.NewStyle().Padding(0).Margin(0).
		BorderForeground(lipgloss.Color("36")).
		Border(lipgloss.NormalBorder(), true)

	for _, c := range model.mainModel.Components() {
		c.SetStyle(s)
	}

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}

type showcaseModel struct {
	mainModel *tealayout.BorderLayout
}

var _ tea.Model = &showcaseModel{}

// Init implements tea.Model.
func (s *showcaseModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (s *showcaseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//rezie etc...
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return s, tea.Quit
		case "n":
			return s, func() tea.Msg {
				region := tealayout.NorthBorderRegion
				s.mainModel.ToggleRegions(region)
				return tea.ClearScreen()
			}
		case "w":
			return s, func() tea.Msg {
				region := tealayout.WestBorderRegion
				s.mainModel.ToggleRegions(region)
				return tea.ClearScreen()
			}
		case "c":
			return s, func() tea.Msg {
				region := tealayout.CenterBorderRegion
				s.mainModel.ToggleRegions(region)
				return tea.ClearScreen()
			}
		case "e":
			return s, func() tea.Msg {
				region := tealayout.EastBorderRegion
				s.mainModel.ToggleRegions(region)
				return tea.ClearScreen()
			}
		case "s":
			return s, func() tea.Msg {
				region := tealayout.SouthBorderRegion
				s.mainModel.ToggleRegions(region)
				return tea.ClearScreen()
			}
		}
	case tea.WindowSizeMsg:
		s.mainModel.Resize(tealayout.Size{H: msg.Height, W: msg.Width})
		return s, nil
	}

	return s, nil
}

// View implements tea.Model.
func (s *showcaseModel) View() string {
	return s.mainModel.View()
}
