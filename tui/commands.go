package tui

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

func (m *Model) updateMemoryTracking() {
	for i := uint32(0); i < 65536; i++ {
		m.previousMemory[i] = m.cpu.Read(uint16(i))
	}
}
