package processor

// The 6502 implements 56 instructions.
//
// Access:     LDA, STA, LDX, STX, LDY, STY
// Transfer:   TAX, TXA, TAY, TYA
// Arithmetic: ADC, SBC, INC, DEC, INX, DEX, INY, DEY
// Shift:      ASL, LSR, ROL, ROR
// Bitwise:    AND, ORA, EOR, BIT
// Compare:    CMP, CPX, CPY
// Branch:     BCC, BCS, BEQ, BNE, BPL, BMI, BVC, BVS
// Jump:       JMP, JSR, RTS, BRK, RTI
// Stack:      PHA, PLA, PHP, PLP, TXS, TSX
// Flags:      CLC, SEC, CLI, SEI, CLD, SED, CLV
// Other:      NOP
//
// https://www.nesdev.org/wiki/Instruction_reference
//
// Some addressing modes may incur an extra clock cycle when a memory access crosses a page boundary. Certain
// instructions allow this additional cycle to be taken. To model this behaviour, both the addressing mode and
// the instruction report whether an extra cycle is possible. If both indicate true, one additional clock cycle
// is added.

type InstructionFunc func(*CPU, AddressInfo) bool

//
// Access Instructions
//

// LDA - Load Accumulator
// Function:  A = memory
// Flags Out: Z, N
func LDA(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.A = cpu.Read(addressInfo.Address)
	cpu.SetZN(cpu.A)
	return true
}

// STA - Store Accumulator
// Function:  memory = A
// Flags Out: None
func STA(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Write(addressInfo.Address, cpu.A)
	return false
}

// LDX - Load X Register
// Function:  X = memory
// Flags Out: Z, N
func LDX(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.X = cpu.Read(addressInfo.Address)
	cpu.SetZN(cpu.X)
	return true
}

// STX - Store X Register
// Function:  memory = X
// Flags Out: None
func STX(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Write(addressInfo.Address, cpu.X)
	return false
}

// LDY - Load Y Register
// Function:  A = memory
// Flags Out: Z, N
func LDY(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Y = cpu.Read(addressInfo.Address)
	cpu.SetZN(cpu.Y)
	return true
}

// STY - Store Y Register
// Function:  memory = X
// Flags Out: None
func STY(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Write(addressInfo.Address, cpu.Y)
	return false
}

//
// Transfer Instructions
//

// TAX - Transfer Accumulator to X Register
// Function:  X = A
// Flags Out: Z, N
func TAX(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.X = cpu.A
	cpu.SetZN(cpu.X)
	return false
}

// TAY - Transfer Accumulator to Y Register
// Function:  Y = A
// Flags Out: Z, N
func TAY(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Y = cpu.A
	cpu.SetZN(cpu.Y)
	return false
}

// TXA - Transfer X Register to Accumulator
// Function:  A = X
// Flags Out: Z, N
func TXA(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.A = cpu.X
	cpu.SetZN(cpu.A)
	return false
}

// TYA - Transfer Y Register to Accumulator
// Function:  A = Y
// Flags Out: Z, N
func TYA(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.A = cpu.Y
	cpu.SetZN(cpu.A)
	return false
}

//
// Arithmetic Instructions
//

// ADC - Add with Carry
// Function:  A = A + memory + C
// Flags Out: C, Z, V, N
func ADC(cpu *CPU, addressInfo AddressInfo) bool {
	value := cpu.Read(addressInfo.Address)
	temp := uint16(cpu.A) + uint16(value) + ternary(cpu.GetFlag(C), uint16(1), uint16(0))
	cpu.SetFlag(C, temp > 0xFF)
	cpu.SetFlag(V, (cpu.A^value)&0x80 == 0 && (uint16(cpu.A)^temp)&0x80 != 0)
	cpu.A = uint8(temp)
	cpu.SetZN(cpu.A)
	return true
}

// SBC - Subtract with Carry
// Function: A = A - memory - !C
// Flags Out: C, Z, V, N
func SBC(cpu *CPU, addressInfo AddressInfo) bool {
	value := cpu.Read(addressInfo.Address)
	temp := uint16(cpu.A) - uint16(value) - ternary(cpu.GetFlag(C), uint16(0), uint16(1))
	cpu.SetFlag(C, temp <= 0xFF)
	cpu.SetFlag(V, (cpu.A^value)&0x80 != 0 && (uint16(cpu.A)^temp)&0x80 != 0)
	cpu.A = uint8(temp)
	cpu.SetZN(cpu.A)
	return true
}

// INC - Increment Memory
// Function: memory = memory + 1
// Flags Out: Z, N
func INC(cpu *CPU, addressInfo AddressInfo) bool {
	value := cpu.Read(addressInfo.Address)
	value++
	cpu.Write(addressInfo.Address, value)
	cpu.SetZN(value)
	return false
}

// DEC - Decrement Memory
// Function: memory = memory - 1
// Flags Out: Z, N
func DEC(cpu *CPU, addressInfo AddressInfo) bool {
	value := cpu.Read(addressInfo.Address)
	value--
	cpu.Write(addressInfo.Address, value)
	cpu.SetZN(value)
	return false
}

// INX - Increment X Register
// Function: X = X + 1
// Flags Out: Z, N
func INX(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.X++
	cpu.SetZN(cpu.X)
	return false
}

// DEX - Decrement X Register
// Function: X = X - 1
// Flags Out: Z, N
func DEX(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.X--
	cpu.SetZN(cpu.X)
	return false
}

// INY - Increment Y Register
// Function: Y = Y + 1
// Flags Out: Z, N
func INY(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Y++
	cpu.SetZN(cpu.Y)
	return false
}

// DEY - Decrement Y Register
// Function: Y = Y - 1
// Flags Out: Z, N
func DEY(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Y--
	cpu.SetZN(cpu.Y)
	return false
}

//
// Shift Instructions
//

// ASL - Arithmetic Shift Left
// Function: value = value << 1
// Flags Out: C, Z, N
func ASL(cpu *CPU, addressInfo AddressInfo) bool {
	if addressInfo.IsAccumulator {
		cpu.SetFlag(C, (cpu.A&0x80) != 0)
		cpu.A <<= 1
		cpu.SetZN(cpu.A)
	} else {
		value := cpu.Read(addressInfo.Address)
		cpu.SetFlag(C, (value&0x80) != 0)
		value <<= 1
		cpu.Write(addressInfo.Address, value)
		cpu.SetZN(value)
	}
	return false
}

// LSR - Logical Shift Right
// Function: value = value >> 1
// Flags Out: C, Z, N
func LSR(cpu *CPU, addressInfo AddressInfo) bool {
	if addressInfo.IsAccumulator {
		cpu.SetFlag(C, (cpu.A&0x01) != 0)
		cpu.A >>= 1
		cpu.SetZN(cpu.A)
	} else {
		value := cpu.Read(addressInfo.Address)
		cpu.SetFlag(C, (value&0x01) != 0)
		value >>= 1
		cpu.Write(addressInfo.Address, value)
		cpu.SetZN(value)
	}
	return false
}

// ROL - Rotate Left
// Function: value = value << 1 through C
// Flags Out: C, Z, N
func ROL(cpu *CPU, addressInfo AddressInfo) bool {
	if addressInfo.IsAccumulator {
		c := cpu.GetFlag(C)
		cpu.SetFlag(C, (cpu.A&0x80) != 0)
		cpu.A = (cpu.A << 1) | (ternary(c, byte(1), byte(0)))
		cpu.SetZN(cpu.A)
	} else {
		c := cpu.GetFlag(C)
		value := cpu.Read(addressInfo.Address)
		cpu.SetFlag(C, (value&0x80) != 0)
		value = (value << 1) | (ternary(c, byte(1), byte(0)))
		cpu.Write(addressInfo.Address, value)
		cpu.SetZN(value)
	}
	return false
}

// ROR - Rotate Right
// Function: value = value >> 1 through C
// Flags Out: C, Z, N
func ROR(cpu *CPU, addressInfo AddressInfo) bool {
	if addressInfo.IsAccumulator {
		c := cpu.GetFlag(C)
		cpu.SetFlag(C, (cpu.A&0x01) != 0)
		cpu.A = (cpu.A >> 1) | ((ternary(c, byte(1), byte(0))) << 7)
		cpu.SetZN(cpu.A)
	} else {
		c := cpu.GetFlag(C)
		value := cpu.Read(addressInfo.Address)
		cpu.SetFlag(C, (value&0x01) != 0)
		value = (value >> 1) | ((ternary(c, byte(1), byte(0))) << 7)
		cpu.Write(addressInfo.Address, value)
		cpu.SetZN(value)
	}
	return false
}

//
// Bitwise Instructions
//

// AND - Bitwise AND
// Function: A = A & memory
// Flags Out: Z, N
func AND(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.A = cpu.A & cpu.Read(addressInfo.Address)
	cpu.SetZN(cpu.A)
	return true
}

// ORA - Bitwise OR
// Function: A = A | memory
// Flags Out: Z, N
func ORA(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.A = cpu.A | cpu.Read(addressInfo.Address)
	cpu.SetZN(cpu.A)
	return true
}

// EOR - Bitwise XOR
// Function: A = A ^ memory
// Flags Out: Z, N
func EOR(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.A = cpu.A ^ cpu.Read(addressInfo.Address)
	cpu.SetZN(cpu.A)
	return true
}

// BIT - Bit Test
// Function: A & memory
// Flags Out: Z, V, N
func BIT(cpu *CPU, addressInfo AddressInfo) bool {
	value := cpu.Read(addressInfo.Address)
	cpu.SetFlag(Z, (cpu.A&value) == 0x00)
	cpu.SetFlag(N, value&0x80 != 0)
	cpu.SetFlag(V, value&0x40 != 0)
	return false
}

//
// Compare Instructions
//

// CMP - Compare Accumulator
// Function: A - memory
// Flags Out: C, Z, N
func CMP(cpu *CPU, addressInfo AddressInfo) bool {
	value := cpu.Read(addressInfo.Address)
	temp := uint16(cpu.A) - uint16(value)
	cpu.SetFlag(C, cpu.A >= value)
	cpu.SetZN(byte(temp & 0x00FF))
	return true
}

// CPX - Compare X Register
// Function: X - memory
// Flags Out: C, Z, N
func CPX(cpu *CPU, addressInfo AddressInfo) bool {
	value := cpu.Read(addressInfo.Address)
	temp := uint16(cpu.X) - uint16(value)
	cpu.SetFlag(C, cpu.X >= value)
	cpu.SetZN(byte(temp & 0x00FF))
	return false
}

// CPY - Compare Y Register
// Function: Y - memory
// Flags Out: C, Z, N
func CPY(cpu *CPU, addressInfo AddressInfo) bool {
	value := cpu.Read(addressInfo.Address)
	temp := uint16(cpu.Y) - uint16(value)
	cpu.SetFlag(C, cpu.Y >= value)
	cpu.SetZN(byte(temp & 0x00FF))
	return false
}

//
// Branch Instructions
//

// BCC - Branch if Carry Clear
// Function = PC = PC + 2 + memory (signed)
func BCC(cpu *CPU, addressInfo AddressInfo) bool {
	if !cpu.GetFlag(C) {
		cpu.addBranchCycles(addressInfo)
		cpu.PC = addressInfo.Address
	}
	return false
}

// BCS - Branch if Carry Set
// Function = PC = PC + 2 + memory (signed)
func BCS(cpu *CPU, addressInfo AddressInfo) bool {
	if cpu.GetFlag(C) {
		cpu.addBranchCycles(addressInfo)
		cpu.PC = addressInfo.Address
	}
	return false
}

// BEQ - Branch if Equal
// Function = PC = PC + 2 + memory (signed)
func BEQ(cpu *CPU, addressInfo AddressInfo) bool {
	if cpu.GetFlag(Z) {
		cpu.addBranchCycles(addressInfo)
		cpu.PC = addressInfo.Address
	}
	return false
}

// BNE - Branch if Not Equal
// Function = PC = PC + 2 + memory (signed)
func BNE(cpu *CPU, addressInfo AddressInfo) bool {
	if !cpu.GetFlag(Z) {
		cpu.addBranchCycles(addressInfo)
		cpu.PC = addressInfo.Address
	}
	return false
}

// BPL - Branch if Positive
// Function = PC = PC + 2 + memory (signed)
func BPL(cpu *CPU, addressInfo AddressInfo) bool {
	if !cpu.GetFlag(N) {
		cpu.addBranchCycles(addressInfo)
		cpu.PC = addressInfo.Address
	}
	return false
}

// BMI - Branch if Negative
// Function = PC = PC + 2 + memory (signed)
func BMI(cpu *CPU, addressInfo AddressInfo) bool {
	if cpu.GetFlag(N) {
		cpu.addBranchCycles(addressInfo)
		cpu.PC = addressInfo.Address
	}
	return false
}

// BVC - Branch if Overflow Clear
func BVC(cpu *CPU, addressInfo AddressInfo) bool {
	if !cpu.GetFlag(V) {
		cpu.addBranchCycles(addressInfo)
		cpu.PC = addressInfo.Address
	}
	return false
}

// BVS - Branch if Overflow Set
// Function = PC = PC + 2 + memory (signed)
func BVS(cpu *CPU, addressInfo AddressInfo) bool {
	if cpu.GetFlag(V) {
		cpu.addBranchCycles(addressInfo)
		cpu.PC = addressInfo.Address
	}
	return false
}

//
// Jump Instructions
//

// JMP - Jump
// Function = PC = memory
func JMP(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.PC = addressInfo.Address
	return false
}

// JSR - Jump to Subroutine
// Function:
//
//	push PC + 2 high byte to stack
//	push PC + 2 low byte to stack
//	PC = memory
//
// Note: the return address on the stack points 1 byte before the start of the next instruction. However, the
// clock function will have already incremented the PC to start of the next instruction so we need to -1 here.
func JSR(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Push16(cpu.PC - 1)
	cpu.PC = addressInfo.Address
	return false
}

// RTS - Return from Subroutine
// Function:
//
//	pull PC low byte from stack
//	pull PC high byte from stack
//	PC = PC + 1
func RTS(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.PC = cpu.Pop16() + 1
	return false
}

// BRK - Force Interrupt (IRQ)
// Function:
//
//	push PC + 2 high byte to stack
//	push PC + 2 low byte to stack
//	push NV11DIZC flags to stack
//	PC = ($FFFE)
func BRK(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Push16(cpu.PC)
	cpu.Push(cpu.Status | 0x10) // 0x10 sets the Break flag to 1 (but only in the value pushed to the stack)
	cpu.SetFlag(I, true)        // Set the "Interrupt Disable" flag
	cpu.PC = cpu.IRQVector()    // Read a value from 0xFFFE and use this as the memory address to jump to
	return false
}

// RTI - Return from Interrupt
// Function:
//
//	pull NVxxDIZC flags from stack
//	pull PC low byte from stack
//	pull PC high byte from stack
func RTI(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Status = cpu.Pop()
	cpu.SetFlag(B, false)
	cpu.SetFlag(U, true)
	cpu.PC = cpu.Pop16()
	return false
}

//
// Stack Instructions
//

// PHA - Push A
// Function:
//
//	($0100 + SP) = A
//	SP = SP - 1
func PHA(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Push(cpu.A)
	return false
}

// PLA - Pull A
// Function:
//
//	SP = SP + 1
//	A = ($0100 + SP)
//
// Flags Out: Z, N
func PLA(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.A = cpu.Pop()
	cpu.SetZN(cpu.A)
	return false
}

// PHP - Push Processor Status (status flags)
// Function:
//
//	($0100 + SP) = NV11DIZC
//	SP = SP - 1
func PHP(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Push(cpu.Status | 0x10) // 0x10 sets the Break flag to 1 (but only in the value pushed to the stack)
	return false
}

// PLP - Pull Processor Status (status flags)
// Function:
//
//	SP = SP + 1
//	NVxxDIZC = ($0100 + SP)
//
// Flags out: C, Z, I, D, V, N
func PLP(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.Status = cpu.Pop()
	cpu.SetFlag(B, false)
	cpu.SetFlag(U, true)
	return false
}

// TXS - Transfer X to Stack Pointer
// Function: SP = X
func TXS(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.SP = cpu.X
	return false
}

// TSX - Transfer Stack Pointer to X
// Function: SP = X
// Flags Out: Z, N
func TSX(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.X = cpu.SP
	cpu.SetZN(cpu.X)
	return false
}

//
// Flag Instructions
//

// CLC - Clear Carry
// Function:  C = 0
// Flags Out: C
func CLC(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.SetFlag(C, false)
	return false
}

// SEC - Set Carry
// Function:  C = 1
// Flags Out: C
func SEC(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.SetFlag(C, true)
	return false
}

// CLI - Clear Interrupt Disable
// Function:  I = 0
// Flags Out: I
func CLI(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.SetFlag(I, false)
	return false
}

// SEI - Set Interrupt Disable
// Function:  I = 0
// Flags Out: I
func SEI(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.SetFlag(I, true)
	return false
}

// CLD - Clear Decimal
// Function:  D = 0
// Flags Out: D
func CLD(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.SetFlag(D, false)
	return false
}

// SED - Set Decimal
// Function:  D = 1
// Flags Out: D
func SED(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.SetFlag(D, true)
	return false
}

// CLV - Clear Overflow
// Function:  V = 0
// Flags Out: V
func CLV(cpu *CPU, addressInfo AddressInfo) bool {
	cpu.SetFlag(V, false)
	return false
}

//
// Other Instructions
//

// NOP - No operation
func NOP(cpu *CPU, addressInfo AddressInfo) bool {
	return false
}

// XXX captures illegal opcodes
func XXX(cpu *CPU, addressInfo AddressInfo) bool {
	return false
}
