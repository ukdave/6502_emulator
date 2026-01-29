package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ukdave/6502_emulator/bus"
	"github.com/ukdave/6502_emulator/processor"
	"github.com/ukdave/6502_emulator/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	binaryPath := parseArgs()

	// Create and start the TUI program
	p := tea.NewProgram(initialModel(binaryPath), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func parseArgs() string {
	// Optional help flags
	help := flag.Bool("help", false, "Show help")
	flag.BoolVar(help, "h", false, "Show help (shorthand)")

	// Parse flags
	flag.Parse()

	// Handle help
	if *help {
		fmt.Printf("Usage: %s [options] [binary-file]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Positional argument (binary file) is optional
	binaryPath := ""
	args := flag.Args()
	if len(args) > 0 {
		binaryPath = args[0]
	}

	return binaryPath
}

func initialModel(binaryPath string) *tui.Model {
	// Create a new bus
	bus := bus.NewSimpleBus()

	// Set the value of the reset vector to 0x8000. This is where our program will start
	bus.Write(0xFFFC, 0x00)
	bus.Write(0xFFFD, 0x80)

	// If a binary path was provided, load that file into memory at 0x8000
	if binaryPath != "" {
		binFile, err := os.ReadFile(binaryPath)
		if err != nil {
			fmt.Printf("Failed to read binary file: %v\n", err)
			os.Exit(1)
		}
		for i, b := range binFile {
			bus.Write(0x8000+uint16(i), b)
		}
	}

	// Create a new CPU
	cpu := processor.NewCPU(bus)
	return tui.NewModel(cpu)
}
