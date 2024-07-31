package main

import (
	// "github.com/charmbracelet/bubbles/key"

	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

type Table struct {
	MaxCol     int
	Column     []Col
	Curser     Grid
	Height     int
	Width      int
	Style      TableStyle
	MaxViewRow int
	RowBuff    PageBuff
	Keybind    TableKeyMap
	ViewBuff   []string
}

func NewTable(col []Col, row [][]string) *Table {
	return &Table{
		Column: col,
		Curser: Grid{
			CurrX:  0,
			CurrY:  0,
			StartX: 0,
			StartY: 0,
		},
		RowBuff: NewPageBuff(row, 2),
		Style:   DefaultStyle(),
		Keybind: DefaultKey(),
	}
}

func (t Table) Init() tea.Cmd { return nil }

func (t Table) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t.UpdateTableView(msg.Height, msg.Width)
		return t, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return t, tea.Quit
			// case key.Matches(msg, t.Keybind.LineUp):
			// 	t.MoveUp(1)
			// case key.Matches(msg, t.Keybind.LineDown):
			// 	t.MoveDown(1)
		}
	}

	return t, nil
}

func (t Table) View() string {
	t.UpdateViewBuff()
	return (t.HeadersView() + "\n" + t.RowViewOutput())
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

// func (t *Table) SingleRowView(idx int) string {
// 	r := make([]string, 0, len(t.Column))
// 	for i, row := range t.Row[idx] {
// 		style := lipgloss.NewStyle().Width(t.Column[i].MxWidth).MaxWidth(t.Column[i].MxWidth).Inline(true)
// 		renderCell := t.Style.DefaultRow.Render(style.Render(runewidth.Truncate(row, t.Column[i].MxWidth, "...")))
// 		r = append(r, t.Style.DefaultRow.Render(renderCell))
// 	}
//
// 	row := lipgloss.JoinHorizontal(lipgloss.Left, r...)
// 	if t.Curser.CurrY == idx {
// 		return t.Style.SelectedRow.Render(row)
// 	}
//
// 	return row
// }

func (t *Table) RowViewOutput() string {
	return lipgloss.JoinVertical(lipgloss.Left, t.ViewBuff...)
}

func (t *Table) UpdateViewBuff() {
	start, end, err := t.RowBuff.GetPageRealIndex(t.RowBuff.currPage)
	if err != nil {
		fmt.Println("Error updating view buffer:", err)
		return
	}

	t.ViewBuff = make([]string, 0, end-start+1)
	for i := start; i <= end && i < len(t.RowBuff.buff); i++ {
		r := make([]string, len(t.Column))
		for j, v := range t.RowBuff.buff[i] {
			if j >= len(t.Column) {
				break
			}
			style := lipgloss.NewStyle().Width(t.Column[j].MxWidth).MaxWidth(t.Column[j].MxWidth).Inline(true)
			renderCell := style.Render(runewidth.Truncate(v, t.Column[j].MxWidth, "..."))
			r[j] = t.Style.DefaultRow.Render(renderCell)
		}
		row := lipgloss.JoinHorizontal(lipgloss.Left, r...)
		if t.Curser.CurrY == i {
			row = t.Style.SelectedRow.Render(row)
		}
		t.ViewBuff = append(t.ViewBuff, row)
	}
}

func (t *Table) ModifyViewBuff() {
	// 1. setup current x, y value & check original buffer need to add somthing
	// 2. modify the buffer for current Changes

}

// Note: this need to called after changing the state
// func (t *Table) UpdateViewBuff() {
// 	s := make([]string, t.RowBuff.pageBuffSize)
// 	start, end, err := t.RowBuff.GetPageRealIndex(t.RowBuff.currPage)
// 	// TODO: error handle
// 	if err != nil {
// 		return
// 	}
//
// 	for i := start; i <= end; i++ {
// 		r := make([]string, len(t.RowBuff.buff[i]))
// 		for _, v := range t.RowBuff.buff[i] {
// 			style := lipgloss.NewStyle().Width(t.Column[i].MxWidth).MaxWidth(t.Column[i].MxWidth).Inline(true)
// 			renderCell := t.Style.DefaultRow.Render(style.Render(runewidth.Truncate(v, t.Column[i].MxWidth, "...")))
// 			r = append(r, t.Style.DefaultRow.Render(renderCell))
// 		}
// 		s = append(s, lipgloss.JoinHorizontal(lipgloss.Left, r...))
// 	}
//
// 	t.ViewBuff = s
// }

func (t *Table) UpdateTableView(height, width int) {
	if t.Height != height || t.Width != width {
		t.Height = height
		t.Width = width

		// Update row buffer size
		t.RowBuff.ChangePageBuffSize(height - 2)

		// Update column widths
		if len(t.Column) > 0 {
			w := width / len(t.Column)
			for i := range t.Column {
				t.Column[i].SetMxWidth(w)
			}
		}

		// Force view buffer update
		t.UpdateViewBuff()
	}
}

// func (t Table) UpdateTableView(height, width int) {
// 	t.Width = width
// 	t.Height = height
// 	if height != t.Height {
// 		// 4* means 2 for header and 2 for footer headerSize + footerSize + RowSize = Height
// 		t.RowBuff.ChangePageBuffSize(height - 4)
// 	}
// 	if width != t.Width {
// 		w := t.Width / len(t.Column)
// 		for i := range t.Column {
// 			t.Column[i].SetMxWidth(w)
// 		}
// 		//TODO: apply WidthChange func
// 	}
// }

func (t *Table) WidthChange(height int) {

}

func (t *Table) MoveUp(line int) {
}

func (t *Table) MoveDown(line int) {
}

type Col struct {
	Name    string
	MxWidth int
}

func (c *Col) SetMxWidth(width int) {
	c.MxWidth = width
}

type Grid struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
	CurrX  int
	CurrY  int
}
