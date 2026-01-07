// Package processor implements a 6502-compatible CPU core used to emulate the behaviour of the original 8-bit
// microprocessor. It defines the CPU structure, registers, status flags, and instruction logic. The CPU interacts
// with the rest of the system through the bus package, which provides access to memory and peripheral devices.
//
// This package is responsible for managing the processor state, including registers such as the accumulator, X and Y
// registers, stack pointer, and program counter, as well as manipulating the individual status flags that reflect
// processor conditions during execution.
package processor

import (
	"github.com/ukdave/6502_emulator/bus"
)

type Flag byte

// The status register stores 8 flags, which are enumerated here.
// The bits have different interpretations depending upon the context and instruction being executed.
const (
	C Flag = (1 << 0) // Carry Bit
	Z Flag = (1 << 1) // Zero
	I Flag = (1 << 2) // Disable Interrupts
	D Flag = (1 << 3) // Decimal Mode (unused in this implementation)
	B Flag = (1 << 4) // Break
	U Flag = (1 << 5) // Unused
	V Flag = (1 << 6) // Overflow
	N Flag = (1 << 7) // Negative
)

type CPU struct {
	bus bus.Bus

	// CPU Core registers, exported for ease of access by external inspectors. This is all the 6502 has.
	A      byte   // Accumulator Register
	X      byte   // X Register
	Y      byte   // Y Register
	SP     byte   // Stack Pointer (points to location on the bus within page 1)
	PC     uint16 // Program Counter
	Status byte   // Status Register

	cycles uint8
}

// NewCPU creates a new CPU instance.
func NewCPU(bus bus.Bus) *CPU {
	c := &CPU{bus: bus}
	c.Reset()
	return c
}

// Reset resets the CPU to its initial powerup state.
func (c *CPU) Reset() {
	c.A = 0x00
	c.X = 0x00
	c.Y = 0x00
	c.SP = 0xFD
	c.PC = c.ResetVector()
	c.Status = 0x24 // Clear all flags except U and I
}

// ResetVector returns the 16-bit address read from the 6502 reset vector ($FFFC–$FFFD), which is loaded into
// the program counter on reset.
func (c *CPU) ResetVector() uint16 {
	return c.Read16(0xFFFC)
}

func (c *CPU) Clock() {
	// TODO
}

// Read reads an 8-bit value from the bus at the specified address.
func (c *CPU) Read(addr uint16) byte {
	return c.bus.Read(addr)
}

// Read16 reads a 16-bit value from the bus at the specified address.
// The value is assumed to be stored least significant byte first (little endian).
func (c *CPU) Read16(addr uint16) uint16 {
	lo := uint16(c.Read(addr))
	hi := uint16(c.Read(addr + 1))
	return (hi << 8) | lo
}

// Write writes an 8-bit value to the bus at ths specified address.
func (c *CPU) Write(addr uint16, data byte) {
	c.bus.Write(addr, data)
}

// GetFlag returns the value of a specific bit of the status register.
func (c *CPU) GetFlag(flag Flag) bool {
	return (c.Status & byte(flag)) > 0
}

// SetFlag sets or clears a specific bit of the status register.
func (c *CPU) SetFlag(flag Flag, value bool) {
	if value {
		c.Status |= byte(flag)
	} else {
		c.Status &= ^byte(flag)
	}
}

// SetZN sets both the Zero and Negative flags based on the given value.
func (c *CPU) SetZN(value byte) {
	c.SetFlag(Z, value == 0x00)
	c.SetFlag(N, value&0x80 > 0)
}

func (c *CPU) addBranchCycles(addressInfo AddressInfo) {
	c.cycles++
	if addressInfo.PageChanged {
		c.cycles++
	}
}

// Cycles returns the number of remaining cycles (or clock ticks) required to complete the current instruction.
func (c *CPU) Cycles() uint8 {
	return c.cycles
}
