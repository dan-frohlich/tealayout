package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type textModel struct {
	text string
}

// Init implements tea.Model.
func (t textModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (t textModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

// View implements tea.Model.
func (t textModel) View() string {
	return t.text
}

var _ tea.Model = textModel{text: "test"}
