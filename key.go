package main

import "github.com/charmbracelet/bubbles/key"

type TableKeyMap struct {
	LineUp   key.Binding
	LineDown key.Binding
}

func DefaultKey() TableKeyMap {
	return TableKeyMap{
		LineUp: key.NewBinding(
			key.WithKeys("up", "k"),
		),
		LineDown: key.NewBinding(
			key.WithKeys("down", "j"),
		),
	}
}
