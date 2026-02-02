package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m *Model) memoryView() string {
	return m.renderMemoryPage(uint16(0x0000)) + "\n\n" + m.renderMemoryPage(m.cpu.ResetVector())
}

func (m *Model) renderMemoryPage(startAddress uint16) string {
	pageStr := ""
	startAddress &= 0xFF00 // Align to page boundary
	disassembledOp := m.cpu.DisassembleOperation(m.cpu.PC)
	for i := startAddress; i <= startAddress+0xF0; i += 16 {
		pageStr += fmt.Sprintf("$%04X: ", i)
		for j := uint16(0); j < 16; j++ {
			addr := i + j
			currentValue := m.cpu.Read(addr)
			hexStr := fmt.Sprintf("%02X", currentValue)

			if m.previousMemory[addr] != currentValue {
				hexStr = m.memoryChangedStyle.Render(hexStr)
			} else if addr >= m.cpu.PC && addr < m.cpu.PC+uint16(disassembledOp.Operation.Size) {
				hexStr = m.currentInstructionStyle.Render(hexStr)
			}

			pageStr += hexStr
			if j < 15 {
				pageStr += " "
			}
		}
		if i < startAddress+0xF0 {
			pageStr += "\n"
		}
	}
	return pageStr
}

func (m *Model) statusView() string {
	running := ""
	if m.running {
		running = m.runningStyle.Render("*** RUNNING ***")
	}
	return m.statusFlags() +
		fmt.Sprintf("PC:  $%04X       Cycles:  %d\n", m.cpu.PC, m.cpu.TotalCycles) +
		fmt.Sprintf("A:   $%02X  %-5s  %s\n", m.cpu.A, fmt.Sprintf("[%d]", m.cpu.A), running) +
		fmt.Sprintf("X:   $%02X  [%d]\n", m.cpu.X, m.cpu.X) +
		fmt.Sprintf("Y:   $%02X  [%d]\n", m.cpu.Y, m.cpu.Y) +
		fmt.Sprintf("SP:  $%02X\n\n", m.cpu.SP) +
		fmt.Sprintf("Reset Vector:  $%04X\n", m.cpu.ResetVector()) +
		fmt.Sprintf("NMI Vector:    $%04X\n", m.cpu.NMIVector()) +
		fmt.Sprintf("IRQ Vector:    $%04X", m.cpu.IRQVector())
}

func (m *Model) statusFlags() string {
	flags := fmt.Sprintf("%08b", m.cpu.Status)
	flagNames := []string{"N", "V", "-", "B", "D", "I", "Z", "C"}
	flagValues := make([]string, 8)

	for i, bit := range flags {
		if flagNames[i] == "-" {
			flagValues[i] = "-"
			continue
		}
		var style lipgloss.Style
		if bit == '1' {
			style = m.statusBitSetStyle
		} else {
			style = m.statusBitClearStyle
		}
		flagValues[i] = style.Render(flagNames[i])
	}

	return fmt.Sprintf("Status:  %s\n         %s  $%02X\n\n",
		strings.Join(flagValues, " "),
		strings.Join(strings.Split(flags, ""), " "),
		m.cpu.Status)
}

// instructionsView renders a disassembly of the instructions near the Program Counter. We start
// disassembly from the Reset Vector or current PC address, whichever is lowest. We need to disassemble
// at least `viewHeight` worth of instructions, but will continue past the current PC for a bit to
// ensure the current instruction appears somewhere in the top half of the view.
func (m *Model) instructionsView(viewHeight int) string {
	addr := min(uint16(m.cpu.ResetVector()), m.cpu.PC)

	lines := []string{}
	numInstructionsAfterPC := 0
	for {
		disassembledOp := m.cpu.DisassembleOperation(addr)

		line := fmt.Sprintf("$%04X: % -9X %s", addr, disassembledOp.Bytes, disassembledOp.Disassembly)
		if addr == m.cpu.PC {
			line = m.currentInstructionStyle.Render("> " + line)
		} else {
			line = ("  " + line)
		}
		lines = append(lines, line)

		addr += uint16(disassembledOp.Operation.Size)

		if addr > m.cpu.PC {
			numInstructionsAfterPC++
		}
		if len(lines) >= viewHeight && numInstructionsAfterPC > viewHeight/2 {
			break
		}
	}
	return strings.Join(lines[len(lines)-viewHeight:], "\n")
}
