package processor

type Operation struct {
	Instruction InstructionFunc
	AddressMode AddressModeFunc
	Size        uint8
	Cycles      uint8
}

// operations is the lookup table for all 6502 instructions.
// It is 16x16 entries which gives 256 instructions. It is arranged so that the bottom
// 4 bits of the opcode choose the column, and the top 4 bits choose the row.
//
// Note that "illegal" opcodes are not currently implemented and are treated as a 1-byte NOP instruction.
var operations = [...]Operation{
	{BRK, IMM, 1, 7}, {ORA, INDX, 2, 6}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {ORA, ZP0, 2, 3}, {ASL, ZP0, 2, 5}, {XXX, IMP, 1, 1}, {PHP, IMP, 1, 3}, {ORA, IMM, 2, 2}, {ASL, ACC, 1, 2}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {ORA, ABS, 3, 4}, {ASL, ABS, 3, 6}, {XXX, IMP, 1, 1},
	{BPL, REL, 2, 2}, {ORA, INDY, 2, 5}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {ORA, ZPX, 2, 4}, {ASL, ZPX, 2, 6}, {XXX, IMP, 1, 1}, {CLC, IMP, 1, 2}, {ORA, ABY, 3, 4}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {ORA, ABX, 3, 4}, {ASL, ABX, 3, 7}, {XXX, IMP, 1, 1},
	{JSR, ABS, 3, 6}, {AND, INDX, 2, 6}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {BIT, ZP0, 2, 3}, {AND, ZP0, 2, 3}, {ROL, ZP0, 2, 5}, {XXX, IMP, 1, 1}, {PLP, IMP, 1, 4}, {AND, IMM, 2, 2}, {ROL, ACC, 1, 2}, {XXX, IMP, 1, 1}, {BIT, ABS, 3, 4}, {AND, ABS, 3, 4}, {ROL, ABS, 3, 6}, {XXX, IMP, 1, 1},
	{BMI, REL, 2, 2}, {AND, INDY, 2, 5}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {AND, ZPX, 2, 4}, {ROL, ZPX, 2, 6}, {XXX, IMP, 1, 1}, {SEC, IMP, 1, 2}, {AND, ABY, 3, 4}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {AND, ABX, 3, 4}, {ROL, ABX, 3, 7}, {XXX, IMP, 1, 1},
	{RTI, IMP, 1, 6}, {EOR, INDX, 2, 6}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {EOR, ZP0, 2, 3}, {LSR, ZP0, 2, 5}, {XXX, IMP, 1, 1}, {PHA, IMP, 1, 3}, {EOR, IMM, 2, 2}, {LSR, ACC, 1, 2}, {XXX, IMP, 1, 1}, {JMP, ABS, 3, 3}, {EOR, ABS, 3, 4}, {LSR, ABS, 3, 6}, {XXX, IMP, 1, 1},
	{BVC, REL, 2, 2}, {EOR, INDY, 2, 5}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {EOR, ZPX, 2, 4}, {LSR, ZPX, 2, 6}, {XXX, IMP, 1, 1}, {CLI, IMP, 1, 2}, {EOR, ABY, 3, 4}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {EOR, ABX, 3, 4}, {LSR, ABX, 3, 7}, {XXX, IMP, 1, 1},
	{RTS, IMP, 1, 6}, {ADC, INDX, 2, 6}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {ADC, ZP0, 2, 3}, {ROR, ZP0, 2, 5}, {XXX, IMP, 1, 1}, {PLA, IMP, 1, 4}, {ADC, IMM, 2, 2}, {ROR, ACC, 1, 2}, {XXX, IMP, 1, 1}, {JMP, IND, 3, 5}, {ADC, ABS, 3, 4}, {ROR, ABS, 3, 6}, {XXX, IMP, 1, 1},
	{BVS, REL, 2, 2}, {ADC, INDY, 2, 5}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {ADC, ZPX, 2, 4}, {ROR, ZPX, 2, 6}, {XXX, IMP, 1, 1}, {SEI, IMP, 1, 2}, {ADC, ABY, 3, 4}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {ADC, ABX, 3, 4}, {ROR, ABX, 3, 7}, {XXX, IMP, 1, 1},
	{XXX, IMP, 1, 1}, {STA, INDX, 2, 6}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {STY, ZP0, 2, 3}, {STA, ZP0, 2, 3}, {STX, ZP0, 2, 3}, {XXX, IMP, 1, 1}, {DEY, IMP, 1, 2}, {XXX, IMP, 1, 1}, {TXA, IMP, 1, 2}, {XXX, IMP, 1, 1}, {STY, ABS, 3, 4}, {STA, ABS, 3, 4}, {STX, ABS, 3, 4}, {XXX, IMP, 1, 1},
	{BCC, REL, 2, 2}, {STA, INDY, 2, 6}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {STY, ZPX, 2, 4}, {STA, ZPX, 2, 4}, {STX, ZPY, 2, 4}, {XXX, IMP, 1, 1}, {TYA, IMP, 1, 2}, {STA, ABY, 3, 5}, {TXS, IMP, 1, 2}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {STA, ABX, 3, 5}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1},
	{LDY, IMM, 2, 2}, {LDA, INDX, 2, 6}, {LDX, IMM, 2, 2}, {XXX, IMP, 1, 1}, {LDY, ZP0, 2, 3}, {LDA, ZP0, 2, 3}, {LDX, ZP0, 2, 3}, {XXX, IMP, 1, 1}, {TAY, IMP, 1, 2}, {LDA, IMM, 2, 2}, {TAX, IMP, 1, 2}, {XXX, IMP, 1, 1}, {LDY, ABS, 3, 4}, {LDA, ABS, 3, 4}, {LDX, ABS, 3, 4}, {XXX, IMP, 1, 1},
	{BCS, REL, 2, 2}, {LDA, INDY, 2, 5}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {LDY, ZPX, 2, 4}, {LDA, ZPX, 2, 4}, {LDX, ZPY, 2, 4}, {XXX, IMP, 1, 1}, {CLV, IMP, 1, 2}, {LDA, ABY, 3, 4}, {TSX, IMP, 1, 2}, {XXX, IMP, 1, 1}, {LDY, ABX, 3, 4}, {LDA, ABX, 3, 4}, {LDX, ABY, 3, 4}, {XXX, IMP, 1, 1},
	{CPY, IMM, 2, 2}, {CMP, INDX, 2, 6}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {CPY, ZP0, 2, 3}, {CMP, ZP0, 2, 3}, {DEC, ZP0, 2, 5}, {XXX, IMP, 1, 1}, {INY, IMP, 1, 2}, {CMP, IMM, 2, 2}, {DEX, IMP, 1, 2}, {XXX, IMP, 1, 1}, {CPY, ABS, 3, 4}, {CMP, ABS, 3, 4}, {DEC, ABS, 3, 6}, {XXX, IMP, 1, 1},
	{BNE, REL, 2, 2}, {CMP, INDY, 2, 5}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {CMP, ZPX, 2, 4}, {DEC, ZPX, 2, 6}, {XXX, IMP, 1, 1}, {CLD, IMP, 1, 2}, {CMP, ABY, 3, 4}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {CMP, ABX, 3, 4}, {DEC, ABX, 3, 7}, {XXX, IMP, 1, 1},
	{CPX, IMM, 2, 2}, {SBC, INDX, 2, 6}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {CPX, ZP0, 2, 3}, {SBC, ZP0, 2, 3}, {INC, ZP0, 2, 5}, {XXX, IMP, 1, 1}, {INX, IMP, 1, 2}, {SBC, IMM, 2, 2}, {NOP, IMP, 1, 2}, {XXX, IMP, 1, 1}, {CPX, ABS, 3, 4}, {SBC, ABS, 3, 4}, {INC, ABS, 3, 6}, {XXX, IMP, 1, 1},
	{BEQ, REL, 2, 2}, {SBC, INDY, 2, 5}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {SBC, ZPX, 2, 4}, {INC, ZPX, 2, 6}, {XXX, IMP, 1, 1}, {SED, IMP, 1, 2}, {SBC, ABY, 3, 4}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {XXX, IMP, 1, 1}, {SBC, ABX, 3, 4}, {INC, ABX, 3, 7}, {XXX, IMP, 1, 1},
}
