package tealayout

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	_ tea.Model     = &BorderLayout{}
	_ Resizable     = &BorderLayout{}
	_ LayoutManager = &BorderLayout{}
)

type BorderLayoutRegionID int

const (
	NorthBorderRegion BorderLayoutRegionID = iota
	WestBorderRegion
	CenterBorderRegion
	EastBorderRegion
	SouthBorderRegion
)

// BorderLayout a border layout

type BorderLayout struct {
	North  *LayoutComponent
	West   *LayoutComponent
	Center *LayoutComponent
	East   *LayoutComponent
	South  *LayoutComponent
	size   Size
}

func NewBorderLayout(opts ...BorderLayoutOpt) *BorderLayout {
	b := &BorderLayout{}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

type BorderLayoutOpt func(*BorderLayout)

func BorderRegionConfigOpt(region BorderLayoutRegionID, c tea.Model, fixedWidth int, fixedHeight int) BorderLayoutOpt {
	switch region {
	case NorthBorderRegion:
		return northComponentOpt(c, fixedHeight)
	case WestBorderRegion:
		return westComponentOpt(c, fixedWidth)
	case CenterBorderRegion:
		return centerComponentOpt(c)
	case EastBorderRegion:
		return eastComponentOpt(c, fixedWidth)
	case SouthBorderRegion:
		return southComponentOpt(c, fixedHeight)
	}
	return func(b *BorderLayout) {}
}

func northComponentOpt(c tea.Model, height int) BorderLayoutOpt {
	return func(b *BorderLayout) {
		if b.North == nil {
			b.North = NewLayoutComponent(c)
		} else {
			b.North.wrapped = c
		}
		b.North.visible = true
		if height > 0 {
			b.North.fixedHeight = height
		}
	}
}

func westComponentOpt(c tea.Model, width int) BorderLayoutOpt {
	return func(b *BorderLayout) {
		if b.West == nil {
			b.West = NewLayoutComponent(c)
		} else {
			b.West.wrapped = c
		}
		b.West.visible = true
		if width > 0 {
			b.West.fixedWidth = width
		}
	}
}

func centerComponentOpt(c tea.Model) BorderLayoutOpt {
	return func(b *BorderLayout) {
		if b.Center == nil {
			b.Center = NewLayoutComponent(c)
		} else {
			b.Center.wrapped = c
		}
		b.Center.visible = true
	}
}

func eastComponentOpt(c tea.Model, width int) BorderLayoutOpt {
	return func(b *BorderLayout) {
		if b.East == nil {
			b.East = NewLayoutComponent(c)
		} else {
			b.East.wrapped = c
		}
		b.East.visible = true
		if width > 0 {
			b.East.fixedWidth = width
		}
	}
}

func southComponentOpt(c tea.Model, height int) BorderLayoutOpt {
	return func(b *BorderLayout) {
		if b.South == nil {
			b.South = NewLayoutComponent(c)
		} else {
			b.South.wrapped = c
		}
		b.South.visible = true
		if height > 0 {
			b.South.fixedHeight = height
		}
	}
}

func ComponentStyling(s lipgloss.Style) BorderLayoutOpt {
	return func(b *BorderLayout) {
		if b.North != nil {
			b.North.viewport.Style = s
		}
		if b.West != nil {
			b.West.viewport.Style = s
		}
		if b.Center != nil {
			b.Center.viewport.Style = s
		}
		if b.East != nil {
			b.East.viewport.Style = s
		}
		if b.South != nil {
			b.South.viewport.Style = s
		}
	}
}

// Components implements LayoutManager.
func (b *BorderLayout) Components() []*LayoutComponent {
	return []*LayoutComponent{b.North, b.West, b.Center, b.East, b.South}
}

func (b *BorderLayout) layout() {
	b.Resize(b.size)
}

// Resize implements Resizable.
func (b *BorderLayout) Resize(sz Size) {
	b.size = sz
	centerHeight := sz.H
	centerWidth := sz.W
	if b.North != nil && b.North.visible {
		c := b.North
		centerHeight -= c.fixedHeight
		c.Resize(Size{H: c.fixedHeight, W: sz.W})
	}
	if b.South != nil && b.South.visible {
		c := b.South
		centerHeight -= c.fixedHeight
		c.Resize(Size{H: c.fixedHeight, W: sz.W})
	}
	if b.West != nil && b.West.visible {
		c := b.West
		centerWidth -= c.fixedWidth
		c.Resize(Size{H: centerHeight, W: c.fixedWidth})
	}
	if b.East != nil && b.East.visible {
		c := b.East
		centerWidth -= c.fixedWidth
		c.Resize(Size{H: centerHeight, W: c.fixedWidth})
	}
	if b.Center != nil && b.Center.visible {
		b.Center.Resize(Size{H: centerHeight, W: centerWidth})
	}
}

// Init implements tea.Model.
func (b *BorderLayout) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (b *BorderLayout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return b, nil
}

// View implements tea.Model.
func (b *BorderLayout) View() string {
	var horzStr []string
	if b.West.Visible() {
		horzStr = append(horzStr, b.West.View())
	}
	if b.Center.Visible() {
		horzStr = append(horzStr, b.Center.View())
	}
	if b.East.Visible() {
		horzStr = append(horzStr, b.East.View())
	}
	var vertStr []string

	if b.North.Visible() {
		vertStr = append(vertStr, b.North.View())
	}
	if len(horzStr) > 0 {
		vertStr = append(vertStr, lipgloss.JoinHorizontal(lipgloss.Top, horzStr...))
	}
	if b.South.Visible() {
		vertStr = append(vertStr, b.South.View())
	}
	return lipgloss.JoinVertical(lipgloss.Left, vertStr...)
}

func (b *BorderLayout) ShowRegions(ids ...BorderLayoutRegionID) {
	b.setRegionVisibilitys(true, ids...)
}

func (b *BorderLayout) HideRegions(ids ...BorderLayoutRegionID) {
	b.setRegionVisibilitys(false, ids...)
}

func (b *BorderLayout) ToggleRegions(ids ...BorderLayoutRegionID) {
	for _, id := range ids {
		switch id {
		case NorthBorderRegion:
			if b.North == nil {
				continue
			}
			b.North.visible = !b.North.visible
		case WestBorderRegion:
			if b.West == nil {
				continue
			}
			b.West.visible = !b.West.visible
		case CenterBorderRegion:
			if b.Center == nil {
				continue
			}
			b.Center.visible = !b.Center.visible
		case EastBorderRegion:
			if b.East == nil {
				continue
			}
			b.East.visible = !b.East.visible
		case SouthBorderRegion:
			if b.South == nil {
				continue
			}
			b.South.visible = !b.South.visible
		}
	}
	if len(ids) > 0 {
		b.layout()
	}
}

func (b *BorderLayout) Visibility(id BorderLayoutRegionID) bool {
	switch id {
	case NorthBorderRegion:
		return b.North.Visible()
	case WestBorderRegion:
		return b.West.Visible()
	case CenterBorderRegion:
		return b.Center.Visible()
	case EastBorderRegion:
		return b.East.Visible()
	case SouthBorderRegion:
		return b.South.Visible()
	}
	return false
}

func (b *BorderLayout) setRegionVisibilitys(show bool, ids ...BorderLayoutRegionID) {
	for _, id := range ids {
		switch id {
		case NorthBorderRegion:
			if b.North == nil {
				continue
			}
			b.North.visible = show
		case WestBorderRegion:
			if b.West == nil {
				continue
			}
			b.West.visible = show
		case CenterBorderRegion:
			if b.Center == nil {
				continue
			}
			b.Center.visible = show
		case EastBorderRegion:
			if b.East == nil {
				continue
			}
			b.East.visible = show
		case SouthBorderRegion:
			if b.South == nil {
				continue
			}
			b.South.visible = show
		}
	}
	if len(ids) > 0 {
		b.layout()
	}
}
