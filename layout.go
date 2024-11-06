package tealayout

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Size size of a component
type Size struct {
	W int
	H int
}

// Resize a component
type Resizable interface {
	// Resize a component
	Resize(sz Size)
}

func NewLayoutComponent(wrapped tea.Model) *LayoutComponent {
	return &LayoutComponent{
		wrapped:  wrapped,
		viewport: viewport.New(0, 0),
	}
}

var (
	_ tea.Model = &LayoutComponent{}
	_ Resizable = &LayoutComponent{}
)

type LayoutComponent struct {
	wrapped     tea.Model
	viewport    viewport.Model
	fixedHeight int
	fixedWidth  int
	visible     bool
}

// Resize implements Resizable.
func (lc *LayoutComponent) Resize(sz Size) {
	if lc == nil {
		return
	}
	lc.viewport.Height = sz.H
	lc.viewport.Width = sz.W
	lc.viewport.SetContent(fmt.Sprintf("%d x %d", sz.W, sz.H))
}

// Init implements tea.Model.
func (lc *LayoutComponent) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (lc *LayoutComponent) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return lc, nil
}

// View implements tea.Model.
func (lc *LayoutComponent) View() string {
	if lc == nil {
		return ""
	}
	x := lc.wrapped.View() + fmt.Sprintf("\n%d x %d", lc.viewport.Width, lc.viewport.Height)
	lc.viewport.SetContent(x)
	log.Println(x)
	return lc.viewport.View()
}

func (lc *LayoutComponent) SetStyle(s lipgloss.Style) {
	if lc == nil {
		return
	}
	lc.viewport.Style = s
}

func (lc *LayoutComponent) Visible() bool {
	return lc != nil && lc.visible
}

type LayoutManager interface {
	Components() []*LayoutComponent
}
