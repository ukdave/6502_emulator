package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) updateDimensions(width, height int) {
	m.width = width
	m.height = height
	m.help.Width = width
}

func (m *Model) step() {
	m.updateMemoryTracking()
	for {
		m.cpu.Clock()
		if m.cpu.Cycles() == 0 {
			break
		}
	}
}

func (m *Model) run() tea.Cmd {
	return func() tea.Msg {
		if m.running {
			m.running = false
		} else {
			m.running = true
			for {
				pcBefore := m.cpu.PC
				m.step()
				m.runUpdateChan <- runUpdateMsg{}
				time.Sleep(time.Duration(m.runDelayMillis) * time.Millisecond)
				if !m.running || m.cpu.PC == 0x0000 || m.cpu.PC == pcBefore {
					break
				}
			}
			m.running = false
			m.runUpdateChan <- runUpdateMsg{}
		}
		return nil
	}
}

func (m *Model) waitForRunUpdateMsg() tea.Cmd {
	return func() tea.Msg {
		return <-m.runUpdateChan
	}
}

func (m *Model) updateMemoryTracking() {
	for i := uint32(0); i < 65536; i++ {
		m.previousMemory[i] = m.cpu.Read(uint16(i))
	}
}
