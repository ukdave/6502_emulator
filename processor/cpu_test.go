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
	// Write value (0x1234) for PC to 0xFFFC
	bus := bus.NewSimpleBus()
	bus.Write(0xFFFC, 0x34)
	bus.Write(0xFFFD, 0x12)

	// Create a new CPU
	cpu := processor.NewCPU(bus)

	assert.Equal(t, uint16(0x1234), cpu.ResetVector(), "Reset Vector should be 0x1234")
}

func IRQVector(t *testing.T) {
	// Write value (0x1234) for PC to 0xFFFE
	bus := bus.NewSimpleBus()
	bus.Write(0xFFFE, 0x34)
	bus.Write(0xFFFF, 0x12)

	// Create a new CPU
	cpu := processor.NewCPU(bus)

	assert.Equal(t, uint16(0x1234), cpu.IRQVector(), "Reset Vector should be 0x1234")
}

func NMIVector(t *testing.T) {
	// Write value (0x1234) for PC to 0xFFFA
	bus := bus.NewSimpleBus()
	bus.Write(0xFFFA, 0x34)
	bus.Write(0xFFFB, 0x12)

	// Create a new CPU
	cpu := processor.NewCPU(bus)

	assert.Equal(t, uint16(0x1234), cpu.NMIVector(), "Reset Vector should be 0x1234")
}

func TestClock(t *testing.T) {
	bus := bus.NewSimpleBus()

	// Set the value of the reset vector to 0x8000. This is where our program will start
	bus.Write(0xFFFC, 0x00)
	bus.Write(0xFFFD, 0x80)

	// Write the instruction "LDA #$05" to memory starting at 0x8000.
	// A9 = opcode for LDA immediate
	// 05 = the literal value
	bus.Write(0x8000, 0xA9)
	bus.Write(0x8001, 0x05)

	// Create a new CPU and check initial state
	cpu := processor.NewCPU(bus)
	assert.Equal(t, uint8(0x00), cpu.A, "Expected the Accumulator to be 0x00")
	assert.Equal(t, uint16(0x8000), cpu.PC, "Expected the Program Counter to be 0x8000")

	// This instruction should take 2 clock cycles to complete
	cpu.Clock()
	cpu.Clock()

	// Check CPU state after executing the instruction
	assert.Equal(t, uint8(0x05), cpu.A, "Expected the Accumulator to be 0x05")
	assert.Equal(t, uint16(0x8002), cpu.PC, "Expected the Program Counter to be 0x8002")
	assert.Equal(t, uint8(0), cpu.Cycles(), "Expected Cycles to be 0")
}

func TestIRQ_Enabled(t *testing.T) {
	bus := bus.NewSimpleBus()
	cpu := processor.NewCPU(bus)

	// Set the Program Counter to a known value
	cpu.PC = 0x2000

	// Write a 16-bit value to memory to 0xFFFE
	cpu.Write16(0xFFFE, 0x1005)

	// Enable interrupts
	cpu.SetFlag(processor.I, false)

	// Read current status flags
	statusFlags := cpu.Status

	// Trigger IRQ
	cpu.IRQ()

	assert.Equal(t, uint16(0x1005), cpu.PC, "Expected PC to be 0x1005")
	assert.Equal(t, true, cpu.GetFlag(processor.I), "Disable Interrupt flag should be true")
	assert.Equal(t, uint8(0xFA), cpu.SP, "Expected stack pointer to be 0xFA (3 bytes pushed)")
	assert.Equal(t, statusFlags&^0x10, cpu.Pop(), "Expected stack to contain status flags with B clear")
	assert.Equal(t, uint16(0x2000), cpu.Pop16(), "Expected stack contain original PC value (0x2000)")
}

func TestIRQ_Disabled(t *testing.T) {
	bus := bus.NewSimpleBus()
	cpu := processor.NewCPU(bus)

	// Set the Program Counter to a known value
	cpu.PC = 0x2000

	// Write a 16-bit value to memory to 0xFFFE
	cpu.Write16(0xFFFE, 0x1005)

	// Disable interrupts
	cpu.SetFlag(processor.I, true)

	// Trigger IRQ
	cpu.IRQ()

	assert.Equal(t, uint16(0x2000), cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(t, uint8(0xFD), cpu.SP, "Expected stack pointer to be 0xFD (nothing pushed)")
}

func TestNMI(t *testing.T) {
	bus := bus.NewSimpleBus()
	cpu := processor.NewCPU(bus)

	// Set the Program Counter to a known value
	cpu.PC = 0x2000

	// Write a 16-bit value to memory to 0xFFFA
	cpu.Write16(0xFFFA, 0x1005)

	// Read current status flags
	statusFlags := cpu.Status

	// Trigger NMI
	cpu.NMI()

	assert.Equal(t, uint16(0x1005), cpu.PC, "Expected PC to be 0x1005")
	assert.Equal(t, true, cpu.GetFlag(processor.I), "Disable Interrupt flag should be true")
	assert.Equal(t, uint8(0xFA), cpu.SP, "Expected stack pointer to be 0xFA (3 bytes pushed)")
	assert.Equal(t, statusFlags&^0x10, cpu.Pop(), "Expected stack to contain status flags with B clear")
	assert.Equal(t, uint16(0x2000), cpu.Pop16(), "Expected stack contain original PC value (0x2000)")
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

func TestWrite(t *testing.T) {
	// Create a new CPU
	bus := bus.NewSimpleBus()
	cpu := processor.NewCPU(bus)

	// Write a value to the bus
	cpu.Write(0x1234, 0xAB)

	assert.Equal(t, uint8(0xAB), bus.Read(0x1234))
}

func TestWrite16(t *testing.T) {
	// Create a new CPU
	bus := bus.NewSimpleBus()
	cpu := processor.NewCPU(bus)

	// Write a value to the bus
	cpu.Write16(0x1234, 0xABCD)

	assert.Equal(t, uint8(0xCD), bus.Read(0x1234))
	assert.Equal(t, uint8(0xAB), bus.Read(0x1235))
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

func TestSetZN(t *testing.T) {
	// Create a new CPU
	cpu := &processor.CPU{}

	// Set the Zero flag and clear the Negative flag
	cpu.SetFlag(processor.Z, true)
	cpu.SetFlag(processor.N, false)

	// Update Zero and Clear flags together
	cpu.SetZN(0xFF)

	assert.False(t, cpu.GetFlag(processor.Z), "Zero flag should be clear")
	assert.True(t, cpu.GetFlag(processor.N), "Negative flag should be set")
}

func TestPush(t *testing.T) {
	// Create a new CPU
	bus := bus.NewSimpleBus()
	cpu := processor.NewCPU(bus)

	// Push a value on to the stack
	cpu.Push(0x12)

	// The stack starts at 0x01FD, so this is where our value should be stored
	assert.Equal(t, uint8(0x12), bus.Read(0x01FD))

	// The stack pointer should now be at 0xFC
	assert.Equal(t, uint8(0xFC), cpu.SP, "Expected the Stack Pointer to be 0xFC")
}

func TestPush16(t *testing.T) {
	// Create a new CPU
	bus := bus.NewSimpleBus()
	cpu := processor.NewCPU(bus)

	// Push a value on to the stack
	cpu.Push16(0x1234)

	// The stack starts at 0x01FD and grows down, so the start of our 16-bit value should
	// stored at 0x01FC with the least significant byte appearing first.
	assert.Equal(t, uint8(0x34), bus.Read(0x01FC))
	assert.Equal(t, uint8(0x12), bus.Read(0x01FD))

	// The stack pointer should now be at 0xFB
	assert.Equal(t, uint8(0xFB), cpu.SP, "Expected the Stack Pointer to be 0xFB")
}

func TestPop(t *testing.T) {
	// Create a new CPU
	bus := bus.NewSimpleBus()
	cpu := processor.NewCPU(bus)

	// Push a value on to the stack
	cpu.Push(0x12)

	assert.Equal(t, uint8(0x12), cpu.Pop(), "Expected to pop 0x12 off the stack")
	assert.Equal(t, uint8(0xFD), cpu.SP, "Expected the Stack Pointer to be 0xFD")
}

func TestPop16(t *testing.T) {
	// Create a new CPU
	bus := bus.NewSimpleBus()
	cpu := processor.NewCPU(bus)

	// Push a value on to the stack
	cpu.Push16(0x1234)

	assert.Equal(t, uint16(0x1234), cpu.Pop16(), "Expected to pop 0x1234 off the stack")
	assert.Equal(t, uint8(0xFD), cpu.SP, "Expected the Stack Pointer to be 0xFD")
}
