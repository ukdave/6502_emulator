// Package tui implements the terminal user interface for the 6502 emulator.
package tui

import (
	"github.com/ukdave/6502_emulator/processor"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type runUpdateMsg struct{}

type Model struct {
	cpu            *processor.CPU
	previousMemory [65536]byte // Track previous memory state to detect changes

	runDelayMillis int
	running        bool
	runUpdateChan  chan runUpdateMsg

	width  int
	height int
	keys   keyMap
	help   help.Model

	boxStyle                lipgloss.Style
	statusBitSetStyle       lipgloss.Style
	statusBitClearStyle     lipgloss.Style
	runningStyle            lipgloss.Style
	currentInstructionStyle lipgloss.Style
	memoryChangedStyle      lipgloss.Style
	helpStyle               lipgloss.Style
}

func NewModel(cpu *processor.CPU, runDelayMillis int) *Model {
	m := &Model{
		cpu:                     cpu,
		runDelayMillis:          runDelayMillis,
		runUpdateChan:           make(chan runUpdateMsg),
		keys:                    keys,
		help:                    help.New(),
		boxStyle:                lipgloss.NewStyle().Padding(0, 1).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("63")),
		statusBitSetStyle:       lipgloss.NewStyle().Foreground(lipgloss.Color("2")),
		statusBitClearStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color("161")),
		runningStyle:            lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
		currentInstructionStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("111")),
		memoryChangedStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("11")),
		helpStyle:               lipgloss.NewStyle().PaddingTop(1),
	}
	m.updateMemoryTracking()
	return m
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.waitForRunUpdateMsg(),
	)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.updateDimensions(msg.Width, msg.Height)
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Step):
			m.step()
		case key.Matches(msg, m.keys.Run):
			return m, m.run()
		case key.Matches(msg, m.keys.IRQ):
			m.cpu.IRQ()
		case key.Matches(msg, m.keys.NMI):
			m.cpu.NMI()
		case key.Matches(msg, m.keys.Reset):
			m.cpu.Reset()
			m.updateMemoryTracking()
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	case runUpdateMsg:
		return m, m.waitForRunUpdateMsg()
	}
	return m, nil
}

func (m *Model) View() tea.View {
	rightColWidth := 40 + m.boxStyle.GetHorizontalFrameSize()
	statusPanelHeight := 12 + m.boxStyle.GetVerticalFrameSize()

	help := m.helpStyle.
		Render(m.help.View(m.keys))

	status := m.boxStyle.
		Width(rightColWidth).
		Height(statusPanelHeight).
		Render(m.statusView())

	memory := m.boxStyle.
		Width(m.width - rightColWidth).
		Height(m.height - lipgloss.Height(help)).
		Render(m.memoryView())

	instructionPanelHeight := lipgloss.Height(memory) - lipgloss.Height(status)
	instructions := m.boxStyle.
		Width(rightColWidth).
		Height(instructionPanelHeight).
		Render(m.instructionsView(instructionPanelHeight - m.boxStyle.GetVerticalFrameSize()))

	rightCol := lipgloss.JoinVertical(lipgloss.Left, status, instructions)
	topRow := lipgloss.JoinHorizontal(lipgloss.Top, memory, rightCol)
	viewStr := lipgloss.JoinVertical(lipgloss.Left, topRow, help)

	v := tea.NewView(viewStr)
	v.AltScreen = true
	return v
}
