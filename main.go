package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	h, b, err := ImportData()
	if err != nil {
		return
	}
	// headers, rows := colAndrow()

	t := NewTable(h, b)
	if _, err := tea.NewProgram(t, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program", err)
		os.Exit(1)
	}
}

func colAndrow() ([]Col, [][]string) {

	h := []Col{
		{Name: "Header"},
		{Name: "Header1"},
		{Name: "Header2"},
		{Name: "Header3"},
	}

	r := [][]string{
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testsdasdasdrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
		{"testrow1", "testrow2", "testrow3", "testrow4"},
	}
	return h, r
}
