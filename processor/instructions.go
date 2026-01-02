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

//
// Transfer Instructions
//

//
// Arithmetic Instructions
//

//
// Shift Instructions
//

//
// Bitwise Instructions
//

// AND - Bitwise Logic AND

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
