package main

import (
	"fmt"
	"os"

	"github.com/ukdave/6502_emulator/bus"
	"github.com/ukdave/6502_emulator/processor"
	"github.com/ukdave/6502_emulator/tui"

	tea "github.com/charmbracelet/bubbletea"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	StartAddress uint16 `short:"s" long:"start" description:"Start address to load the binary file into memory" default:"0x8000"`

	Args struct {
		BinaryPath string `positional-arg-name:"binary_file" description:"Path to the binary file to load into memory"`
	} `positional-args:"yes"`
}

func main() {
	_, err := flags.Parse(&opts)
	if flags.WroteHelp(err) {
		os.Exit(0)
	}
	if err != nil {
		os.Exit(1)
	}

	// Create and start the TUI program
	p := tea.NewProgram(initialModel(opts.Args.BinaryPath, opts.StartAddress), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel(binaryPath string, startAddress uint16) *tui.Model {
	// Create a new bus
	bus := bus.NewSimpleBus()

	// Set the value of the reset vector to startAddress. This is where our program will start
	bus.Write(0xFFFC, uint8(startAddress&0xFF))
	bus.Write(0xFFFD, uint8((startAddress>>8)&0xFF))

	// If a binary path was provided, load that file into memory at startAddress
	if binaryPath != "" {
		binFile, err := os.ReadFile(binaryPath)
		if err != nil {
			fmt.Printf("Failed to read binary file: %v\n", err)
			os.Exit(1)
		}
		for i, b := range binFile {
			bus.Write(startAddress+uint16(i), b)
		}
	}

	// Create a new CPU
	cpu := processor.NewCPU(bus)
	return tui.NewModel(cpu)
}
