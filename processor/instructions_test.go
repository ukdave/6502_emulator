package processor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ukdave/6502_emulator/bus"
	"github.com/ukdave/6502_emulator/processor"
)

type InstructionsSuite struct {
	suite.Suite
	bus bus.Bus
	cpu *processor.CPU
}

func TestInstructionsSuite(t *testing.T) {
	suite.Run(t, new(InstructionsSuite))
}

func (suite *InstructionsSuite) SetupTest() {
	suite.bus = bus.NewSimpleBus()
	suite.cpu = processor.NewCPU(suite.bus)
}

//
// Access Instructions
//

func (suite *InstructionsSuite) TestLDA() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x12)

	// Set the Accumulator to a known value
	suite.cpu.A = 0x00

	// Execute LDA instruction
	extraCycle := processor.LDA(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x12), suite.cpu.A, "Accumulator should be 0x12")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestLDA_ZeroValue() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x00)

	// Set the Accumulator to a known value
	suite.cpu.A = 0x12

	// Execute LDA instruction
	extraCycle := processor.LDA(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.A, "Accumulator should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestLDA_NegativeValue() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0xFF)

	// Set the Accumulator to a known value
	suite.cpu.A = 0x00

	// Execute LDA instruction
	extraCycle := processor.LDA(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.A, "Accumulator should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestSTA() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0x34

	// Execute STA instruction
	extraCycle := processor.STA(suite.cpu, processor.AddressInfo{Address: 0x2000})

	// Read back the value from memory
	value := suite.bus.Read(0x2000)

	assert.Equal(suite.T(), uint8(0x34), value, "Memory at 0x2000 should be 0x34")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestLDX() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x12)

	// Set the X Register to a known value
	suite.cpu.X = 0x00

	// Execute LDX instruction
	extraCycle := processor.LDX(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x12), suite.cpu.X, "X Register should be 0x12")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestLDX_ZeroValue() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x00)

	// Set the X Register to a known value
	suite.cpu.X = 0x12

	// Execute LDX instruction
	extraCycle := processor.LDX(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.X, "X Register should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestLDX_NegativeValue() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0xFF)

	// Set the X Register to a known value
	suite.cpu.X = 0x00

	// Execute LDX instruction
	extraCycle := processor.LDX(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.X, "X Register should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestSTX() {
	// Set the X Register to a known value
	suite.cpu.X = 0x34

	// Execute STX instruction
	extraCycle := processor.STX(suite.cpu, processor.AddressInfo{Address: 0x2000})

	// Read back the value from memory
	value := suite.bus.Read(0x2000)

	assert.Equal(suite.T(), uint8(0x34), value, "Memory at 0x2000 should be 0x34")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestLDY() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x12)

	// Set the Y Register to a known value
	suite.cpu.Y = 0x00

	// Execute LDY instruction
	extraCycle := processor.LDY(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x12), suite.cpu.Y, "Y Register should be 0x12")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestLDY_ZeroValue() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x00)

	// Set the Y Register to a known value
	suite.cpu.Y = 0x12

	// Execute LDY instruction
	extraCycle := processor.LDY(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.Y, "Y Register should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestLDY_NegativeValue() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0xFF)

	// Set the Y Register to a known value
	suite.cpu.Y = 0x00

	// Execute LDY instruction
	extraCycle := processor.LDY(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.Y, "Y Register should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestSTY() {
	// Set the Y Register to a known value
	suite.cpu.Y = 0x34

	// Execute STY instruction
	extraCycle := processor.STY(suite.cpu, processor.AddressInfo{Address: 0x2000})

	// Read back the value from memory
	value := suite.bus.Read(0x2000)

	assert.Equal(suite.T(), uint8(0x34), value, "Memory at 0x2000 should be 0x34")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

//
// Transfer Instructions
//

func (suite *InstructionsSuite) TestTAX() {
	// Set the Accumulator and X Register to known values
	suite.cpu.A = 0x56
	suite.cpu.X = 0x00

	// Execute TAX instruction
	extraCycle := processor.TAX(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0x56), suite.cpu.X, "X Register should be 0x56")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTAX_ZeroValue() {
	// Set the Accumulator and X Register to known values
	suite.cpu.A = 0x00
	suite.cpu.X = 0x12

	// Execute TAX instruction
	extraCycle := processor.TAX(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.X, "X Register should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTAX_NegativeValue() {
	// Set the Accumulator and X Register to known values
	suite.cpu.A = 0xFF
	suite.cpu.X = 0x00

	// Execute TAX instruction
	extraCycle := processor.TAX(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.X, "X Register should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTAY() {
	// Set the Accumulator and Y Register to known values
	suite.cpu.A = 0x56
	suite.cpu.Y = 0x00

	// Execute TAY instruction
	extraCycle := processor.TAY(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0x56), suite.cpu.Y, "Y Register should be 0x56")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTAY_ZeroValue() {
	// Set the Accumulator and Y Register to known values
	suite.cpu.A = 0x00
	suite.cpu.Y = 0x12

	// Execute TAY instruction
	extraCycle := processor.TAY(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.Y, "Y Register should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTAY_NegativeValue() {
	// Set the Accumulator and Y Register to known values
	suite.cpu.A = 0xFF
	suite.cpu.Y = 0x00

	// Execute TAY instruction
	extraCycle := processor.TAY(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.Y, "Y Register should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTXA() {
	// Set the X Register and Accumulator to known values
	suite.cpu.X = 0x56
	suite.cpu.A = 0x00

	// Execute TXA instruction
	extraCycle := processor.TXA(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0x56), suite.cpu.A, "Accumulator should be 0x56")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTXA_ZeroValue() {
	// Set the X Register and Accumulator to known values
	suite.cpu.X = 0x00
	suite.cpu.A = 0x12

	// Execute TXA instruction
	extraCycle := processor.TXA(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.A, "Accumulator should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTXA_NegativeValue() {
	// Set the X Register and Accumulator to known values
	suite.cpu.X = 0xFF
	suite.cpu.A = 0x00

	// Execute TXA instruction
	extraCycle := processor.TXA(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.A, "Accumulator should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTYA() {
	// Set the Y Register and Accumulator to known values
	suite.cpu.Y = 0x56
	suite.cpu.A = 0x00

	// Execute TYA instruction
	extraCycle := processor.TYA(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0x56), suite.cpu.A, "Accumulator should be 0x56")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTYA_ZeroValue() {
	// Set the Y Register and Accumulator to known values
	suite.cpu.Y = 0x00
	suite.cpu.A = 0x12

	// Execute TYA instruction
	extraCycle := processor.TYA(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.A, "Accumulator should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestTYA_NegativeValue() {
	// Set the Y Register and Accumulator to known values
	suite.cpu.Y = 0xFF
	suite.cpu.A = 0x00

	// Execute TYA instruction
	extraCycle := processor.TYA(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.A, "Accumulator should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

//
// Arithmetic Instructions
//

//
// Shift Instructions
//

//
// Bitwise Instructions
//

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

func (suite *InstructionsSuite) TestNOP() {
	// Execute NOP instruction
	extraCycle := processor.NOP(suite.cpu, processor.AddressInfo{})

	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestXXX() {
	// Execute illegal instruction
	extraCycle := processor.XXX(suite.cpu, processor.AddressInfo{})

	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}
