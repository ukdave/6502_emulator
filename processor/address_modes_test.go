package processor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ukdave/6502_emulator/bus"
	"github.com/ukdave/6502_emulator/processor"
)

type AddressModesSuite struct {
	suite.Suite
	bus bus.Bus
	cpu *processor.CPU
}

func TestAddressModesSuite(t *testing.T) {
	suite.Run(t, new(AddressModesSuite))
}

func (suite *AddressModesSuite) SetupTest() {
	suite.bus = bus.NewSimpleBus()
	suite.cpu = processor.NewCPU(suite.bus)
}

func (suite *AddressModesSuite) TestACC() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xAA) // some random data
	suite.bus.Write(0x0001, 0xBB) // some more random data
	suite.bus.Write(0x0002, 0xCC) // even more random data

	// Set the Program Counter to 0x0001
	suite.cpu.PC = 0x0001

	// Use accumulator addressing mode
	result := processor.ACC(suite.cpu)

	assert.Equal(suite.T(), uint16(0x0000), result.Address, "Expected address to be 0x0000")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.True(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be true")
}

func (suite *AddressModesSuite) TestIMM() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xAA) // some random data
	suite.bus.Write(0x0001, 0xBB) // some more random data
	suite.bus.Write(0x0002, 0xCC) // even more random data

	// Set the Program Counter to 0x0001
	suite.cpu.PC = 0x0001

	// Use immediate addressing mode to get the target address
	result := processor.IMM(suite.cpu)

	assert.Equal(suite.T(), uint16(0x0002), result.Address, "Expected address to be 0x0002")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestABS() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xAA) // some random data
	suite.bus.Write(0x0001, 0xBB) // some more random data
	suite.bus.Write(0x0002, 0xCC) // even more random data
	suite.bus.Write(0x0003, 0xDD) // yet more random data

	// Set the Program Counter to 0x0000
	suite.cpu.PC = 0x0001

	// Use absolute addressing mode to get the target address
	result := processor.ABS(suite.cpu)

	assert.Equal(suite.T(), uint16(0xDDCC), result.Address, "Expected address to be 0xDDCC")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestABX() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xA0) // some random data
	suite.bus.Write(0x0001, 0xB0) // some more random data
	suite.bus.Write(0x0002, 0xC0) // even more random data
	suite.bus.Write(0x0003, 0xD0) // yet more random data

	// Set the Program Counter to 0x0000
	suite.cpu.PC = 0x0001
	suite.cpu.X = 0x02

	// Use absolute with X offset addressing mode to get the target address
	result := processor.ABX(suite.cpu)

	assert.Equal(suite.T(), uint16(0xD0C2), result.Address, "Expected address to be 0xD0C2")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestABX_PageChange() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xA0) // some random data
	suite.bus.Write(0x0001, 0xB0) // some more random data
	suite.bus.Write(0x0002, 0xFE) // even more random data
	suite.bus.Write(0x0003, 0xD0) // yet more random data

	// Set the Program Counter to 0x0000
	suite.cpu.PC = 0x0001
	suite.cpu.X = 0x02

	// Use absolute with X offset addressing mode to get the target address
	result := processor.ABX(suite.cpu)

	assert.Equal(suite.T(), uint16(0xD100), result.Address, "Expected address to be 0xD100")
	assert.True(suite.T(), result.PageChanged, "Expected pageChanged to be true")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestABY() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xA0) // some random data
	suite.bus.Write(0x0001, 0xB0) // some more random data
	suite.bus.Write(0x0002, 0xC0) // even more random data
	suite.bus.Write(0x0003, 0xD0) // yet more random data

	// Set the Program Counter to 0x0000
	suite.cpu.PC = 0x0001
	suite.cpu.Y = 0x02

	// Use absolute with Y offset addressing mode to get the target address
	result := processor.ABY(suite.cpu)

	assert.Equal(suite.T(), uint16(0xD0C2), result.Address, "Expected address to be 0xD0C2")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestABY_PageChange() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xA0) // some random data
	suite.bus.Write(0x0001, 0xB0) // some more random data
	suite.bus.Write(0x0002, 0xFE) // even more random data
	suite.bus.Write(0x0003, 0xD0) // yet more random data

	// Set the Program Counter to 0x0000
	suite.cpu.PC = 0x0001
	suite.cpu.Y = 0x02

	// Use absolute with Y offset addressing mode to get the target address
	result := processor.ABY(suite.cpu)

	assert.Equal(suite.T(), uint16(0xD100), result.Address, "Expected address to be 0xD100")
	assert.True(suite.T(), result.PageChanged, "Expected pageChanged to be true")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestZP0() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xAA) // some random data
	suite.bus.Write(0x0001, 0xBB) // some more random data
	suite.bus.Write(0x0002, 0xCC) // even more random data

	// Set the Program Counter to 0x0000
	suite.cpu.PC = 0x0001

	// Use zero page addressing mode to get the target address
	result := processor.ZP0(suite.cpu)

	assert.Equal(suite.T(), uint16(0x00CC), result.Address, "Expected address to be 0x00CC")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestZPX() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xA0) // some random data
	suite.bus.Write(0x0001, 0xB0) // some more random data
	suite.bus.Write(0x0002, 0xC0) // even more random data

	// Set the Program Counter to 0x0000
	suite.cpu.PC = 0x0001
	suite.cpu.X = 0x0002

	// Use zero page with X offset addressing mode to get the target address
	result := processor.ZPX(suite.cpu)

	assert.Equal(suite.T(), uint16(0x00C2), result.Address, "Expected address to be 0x00C2")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestZPX_PageWrapping() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xA0) // some random data
	suite.bus.Write(0x0001, 0xB0) // some more random data
	suite.bus.Write(0x0002, 0xFE) // even more random data

	// Set the Program Counter to 0x0000
	suite.cpu.PC = 0x0001
	suite.cpu.X = 0x0004

	// Use zero page with X offset addressing mode to get the target address
	result := processor.ZPX(suite.cpu)

	assert.Equal(suite.T(), uint16(0x0002), result.Address, "Expected address to be 0x0002")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestZPY() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xA0) // some random data
	suite.bus.Write(0x0001, 0xB0) // some more random data
	suite.bus.Write(0x0002, 0xC0) // even more random data

	// Set the Program Counter to 0x0000
	suite.cpu.PC = 0x0001
	suite.cpu.Y = 0x0002

	// Use zero page with Y offset addressing mode to get the target address
	result := processor.ZPY(suite.cpu)

	assert.Equal(suite.T(), uint16(0x00C2), result.Address, "Expected address to be 0x00C2")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestZPY_PageWrapping() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xA0) // some random data
	suite.bus.Write(0x0001, 0xB0) // some more random data
	suite.bus.Write(0x0002, 0xFE) // even more random data

	// Set the Program Counter to 0x0000
	suite.cpu.PC = 0x0001
	suite.cpu.Y = 0x0004

	// Use zero page with Y offset addressing mode to get the target address
	result := processor.ZPY(suite.cpu)

	assert.Equal(suite.T(), uint16(0x0002), result.Address, "Expected address to be 0x0002")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestIMP() {
	// Write some random data to memory at address 0x0000
	suite.bus.Write(0x0000, 0xAA) // some random data
	suite.bus.Write(0x0001, 0xBB) // some more random data
	suite.bus.Write(0x0002, 0xCC) // even more random data

	// Set the Program Counter to 0x0001
	suite.cpu.PC = 0x0001

	// Use implied addressing mode
	result := processor.IMP(suite.cpu)

	assert.Equal(suite.T(), uint16(0x0000), result.Address, "Expected address to be 0x0000")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestREL_PositiveOffset() {
	// Write instruction "BEQ 5" into memory at address 0x2000
	suite.bus.Write(0x2000, 0xF0) // opcode
	suite.bus.Write(0x2001, 0x05) // offset (+5)

	// Set the Program Counter to 0x2000
	suite.cpu.PC = 0x2000

	// Use relative addressing mode to get the target address
	result := processor.REL(suite.cpu)

	assert.Equal(suite.T(), uint16(0x2007), result.Address, "Expected address to be 0x2007")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestREL_PositiveOffsetAndPageChange() {
	// Write instruction "BEQ 5" into memory at address 0x20FA
	suite.bus.Write(0x20FA, 0xF0) // opcode
	suite.bus.Write(0x20FB, 0x05) // offset (+5)

	// Set the Program Counter to 0x20FA
	suite.cpu.PC = 0x20FA

	// Use relative addressing mode to get the target address
	result := processor.REL(suite.cpu)

	assert.Equal(suite.T(), uint16(0x2101), result.Address, "Expected address to be 0x2101")
	assert.True(suite.T(), result.PageChanged, "Expected pageChanged to be be true")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestREL_NegativeOffset() {
	// Write instruction "BEQ -5" into memory at address 0x2005
	suite.bus.Write(0x2005, 0xF0) // opcode
	suite.bus.Write(0x2006, 0xFB) // offset (-5)

	// Set the Program Counter to 0x2005
	suite.cpu.PC = 0x2005

	// Use relative addressing mode to get the target address
	result := processor.REL(suite.cpu)

	assert.Equal(suite.T(), uint16(0x2002), result.Address, "Expected address to be 0x2002")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestREL_NegativeOffsetAndPageChange() {
	// Write instruction "BEQ -5" into memory at address 0x2000
	suite.bus.Write(0x2000, 0xF0) // opcode
	suite.bus.Write(0x2001, 0xFB) // offset (-5)

	// Set the Program Counter to 0x2000
	suite.cpu.PC = 0x2000

	// Use relative addressing mode to get the target address
	result := processor.REL(suite.cpu)

	assert.Equal(suite.T(), uint16(0x1FFD), result.Address, "Expected address to be 0x1FFD")
	assert.True(suite.T(), result.PageChanged, "Expected pageChanged to be be true")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestIND() {
	// Write instruction "JMP $1234" into memory at address 0x0000
	suite.bus.Write(0x0000, 0x6C) // opcode
	suite.bus.Write(0x0001, 0x34) // lo byte
	suite.bus.Write(0x0002, 0x12) // hi byte

	// Write final target address 0x5678 into memory at address 0x1234
	suite.bus.Write(0x1234, 0x78) // lo byte
	suite.bus.Write(0x1235, 0x56) // hi byte

	// Set the Program Counter to the start of the instruction
	suite.cpu.PC = 0x0000

	// Use indirect addressing mode to get the target address
	result := processor.IND(suite.cpu)

	assert.Equal(suite.T(), uint16(0x5678), result.Address, "Expected address to be 0x5678")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestIND_PageBug() {
	// Write instruction "JMP $12FF" into memory at address 0x0000
	suite.bus.Write(0x0000, 0x6C) // opcode
	suite.bus.Write(0x0001, 0xFF) // lo byte
	suite.bus.Write(0x0002, 0x12) // hi byte

	// Write final target address 0x5678 into memory at address 0x12FF
	suite.bus.Write(0x12FF, 0x78) // lo byte
	suite.bus.Write(0x1200, 0x56) // hi byte (because of the page wrapping bug we write the correct hi bytes to 1200 not 1300)
	suite.bus.Write(0x1300, 0x99) // hi byte (write an incorrect hi byte at 0x12FF + 1 so we can clearly see if we haven't implemented the page wrapping bug)

	// Set the Program Counter to the start of the instruction
	suite.cpu.PC = 0x0000

	// Use indirect addressing mode to get the final address
	result := processor.IND(suite.cpu)

	if result.Address != 0x5678 {
		if result.Address == 0x9978 {
			suite.T().Errorf("Expected 0x5678, but got 0x%04X - page wrapping bug not implemented", result.Address)
		} else {
			suite.T().Errorf("Expected 0x5678, but got 0x%04X", result.Address)
		}
	}
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestINDX() {
	// Write instruction "LDA ($40,X)" into memory at address 0x2000
	suite.bus.Write(0x2000, 0xA1) // opcode
	suite.bus.Write(0x2001, 0x40) // operand

	// Write some data to the zero page
	suite.bus.Write(0x0050, 0x78)
	suite.bus.Write(0x0051, 0x12)

	// Set the Program Counter to the start of the instruction
	suite.cpu.PC = 0x2000
	suite.cpu.X = 0x10

	// Use indexed indirect X addressing mode to get the target address
	result := processor.INDX(suite.cpu)

	assert.Equal(suite.T(), uint16(0x1278), result.Address, "Expected address to be 0x1278")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestINDX_PageWrapping() {
	// Write instruction "LDA ($F5,X)" into memory at address 0x2000
	suite.bus.Write(0x2000, 0xA1) // opcode
	suite.bus.Write(0x2001, 0xF5) // operand

	// Write some data to the zero page
	suite.bus.Write(0x0015, 0x22)
	suite.bus.Write(0x0016, 0x33)

	// Set the Program Counter to the start of the instruction
	suite.cpu.PC = 0x2000
	suite.cpu.X = 0x20

	// Use indexed indirect X addressing mode to get the target address
	result := processor.INDX(suite.cpu)

	assert.Equal(suite.T(), uint16(0x3322), result.Address, "Expected address to be 0x3322")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestINDY() {
	// Write instruction "LDA ($40),Y" into memory at address 0x2000
	suite.bus.Write(0x2000, 0xB1) // opcode
	suite.bus.Write(0x2001, 0x40) // operand

	// Write some data to the zero page
	suite.bus.Write(0x0040, 0x78)
	suite.bus.Write(0x0041, 0x12)

	// Set the Program Counter to the start of the instruction
	suite.cpu.PC = 0x2000
	suite.cpu.Y = 0x0A

	// Use indirect indexed Y addressing mode to get the target address
	result := processor.INDY(suite.cpu)

	assert.Equal(suite.T(), uint16(0x1282), result.Address, "Expected address to be 0x1282")
	assert.False(suite.T(), result.PageChanged, "Expected pageChanged to be be false")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}

func (suite *AddressModesSuite) TestINDY_PageWrapping() {
	// Write instruction "LDA ($40),Y" into memory at address 0x2000
	suite.bus.Write(0x2000, 0xB1) // opcode
	suite.bus.Write(0x2001, 0x40) // operand

	// Write some data to the zero page
	suite.bus.Write(0x0040, 0x78)
	suite.bus.Write(0x0041, 0x12)

	// Set the Program Counter to the start of the instruction
	suite.cpu.PC = 0x2000
	suite.cpu.Y = 0xFF

	// Use indirect indexed Y addressing mode to get the target address
	result := processor.INDY(suite.cpu)

	assert.Equal(suite.T(), uint16(0x1377), result.Address, "Expected address to be 0x1377")
	assert.True(suite.T(), result.PageChanged, "Expected pageChanged to be true")
	assert.False(suite.T(), result.IsAccumulator, "Expected IsAccumulator to be false")
}
