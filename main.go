package main

import (
	"fmt"

	"github.com/ukdave/6502_emulator/bus"
	"github.com/ukdave/6502_emulator/processor"
)

func main() {
	// Create a new bus
	bus := bus.NewSimpleBus()

	// Example usage of write and read
	bus.Write(0x1000, 0x42)
	data := bus.Read(0x1000)

	// Print out the data
	fmt.Printf("Data at address 0x1000: 0x%X\n", data)

	cpu := processor.NewCPU(bus)
	cpu.Reset()
}
