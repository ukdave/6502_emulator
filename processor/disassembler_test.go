package processor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ukdave/6502_emulator/bus"
	"github.com/ukdave/6502_emulator/processor"
)

type DisassembleOperationSuite struct {
	suite.Suite
	bus bus.Bus
	cpu *processor.CPU
}

func TestDisassembleOperationSuite(t *testing.T) {
	suite.Run(t, new(DisassembleOperationSuite))
}

func (suite *DisassembleOperationSuite) SetupTest() {
	suite.bus = bus.NewSimpleBus()
	suite.cpu = processor.NewCPU(suite.bus)
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_Accumulator() {
	// ASL ACC (0x0A at address 0x0000)
	suite.bus.Write(0x0000, 0x0A)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0x0A}, result.Bytes, "Expected operand to be [0x0A]")
	assert.Equal(suite.T(), uint16(0), result.Operand, "Expected operand to be 0")
	assert.Equal(suite.T(), "ASL", result.Operation.Name(), "Expected operation name to be ASL")
	assert.Equal(suite.T(), "ACC", result.Operation.AddressModeName(), "Expected address mode to be ACC")
	assert.Equal(suite.T(), "ASL A {ACC}", result.Disassembly, "Expected disassembly to be 'ASL A {ACC}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_Immediate() {
	// LDA #$42 (0xA9 0x42 at address 0x0000)
	suite.bus.Write(0x0000, 0xA9)
	suite.bus.Write(0x0001, 0x42)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0xA9, 0x42}, result.Bytes, "Expected operand to be [0xA9, 0x42]")
	assert.Equal(suite.T(), uint16(0x42), result.Operand, "Expected operand to be 0x42")
	assert.Equal(suite.T(), "LDA", result.Operation.Name(), "Expected operation name to be LDA")
	assert.Equal(suite.T(), "IMM", result.Operation.AddressModeName(), "Expected address mode to be IMM")
	assert.Equal(suite.T(), "LDA #$42 {IMM}", result.Disassembly, "Expected disassembly to be 'LDA #$42 {IMM}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_Absolute() {
	// LDA $1234 (0xAD 0x34 0x12 at address 0x0000)
	suite.bus.Write(0x0000, 0xAD)
	suite.bus.Write(0x0001, 0x34)
	suite.bus.Write(0x0002, 0x12)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0xAD, 0x34, 0x12}, result.Bytes, "Expected operand to be [0xAD, 0x34, 0x12]")
	assert.Equal(suite.T(), uint16(0x1234), result.Operand, "Expected operand to be 0x1234")
	assert.Equal(suite.T(), "LDA", result.Operation.Name(), "Expected operation name to be LDA")
	assert.Equal(suite.T(), "ABS", result.Operation.AddressModeName(), "Expected address mode to be ABS")
	assert.Equal(suite.T(), "LDA $1234 {ABS}", result.Disassembly, "Expected disassembly to be 'LDA $1234 {ABS}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_AbsoluteXIndexed() {
	// LDA $1234,X (0xBD 0x34 0x12 at address 0x0000)
	suite.bus.Write(0x0000, 0xBD)
	suite.bus.Write(0x0001, 0x34)
	suite.bus.Write(0x0002, 0x12)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0xBD, 0x34, 0x12}, result.Bytes, "Expected operand to be [0xBD, 0x34, 0x12]")
	assert.Equal(suite.T(), uint16(0x1234), result.Operand, "Expected operand to be 0x1234")
	assert.Equal(suite.T(), "LDA", result.Operation.Name(), "Expected operation name to be LDA")
	assert.Equal(suite.T(), "ABX", result.Operation.AddressModeName(), "Expected address mode to be ABX")
	assert.Equal(suite.T(), "LDA $1234,X {ABX}", result.Disassembly, "Expected disassembly to be 'LDA $1234,X {ABX}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_AbsoluteYIndexed() {
	// LDA $1234,Y (0xB9 0x34 0x12 at address 0x0000)
	suite.bus.Write(0x0000, 0xB9)
	suite.bus.Write(0x0001, 0x34)
	suite.bus.Write(0x0002, 0x12)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0xB9, 0x34, 0x12}, result.Bytes, "Expected operand to be [0xB9, 0x34, 0x12]")
	assert.Equal(suite.T(), uint16(0x1234), result.Operand, "Expected operand to be 0x1234")
	assert.Equal(suite.T(), "LDA", result.Operation.Name(), "Expected operation name to be LDA")
	assert.Equal(suite.T(), "ABY", result.Operation.AddressModeName(), "Expected address mode to be ABY")
	assert.Equal(suite.T(), "LDA $1234,Y {ABY}", result.Disassembly, "Expected disassembly to be 'LDA $1234,Y {ABY}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_ZeroPage() {
	// LDA $42 (0xA5 0x42 at address 0x0000)
	suite.bus.Write(0x0000, 0xA5)
	suite.bus.Write(0x0001, 0x42)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0xA5, 0x42}, result.Bytes, "Expected operand to be [0xA5, 0x42]")
	assert.Equal(suite.T(), uint16(0x42), result.Operand, "Expected operand to be 0x42")
	assert.Equal(suite.T(), "LDA", result.Operation.Name(), "Expected operation name to be LDA")
	assert.Equal(suite.T(), "ZP0", result.Operation.AddressModeName(), "Expected address mode to be ZP0")
	assert.Equal(suite.T(), "LDA $42 {ZP0}", result.Disassembly, "Expected disassembly to be 'LDA $42 {ZP0}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_ZeroPageXIndexed() {
	// LDA $42,X (0xB5 0x42 at address 0x0000)
	suite.bus.Write(0x0000, 0xB5)
	suite.bus.Write(0x0001, 0x42)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0xB5, 0x42}, result.Bytes, "Expected operand to be [0xB5, 0x42]")
	assert.Equal(suite.T(), uint16(0x42), result.Operand, "Expected operand to be 0x42")
	assert.Equal(suite.T(), "LDA", result.Operation.Name(), "Expected operation name to be LDA")
	assert.Equal(suite.T(), "ZPX", result.Operation.AddressModeName(), "Expected address mode to be ZPX")
	assert.Equal(suite.T(), "LDA $42,X {ZPX}", result.Disassembly, "Expected disassembly to be 'LDA $42,X {ZPX}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_ZeroPageYIndexed() {
	// STX $42,Y (0x96 0x42 at address 0x0000)
	suite.bus.Write(0x0000, 0x96)
	suite.bus.Write(0x0001, 0x42)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0x96, 0x42}, result.Bytes, "Expected operand to be [0x96, 0x42]")
	assert.Equal(suite.T(), uint16(0x42), result.Operand, "Expected operand to be 0x42")
	assert.Equal(suite.T(), "STX", result.Operation.Name(), "Expected operation name to be STX")
	assert.Equal(suite.T(), "ZPY", result.Operation.AddressModeName(), "Expected address mode to be ZPY")
	assert.Equal(suite.T(), "STX $42,Y {ZPY}", result.Disassembly, "Expected disassembly to be 'STX $42,Y {ZPY}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_Implied() {
	// NOP (0xEA at address 0x0000)
	suite.bus.Write(0x0000, 0xEA)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0xEA}, result.Bytes, "Expected operand to be [0xEA]")
	assert.Equal(suite.T(), uint16(0), result.Operand, "Expected operand to be 0")
	assert.Equal(suite.T(), "NOP", result.Operation.Name(), "Expected operation name to be NOP")
	assert.Equal(suite.T(), "IMP", result.Operation.AddressModeName(), "Expected address mode to be IMP")
	assert.Equal(suite.T(), "NOP {IMP}", result.Disassembly, "Expected disassembly to be 'NOP {IMP}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_RelativePositiveOffset() {
	// BEQ $05 (0xF0 0x05 at address 0x2000)
	// Target address = 0x2000 + 2 (instruction size) + 0x05 = 0x2007
	suite.bus.Write(0x2000, 0xF0)
	suite.bus.Write(0x2001, 0x05)

	result := suite.cpu.DisassembleOperation(0x2000)

	assert.Equal(suite.T(), []byte{0xF0, 0x05}, result.Bytes, "Expected operand to be [0xF0, 0x05]")
	assert.Equal(suite.T(), uint16(0x05), result.Operand, "Expected operand to be 0x05")
	assert.Equal(suite.T(), "BEQ", result.Operation.Name(), "Expected operation name to be BEQ")
	assert.Equal(suite.T(), "REL", result.Operation.AddressModeName(), "Expected address mode to be REL")
	assert.Equal(suite.T(), "BEQ $05 [$2007] {REL}", result.Disassembly, "Expected disassembly to be 'BEQ $05 [$2007] {REL}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_RelativeNegativeOffset() {
	// BEQ -5 (0xF0 0xFB at address 0x2005)
	// Target address = 0x2005 + 2 (instruction size) + 0xFB - 0x100 = 0x2002
	suite.bus.Write(0x2005, 0xF0)
	suite.bus.Write(0x2006, 0xFB)

	result := suite.cpu.DisassembleOperation(0x2005)

	assert.Equal(suite.T(), []byte{0xF0, 0xFB}, result.Bytes, "Expected operand to be [0xF0, 0xFB]")
	assert.Equal(suite.T(), uint16(0xFB), result.Operand, "Expected operand to be 0xFB")
	assert.Equal(suite.T(), "BEQ", result.Operation.Name(), "Expected operation name to be BEQ")
	assert.Equal(suite.T(), "REL", result.Operation.AddressModeName(), "Expected address mode to be REL")
	assert.Equal(suite.T(), "BEQ $FB [$2002] {REL}", result.Disassembly, "Expected disassembly to be 'BEQ $FB [$2002] {REL}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_Indirect() {
	// JMP ($1234) (0x6C 0x34 0x12 at address 0x0000)
	suite.bus.Write(0x0000, 0x6C)
	suite.bus.Write(0x0001, 0x34)
	suite.bus.Write(0x0002, 0x12)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0x6C, 0x34, 0x12}, result.Bytes, "Expected operand to be [0x6C, 0x34, 0x12]")
	assert.Equal(suite.T(), uint16(0x1234), result.Operand, "Expected operand to be 0x1234")
	assert.Equal(suite.T(), "JMP", result.Operation.Name(), "Expected operation name to be JMP")
	assert.Equal(suite.T(), "IND", result.Operation.AddressModeName(), "Expected address mode to be IND")
	assert.Equal(suite.T(), "JMP ($1234) {IND}", result.Disassembly, "Expected disassembly to be 'JMP ($1234) {IND}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_IndexedIndirectX() {
	// LDA ($40,X) (0xA1 0x40 at address 0x0000)
	suite.bus.Write(0x0000, 0xA1)
	suite.bus.Write(0x0001, 0x40)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0xA1, 0x40}, result.Bytes, "Expected operand to be [0xA1, 0x40]")
	assert.Equal(suite.T(), uint16(0x40), result.Operand, "Expected operand to be 0x40")
	assert.Equal(suite.T(), "LDA", result.Operation.Name(), "Expected operation name to be LDA")
	assert.Equal(suite.T(), "INDX", result.Operation.AddressModeName(), "Expected address mode to be INDX")
	assert.Equal(suite.T(), "LDA ($40,X) {INDX}", result.Disassembly, "Expected disassembly to be 'LDA ($40,X) {INDX}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_IndirectIndexedY() {
	// LDA ($40),Y (0xB1 0x40 at address 0x0000)
	suite.bus.Write(0x0000, 0xB1)
	suite.bus.Write(0x0001, 0x40)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0xB1, 0x40}, result.Bytes, "Expected operand to be [0xB1, 0x40]")
	assert.Equal(suite.T(), uint16(0x40), result.Operand, "Expected operand to be 0x40")
	assert.Equal(suite.T(), "LDA", result.Operation.Name(), "Expected operation name to be LDA")
	assert.Equal(suite.T(), "INDY", result.Operation.AddressModeName(), "Expected address mode to be INDY")
	assert.Equal(suite.T(), "LDA ($40),Y {INDY}", result.Disassembly, "Expected disassembly to be 'LDA ($40),Y {INDY}'")
}

func (suite *DisassembleOperationSuite) TestDisassembleOperation_IllegalOpcode() {
	suite.bus.Write(0x0000, 0x02)

	result := suite.cpu.DisassembleOperation(0x0000)

	assert.Equal(suite.T(), []byte{0x02}, result.Bytes, "Expected operand to be [0x02]")
	assert.Equal(suite.T(), uint16(0x00), result.Operand, "Expected operand to be 0x00")
	assert.Equal(suite.T(), "???", result.Operation.Name(), "Expected operation name to be ???")
	assert.Equal(suite.T(), "IMP", result.Operation.AddressModeName(), "Expected address mode to be IMP")
	assert.Equal(suite.T(), "??? {IMP}", result.Disassembly, "Expected disassembly to be '??? {IMP}'")
}
