package processor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ukdave/6502_emulator/bus"
	"github.com/ukdave/6502_emulator/processor"
)

func TestNewCPU(t *testing.T) {
	// Write starting value (0x1234) for PC to 0xFFFC
	bus := bus.NewSimpleBus()
	bus.Write(0xFFFC, 0x34)
	bus.Write(0xFFFD, 0x12)

	// Create a new CPU
	cpu := processor.NewCPU(bus)

	assert.Equal(t, uint8(0), cpu.A, "Accumulator should be 0")
	assert.Equal(t, uint8(0), cpu.X, "X Register should be 0")
	assert.Equal(t, uint8(0), cpu.Y, "Y Register should be 0")
	assert.Equal(t, uint8(0xFD), cpu.SP, "Stack Pointer should be 0xFD")
	assert.Equal(t, uint16(0x1234), cpu.PC, "Program Counter should be 0x1234")
	assert.Equal(t, uint8(0b00100100), cpu.Status, "Status Flags should be 0b00100100")
}

func TestResetVector(t *testing.T) {
	// Write starting value (0x1234) for PC to 0xFFFC
	bus := bus.NewSimpleBus()
	bus.Write(0xFFFC, 0x34)
	bus.Write(0xFFFD, 0x12)

	// Create a new CPU
	cpu := processor.NewCPU(bus)

	assert.Equal(t, uint16(0x1234), cpu.ResetVector(), "Reset Vector should be 0x1234")
}

func TestRead(t *testing.T) {
	bus := bus.NewSimpleBus()
	bus.Write(0x1234, 0xAB)

	cpu := processor.NewCPU(bus)

	assert.Equal(t, uint8(0xAB), cpu.Read(0x1234))
}

func TestRead16(t *testing.T) {
	// Create a new bus and write the value 0xABCD at address 0x1235.
	// The 6502 is little endian so we write the least significant byte first.
	bus := bus.NewSimpleBus()
	bus.Write(0x1234, 0xCD)
	bus.Write(0x1235, 0xAB)

	// Create a new CPU
	cpu := processor.NewCPU(bus)

	assert.Equal(t, uint16(0xABCD), cpu.Read16(0x1234))
}

func TestGetFlag(t *testing.T) {
	// Create a new CPU
	cpu := &processor.CPU{}

	// Get the Zero flag (2nd bit) when set
	cpu.Status = 0b00000010
	assert.True(t, cpu.GetFlag(processor.Z), "Zero flag should be set")

	// Get the Overflow flag (7th bit) when set
	cpu.Status = 0b01000000
	assert.True(t, cpu.GetFlag(processor.V), "Overflow flag should be set")

	// Get both the Zero and Overflow flags (2nd and 7th bits) when set
	cpu.Status = 0b01000010
	assert.True(t, cpu.GetFlag(processor.Z), "Zero flag should be set")
	assert.True(t, cpu.GetFlag(processor.V), "Overflow flag should be set")

	// Get the Zero flag (2nd bit) when clear
	cpu.Status = 0b11111101
	assert.False(t, cpu.GetFlag(processor.Z), "Zero flag should be clear")

	// Get the Overflow flag (7th bit) when clear
	cpu.Status = 0b10111111
	assert.False(t, cpu.GetFlag(processor.V), "Overflow flag should be set")

	// Get both the Zero and Overflow flags (2nd and 7th bits) when clear
	cpu.Status = 0b10111101
	assert.False(t, cpu.GetFlag(processor.Z), "Zero flag should be clear")
	assert.False(t, cpu.GetFlag(processor.V), "Overflow flag should be clear")
}

func TestSetFlag(t *testing.T) {
	// Create a new CPU
	cpu := &processor.CPU{}

	// Set the Zero flag (2nd bit)
	cpu.Status = 0x00
	cpu.SetFlag(processor.Z, true)
	assert.True(t, cpu.GetFlag(processor.Z), "Zero flag should be set")

	// Set the Overflow flag (7th bit)
	cpu.Status = 0x00
	cpu.SetFlag(processor.V, true)
	assert.True(t, cpu.GetFlag(processor.V), "Overflow flag should be set")

	// Set both the Zero and Overflow flags (2nd and 7th bits)
	cpu.Status = 0x00
	cpu.SetFlag(processor.Z, true)
	cpu.SetFlag(processor.V, true)
	assert.Equal(t, uint8(0b01000010), cpu.Status, "Zero and Overflow flags should be set")

	// Clear the Zero flag (2nd bit)
	cpu.Status = 0xFF
	cpu.SetFlag(processor.Z, false)
	assert.Equal(t, uint8(0b11111101), cpu.Status, "Zero flag should be clear")

	// Clear the Overflow flag (7th bit)
	cpu.Status = 0xFF
	cpu.SetFlag(processor.V, false)
	assert.Equal(t, uint8(0b10111111), cpu.Status, "Overflow flag should be clear")

	// Clear both the Zero and Overflow flags (2nd and 7th bits)
	cpu.Status = 0xFF
	cpu.SetFlag(processor.Z, false)
	cpu.SetFlag(processor.V, false)
	assert.Equal(t, uint8(0b10111101), cpu.Status, "Zero and Overflow flags should be clear")
}
