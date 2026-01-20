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

// IRQVector returns the 16-bit address read from the 6502 IRQ/BRK vector ($FFFE–$FFFF), which is loaded into
// the program counter on an IRQ or BRK.
func (c *CPU) IRQVector() uint16 {
	return c.Read16(0xFFFE)
}

// NMIVector returns the 16-bit address read from the 6502 NMI vector ($FFFA–$FFFB), which is loaded into
// the program counter on a non-maskable interrupt.
func (c *CPU) NMIVector() uint16 {
	return c.Read16(0xFFFA)
}

// Clock advances the CPU by a single clock cycle.
//
// 6502 instructions take a variable number of cycles to complete. On real hardware, each cycle performs a small
// portion of the instruction via internal micro-operations. This emulator executes the full instruction atomically,
// but still models timing by tracking the number of cycles the instruction consumes. Each call to Clock decrements
// the remaining cycle count, and when it reaches zero the instruction is considered complete.
func (c *CPU) Clock() {
	if c.cycles > 0 {
		c.cycles--
		return
	}

	opcode := c.Read(c.PC)
	op := operations[opcode]

	// Get the address information/operand using the appropriate address mode for this operation.
	// Note that not all instructions require an operand (e.g. NOP, INX, CLC).
	addressInfo := op.AddressMode(c)

	// Increment the Program Counter (PC) by the size of this operation. We do this *before* executing the
	// instruction as some instructions may alter PC directly (e.g. branch instructions).
	c.PC += uint16(op.Size)

	// Get the starting number of cycles for this operation
	c.cycles = op.Cycles

	// Perform operation
	extraCycle := op.Instruction(c, addressInfo)

	// Several addressing modes have the potential to require an additional clock cycle if they cross a page
	// boundary. This is combined with several instructions that enable this additional clock cycle. If both
	// the instruction and address function return true, then an additional clock cycle is required.
	//
	// Branch instructions require an additional clock cycle if the branch is taken, and a second additional
	// clock cycle if the page boundary is crossed. Our branch instruction functions always return false and
	// handle the addition any extra cycles themselves.
	if extraCycle && addressInfo.PageChanged {
		c.cycles++
	}

	// Decrement the number of cycles remaining for this instruction
	c.cycles--
}

// IRQ performs an Interrupt Request (IRQ) sequence.
// The current Program Counter and status flags are pushed onto the stack and then we jump to the address
// stored in the IRQ vector.
func (c *CPU) IRQ() {
	// Only run if the Disable Interrupts flag is clear
	if !c.GetFlag(I) {
		c.Push16(c.PC)
		c.Push(c.Status &^ 0x10) // 0x10 clears the Break flag to 1 (but only in the value pushed to the stack)
		c.SetFlag(I, true)       // Set the "Interrupt Disable" flag
		c.PC = c.IRQVector()
		c.cycles += 7
	}
}

// NMI performs a Non-Maskable Interrupt (NMI) sequence.
// The current Program Counter and status flags are pushed onto the stack and then we jump to the address
// stored in the NMI vector.
func (c *CPU) NMI() {
	c.Push16(c.PC)
	c.Push(c.Status &^ 0x10) // 0x10 clears the Break flag to 1 (but only in the value pushed to the stack)
	c.SetFlag(I, true)       // Set the "Interrupt Disable" flag
	c.PC = c.NMIVector()
	c.cycles += 7
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

// Write16 writes a 16-bit value to the bus at the specified address.
// The value is written least significant byte first (little endian).
func (c *CPU) Write16(addr uint16, data uint16) {
	c.bus.Write(addr, uint8(data&0xFF))
	c.bus.Write(addr+1, uint8(data>>8))
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

// Push pushes an 8-bit value onto the stack.
func (c *CPU) Push(value uint8) {
	// Remember that the stack is stored in page 1 (so we need to add 0x100 to the value of the stack pointer).
	// Also, the stack pointer starts at 0xFD after a reset and grows down, so we need to decrement it after pushing.
	c.Write(0x100|uint16(c.SP), value)
	c.SP--
}

// Push16 pushes a 16-bit value onto the stack.
// The 6502 is little endian (LSB first) but the stack grows down so we push the most significant byte first.
func (c *CPU) Push16(value uint16) {
	hi := uint8(value >> 8)
	lo := uint8(value & 0xFF)
	c.Push(hi)
	c.Push(lo)
}

// Pop pops an 8-bit value off the stack.
func (c *CPU) Pop() uint8 {
	c.SP++
	return c.Read(0x100 | uint16(c.SP))
}

// Pop16 pops a 16-bit value off the stack.
// The value is assumed to be stored least significant byte first (little endian).
func (c *CPU) Pop16() uint16 {
	lo := uint16(c.Pop())
	hi := uint16(c.Pop())
	return hi<<8 | lo
}
