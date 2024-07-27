package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

type Table struct {
	Column     []Col
	ViewPoint  viewport.Model
	Curser     Grid
	Height     int
	Width      int
	Style      TableStyle
	MaxViewRow int
	RowBuff    PageBuff
}

func NewTable(col []Col, row [][]string) Table {
	return Table{
		Column: col,
		Curser: Grid{
			CurrX:  0,
			CurrY:  0,
			StartX: 0,
			StartY: 0,
		},
	}
}

func (t Table) Init() tea.Cmd { return nil }

// func (t Table) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// }

func (t Table) View() string {
	return (t.HeadersView() + "\n" + "")
}

func (t *Table) HeadersView() string {
	s := make([]string, 0, len(t.Column))
	for _, col := range t.Column {
		style := lipgloss.NewStyle().Width(col.MxWidth).MaxWidth(col.MxWidth).Inline(true)
		renderCell := style.Render(runewidth.Truncate(col.Name, col.MxWidth, "..."))
		s = append(s, t.Style.Header.Render(renderCell))
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, s...)
}

func (t *Table) SingleRowView(idx int) string {
	r := make([]string, 0, len(t.Column))
	for i, row := range t.Row[idx] {
		style := lipgloss.NewStyle().Width(t.Column[i].MxWidth).MaxWidth(t.Column[i].MxWidth).Inline(true)
		renderCell := t.Style.DefaultRow.Render(style.Render(runewidth.Truncate(row, t.Column[i].MxWidth, "...")))
		r = append(r, t.Style.DefaultRow.Render(renderCell))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Left, r...)
	if t.Curser.CurrY == idx {
		return t.Style.SelectedRow.Render(row)
	}

	return row
}

func (t *Table) MoveUp(line int) {
}

// func (t *Table) UpdateViewport() {
// 	if t.Curser.CurrY < t.Curser.StartY {
// 		if t.Curser.CurrY <= 0 {
// 			t.Curser.CurrY = 0
// 		} else {
//
// 		}
// 	} else if t.Curser.CurrY > t.Curser.EndY {
// 	} else {
// 	}
// }

type Col struct {
	Name    string
	MxWidth int
}

type Grid struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
	CurrX  int
	CurrY  int
}

type TableStyle struct {
	Header      lipgloss.Style
	SelectedRow lipgloss.Style
	DefaultRow  lipgloss.Style
}
