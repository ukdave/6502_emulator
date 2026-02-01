package tui

import "github.com/charmbracelet/bubbles/key"

// keyMap defines a set of keybindings. To work for help it must satisfy key.Map.
type keyMap struct {
	Step  key.Binding
	Run   key.Binding
	Reset key.Binding
	IRQ   key.Binding
	NMI   key.Binding
	Quit  key.Binding
}

var keys = keyMap{
	Step: key.NewBinding(
		key.WithKeys(" ", "enter"),
		key.WithHelp("space/enter", "Step"),
	),
	Run: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "Run/Stop"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "Reset"),
	),
	IRQ: key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", "IRQ"),
	),
	NMI: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "NMI"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "Quit"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Step, k.Run, k.Reset, k.IRQ, k.NMI, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
