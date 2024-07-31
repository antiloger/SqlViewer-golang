package main

import "github.com/charmbracelet/lipgloss"

type TableStyle struct {
	Header      lipgloss.Style
	SelectedRow lipgloss.Style
	DefaultRow  lipgloss.Style
}

func DefaultStyle() TableStyle {
	return TableStyle{
		Header:      lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212")).Padding(0, 1).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240")).BorderBottom(true),
		SelectedRow: lipgloss.NewStyle().Bold(true).Padding(0, 1).Foreground(lipgloss.Color("229")).Background(lipgloss.Color("57")),
		DefaultRow:  lipgloss.NewStyle().Padding(0, 1),
	}
}
