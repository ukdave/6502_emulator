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

//
// Branch Instructions
//

//
// Jump Instructions
//

//
// Stack Instructions
//

//
// Flag Instructions
//

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
