package main

import (
	"fmt"
	"os"

	"github.com/ukdave/6502_emulator/bus"
	"github.com/ukdave/6502_emulator/processor"
	"github.com/ukdave/6502_emulator/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() *tui.Model {
	// Create a new bus
	bus := bus.NewSimpleBus()

	// Set the value of the reset vector to 0x8000. This is where our program will start
	bus.Write(0xFFFC, 0x00)
	bus.Write(0xFFFD, 0x80)

	// Write a very simple program to memory starting at 0x8000.
	// This program will multiply 10 (0x0A) by 3 (0x03) using repeated addition.
	bytes := []byte{
		0xA2, 0x0A, 0x8E, 0x00, 0x00, 0xA2, 0x03, 0x8E, 0x01, 0x00, 0xAC, 0x00, 0x00, 0xA9, 0x00, 0x18,
		0x6D, 0x01, 0x00, 0x88, 0xD0, 0xFA, 0x8D, 0x02, 0x00, 0xEA, 0xEA, 0xEA}
	for i, b := range bytes {
		bus.Write(0x8000+uint16(i), b)
	}

	// Create a new CPU
	cpu := processor.NewCPU(bus)
	return tui.NewModel(cpu)
}
