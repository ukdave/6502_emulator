package processor

import "fmt"

// DisassembledOperation contains the disassembled representation of an operation.
type DisassembledOperation struct {
	Bytes       []byte
	Operand     uint16
	Operation   Operation
	Disassembly string
}

// DisassembleOperation decodes an operation at the given address and returns a DisassembledOperation struct.
func (c *CPU) DisassembleOperation(addr uint16) DisassembledOperation {
	opcode := c.Read(addr)
	op := c.GetOperation(opcode)

	bytes := make([]byte, op.Size)
	for i := 0; i < int(op.Size); i++ {
		bytes[i] = c.Read(addr + uint16(i))
	}

	operand := uint16(0)
	for i := 0; i < len(bytes)-1; i++ {
		operand |= uint16(bytes[i+1]) << (8 * i)
	}

	// Generate the disassembly string
	var disassembly string
	switch op.AddressModeName() {
	case "ACC":
		disassembly = fmt.Sprintf("%s A {%s}", op.Name(), op.AddressModeName())
	case "IMM":
		disassembly = fmt.Sprintf("%s #$%02X {%s}", op.Name(), uint8(operand), op.AddressModeName())
	case "ABS":
		disassembly = fmt.Sprintf("%s $%04X {%s}", op.Name(), operand, op.AddressModeName())
	case "ABX":
		disassembly = fmt.Sprintf("%s $%04X,X {%s}", op.Name(), operand, op.AddressModeName())
	case "ABY":
		disassembly = fmt.Sprintf("%s $%04X,Y {%s}", op.Name(), operand, op.AddressModeName())
	case "ZP0":
		disassembly = fmt.Sprintf("%s $%02X {%s}", op.Name(), uint8(operand), op.AddressModeName())
	case "ZPX":
		disassembly = fmt.Sprintf("%s $%02X,X {%s}", op.Name(), uint8(operand), op.AddressModeName())
	case "ZPY":
		disassembly = fmt.Sprintf("%s $%02X,Y {%s}", op.Name(), uint8(operand), op.AddressModeName())
	case "IMP":
		disassembly = fmt.Sprintf("%s {%s}", op.Name(), op.AddressModeName())
	case "REL":
		offset := uint8(operand)
		// Calculate relative address from instruction address
		targetAddr := addr + uint16(op.Size) + uint16(offset)
		if offset >= 0x80 {
			targetAddr -= 0x100
		}
		disassembly = fmt.Sprintf("%s $%02X [$%04X] {%s}", op.Name(), offset, targetAddr, op.AddressModeName())
	case "IND":
		disassembly = fmt.Sprintf("%s ($%04X) {%s}", op.Name(), operand, op.AddressModeName())
	case "INDX":
		disassembly = fmt.Sprintf("%s ($%02X,X) {%s}", op.Name(), uint8(operand), op.AddressModeName())
	case "INDY":
		disassembly = fmt.Sprintf("%s ($%02X),Y {%s}", op.Name(), uint8(operand), op.AddressModeName())
	default:
		disassembly = fmt.Sprintf("%s {%s}", op.Name(), op.AddressModeName())
	}

	return DisassembledOperation{
		Bytes:       bytes,
		Operand:     operand,
		Operation:   op,
		Disassembly: disassembly,
	}
}
