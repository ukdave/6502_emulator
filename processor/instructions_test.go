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

func (suite *InstructionsSuite) TestADC() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x02)

	// Set the accumulator and carry flag to known values
	suite.cpu.A = 0x03
	suite.cpu.SetFlag(processor.C, false)

	// Execute ADC instruction
	extraCycle := processor.ADC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x05), suite.cpu.A, "Accumulator should be 0x05")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestADC_InitialCarrySet() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x02)

	// Set the accumulator and carry flag to known values
	suite.cpu.A = 0x03
	suite.cpu.SetFlag(processor.C, true)

	// Execute ADC instruction
	extraCycle := processor.ADC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x06), suite.cpu.A, "Accumulator should be 0x06")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestADC_CarryOut() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x01)

	// Set the accumulator and carry flag to known values
	suite.cpu.A = 0xFF
	suite.cpu.SetFlag(processor.C, false)

	// Execute ADC instruction
	extraCycle := processor.ADC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.A, "Accumulator should wrap to 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be set")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be set")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestADC_Overflow() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x01)

	// Set the accumulator and carry flag to known values
	suite.cpu.A = 0x7F // +127
	suite.cpu.SetFlag(processor.C, false)

	// Execute ADC instruction
	extraCycle := processor.ADC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x80), suite.cpu.A, "Accumulator should be 0x80")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be set")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be set")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestADC_NegativeOverflow() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x80) // -128

	// Set the accumulator and carry flag to known values
	suite.cpu.A = 0x80 // -128
	suite.cpu.SetFlag(processor.C, false)

	// Execute ADC instruction
	extraCycle := processor.ADC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.A, "Accumulator should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be set")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be set")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be set")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestSBC() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x02)

	// Set the accumulator and carry flag to known values
	suite.cpu.A = 0x05
	suite.cpu.SetFlag(processor.C, true)

	// Execute SBC instruction
	extraCycle := processor.SBC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x03), suite.cpu.A, "Accumulator should be 0x03")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestSBC_Borrow() {
	suite.bus.Write(0x2000, 0x05)

	suite.cpu.A = 0x02
	suite.cpu.SetFlag(processor.C, true)

	extraCycle := processor.SBC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0xFD), suite.cpu.A, "Accumulator should wrap to 0xFD")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be cleared (borrow)")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be set")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestSBC_ZeroResult() {
	suite.bus.Write(0x2000, 0x05)

	suite.cpu.A = 0x05
	suite.cpu.SetFlag(processor.C, true)

	extraCycle := processor.SBC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.A, "Accumulator should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should remain set")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be set")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestSBC_InitialCarryClear() {
	suite.bus.Write(0x2000, 0x01)

	suite.cpu.A = 0x02
	suite.cpu.SetFlag(processor.C, false)

	extraCycle := processor.SBC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.A, "Accumulator should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be set")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be set")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestSBC_Overflow() {
	suite.bus.Write(0x2000, 0xFF) // -1

	suite.cpu.A = 0x7F // +127
	suite.cpu.SetFlag(processor.C, true)

	extraCycle := processor.SBC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x80), suite.cpu.A, "Accumulator should be 0x80")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be cleared")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be set")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be set")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestINC() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x05)

	// Execute INC instruction
	extraCycle := processor.INC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x06), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0x06")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestINC_ZeroResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0xFF)

	// Execute INC instruction
	extraCycle := processor.INC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x00), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestINC_NegativeResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0xFE)

	// Execute INC instruction
	extraCycle := processor.INC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0xFF), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestDEC() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x05)

	// Execute DEC instruction
	extraCycle := processor.DEC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x04), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0x04")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestDEC_ZeroResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x01)

	// Execute DEC instruction
	extraCycle := processor.DEC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x00), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestDEC_NegativeResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x00)

	// Execute DEC instruction
	extraCycle := processor.DEC(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0xFF), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestINX() {
	// Set the X register to a known value
	suite.cpu.X = 0x05

	// Execute INX instruction
	extraCycle := processor.INX(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0x06), suite.cpu.X, "X should be 0x06")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestINX_ZeroResult() {
	// Set the X register to a known value
	suite.cpu.X = 0xFF

	// Execute INX instruction
	extraCycle := processor.INX(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.X, "X should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestINX_NegativeResult() {
	// Set the X register to a known value
	suite.cpu.X = 0xFE

	// Execute INX instruction
	extraCycle := processor.INX(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.X, "X should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestDEX() {
	// Set the X register to a known value
	suite.cpu.X = 0x05

	// Execute DEX instruction
	extraCycle := processor.DEX(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0x04), suite.cpu.X, "X should be 0x04")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestDEX_ZeroResult() {
	// Set the X register to a known value
	suite.cpu.X = 0x01

	// Execute DEX instruction
	extraCycle := processor.DEX(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.X, "X should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestDEX_NegativeResult() {
	// Set the X register to a known value
	suite.cpu.X = 0x00

	// Execute DEX instruction
	extraCycle := processor.DEX(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.X, "X should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestINY() {
	// Set the Y register to a known value
	suite.cpu.Y = 0x05

	// Execute INY instruction
	extraCycle := processor.INY(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0x06), suite.cpu.Y, "Y should be 0x06")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestINY_ZeroResult() {
	// Set the Y register to a known value
	suite.cpu.Y = 0xFF

	// Execute INY instruction
	extraCycle := processor.INY(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.Y, "Y should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestINY_NegativeResult() {
	// Set the Y register to a known value
	suite.cpu.Y = 0xFE

	// Execute INY instruction
	extraCycle := processor.INY(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.Y, "Y should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestDEY() {
	// Set the Y register to a known value
	suite.cpu.Y = 0x05

	// Execute DEX instruction
	extraCycle := processor.DEY(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0x04), suite.cpu.Y, "Y should be 0x04")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestDEY_ZeroResult() {
	// Set the Y register to a known value
	suite.cpu.Y = 0x01

	// Execute DEY instruction
	extraCycle := processor.DEY(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.Y, "Y should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestDEY_NegativeResult() {
	// Set the Y register to a known value
	suite.cpu.Y = 0x00

	// Execute DEY instruction
	extraCycle := processor.DEY(suite.cpu, processor.AddressInfo{Address: 0x00})

	assert.Equal(suite.T(), uint8(0xFF), suite.cpu.Y, "Y should be 0xFF")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

//
// Shift Instructions
//

func (suite *InstructionsSuite) TestASL() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000001)

	// Execute ASL instruction
	extraCycle := processor.ASL(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000010), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b00000010")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestASL_Carry() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b11000000)

	// Execute ASL instruction
	extraCycle := processor.ASL(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b10000000), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b10000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestASL_ZeroResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b10000000)

	// Execute ASL instruction
	extraCycle := processor.ASL(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000000), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b00000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestASL_NegativeResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b01100000)

	// Execute ASL instruction
	extraCycle := processor.ASL(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b11000000), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b11000000")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestASL_Accumulator() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b00000001

	// Execute ASL instruction in Accumulator mode
	extraCycle := processor.ASL(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b00000010), suite.cpu.A, "Accumulator should be 0b00000010")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestASL_AccumulatorAndCarry() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b11000000

	// Execute ASL instruction in Accumulator mode
	extraCycle := processor.ASL(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b10000000), suite.cpu.A, "Accumulator should be 0b10000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestASL_AccumulatorAndZeroResult() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b10000000

	// Execute ASL instruction in Accumulator mode
	extraCycle := processor.ASL(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b00000000), suite.cpu.A, "Accumulator should be 0b00000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestASL_AccumulatorAndNegativeResult() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b01100000

	// Execute ASL instruction in Accumulator mode
	extraCycle := processor.ASL(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b11000000), suite.cpu.A, "Accumulator should be 0b11000000")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestLSR() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000010)

	// Execute LSR instruction
	extraCycle := processor.LSR(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b0000001), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b0000001")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestLSR_Carry() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000011)

	// Execute LSR instruction
	extraCycle := processor.LSR(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000001), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b00000001")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestLSR_ZeroResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000001)

	// Execute LSR instruction
	extraCycle := processor.LSR(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000000), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b00000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestLSR_Accumulator() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b00000010

	// Execute LSR instruction in Accumulator mode
	extraCycle := processor.LSR(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b00000001), suite.cpu.A, "Accumulator should be 0b00000001")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestLSR_AccumulatorAndCarry() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b00000011

	// Execute LSR instruction in Accumulator mode
	extraCycle := processor.LSR(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b00000001), suite.cpu.A, "Accumulator should be 0b00000001")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestLSR_AccumulatorAndZeroResult() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b00000001

	// Execute LSR instruction in Accumulator mode
	extraCycle := processor.LSR(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b00000000), suite.cpu.A, "Accumulator should be 0b00000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROL() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000010)
	suite.cpu.SetFlag(processor.C, false)

	// Execute ROL instruction
	extraCycle := processor.ROL(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000100), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b00000100")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROL_Negative() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b01000000)
	suite.cpu.SetFlag(processor.C, false)

	// Execute ROL instruction
	extraCycle := processor.ROL(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b10000000), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b10000000")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROL_CarryOut() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b10000000)
	suite.cpu.SetFlag(processor.C, false)

	// Execute ROL instruction
	extraCycle := processor.ROL(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000000), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b00000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROL_CarryIn() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000000)
	suite.cpu.SetFlag(processor.C, true)

	// Execute ROL instruction
	extraCycle := processor.ROL(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000001), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b00000001")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROL_Accumulator() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b00000010
	suite.cpu.SetFlag(processor.C, false)

	// Execute ROL instruction
	extraCycle := processor.ROL(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b00000100), suite.cpu.A, "Accumulator should be 0b00000100")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROL_AccumulatorAndNegative() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b01000000
	suite.cpu.SetFlag(processor.C, false)

	// Execute ROL instruction
	extraCycle := processor.ROL(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b10000000), suite.cpu.A, "Accumulator should be 0b10000000")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROL_AccumulatorAndCarryOut() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b10000000
	suite.cpu.SetFlag(processor.C, false)

	// Execute ROL instruction
	extraCycle := processor.ROL(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b00000000), suite.cpu.A, "Accumulator should be 0b00000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROL_AccumulatorAndCarryIn() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b00000000
	suite.cpu.SetFlag(processor.C, true)

	// Execute ROL instruction
	extraCycle := processor.ROL(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b00000001), suite.cpu.A, "Accumulator should be 0b00000001")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROR() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000010)
	suite.cpu.SetFlag(processor.C, false)

	// Execute ROR instruction
	extraCycle := processor.ROR(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000001), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b00000001")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROR_CarryOut() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000001)
	suite.cpu.SetFlag(processor.C, false)

	// Execute ROR instruction
	extraCycle := processor.ROR(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000000), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b00000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROR_CarryIn() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000000)
	suite.cpu.SetFlag(processor.C, true)

	// Execute ROR instruction
	extraCycle := processor.ROR(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b10000000), suite.bus.Read(0x2000), "Memory at 0x2000 should be 0b10000000")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROR_Accumulator() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b00000010
	suite.cpu.SetFlag(processor.C, false)

	// Execute ROR instruction
	extraCycle := processor.ROR(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b00000001), suite.cpu.A, "Accumulator should be 0b00000001")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROR_AccumulatorAndCarryOut() {
	// Set the Accumulator to a known value
	suite.cpu.A = 0b00000001
	suite.cpu.SetFlag(processor.C, false)

	// Execute ROR instruction
	extraCycle := processor.ROR(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b00000000), suite.cpu.A, "Accumulator should be 0b00000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestROR_AccumulatorAndCarryIn() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000000)
	suite.cpu.SetFlag(processor.C, true)

	// Execute ROR instruction
	extraCycle := processor.ROR(suite.cpu, processor.AddressInfo{IsAccumulator: true})

	assert.Equal(suite.T(), uint8(0b10000000), suite.cpu.A, "Accumulator should be 0b10000000")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Carry flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

//
// Bitwise Instructions
//

func (suite *InstructionsSuite) TestAND() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00001111)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b01010101

	// Execute AND instruction
	extraCycle := processor.AND(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000101), suite.cpu.A, "Accumulator should be 0b00000101")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestAND_ZeroResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00001111)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b11110000

	// Execute AND instruction
	extraCycle := processor.AND(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000000), suite.cpu.A, "Accumulator should be 0b00000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestAND_NegativeResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b11110000)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b10101010

	// Execute AND instruction
	extraCycle := processor.AND(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b10100000), suite.cpu.A, "Accumulator should be 0b10100000")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestORA() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00001111)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b01010101

	// Execute ORA instruction
	extraCycle := processor.ORA(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b01011111), suite.cpu.A, "Accumulator should be 0b01011111")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestORA_ZeroResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x00)

	// Set the Accumulator to a known value
	suite.cpu.A = 0x00

	// Execute ORA instruction
	extraCycle := processor.ORA(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0x00), suite.cpu.A, "Accumulator should be 0x00")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestORA_NegativeResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b10101010)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b01010101

	// Execute ORA instruction
	extraCycle := processor.ORA(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b11111111), suite.cpu.A, "Accumulator should be 0b11111111")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestEOR() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00101000)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b00101101

	// Execute EOR instruction
	extraCycle := processor.EOR(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000101), suite.cpu.A, "Accumulator should be 0b00000101")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestEOR_ZeroResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b01010101)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b01010101

	// Execute EOR instruction
	extraCycle := processor.EOR(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b00000000), suite.cpu.A, "Accumulator should be 0b00000000")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestEOR_NegativeResult() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b10101010)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b01010101

	// Execute EOR instruction
	extraCycle := processor.EOR(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint8(0b11111111), suite.cpu.A, "Accumulator should be 0b11111111")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestBIT_Zero() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b00000010)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b00001111

	// Execute BIT instruction
	extraCycle := processor.BIT(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBIT_Negative() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b10000000)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b00000001

	// Execute BIT instruction
	extraCycle := processor.BIT(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBIT_Overflow() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0b01000000)

	// Set the Accumulator to a known value
	suite.cpu.A = 0b00000001

	// Execute BIT instruction
	extraCycle := processor.BIT(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.V), "Overflow flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

//
// Compare Instructions
//

func (suite *InstructionsSuite) TestCMP_LessThan() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x02)

	// Set the Accumulator to a known value
	suite.cpu.A = 0x03

	// Execute CMP instruction
	extraCycle := processor.CMP(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Cary flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestCMP_GreaterThan() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x03)

	// Set the Accumulator to a known value
	suite.cpu.A = 0x02

	// Execute CMP instruction
	extraCycle := processor.CMP(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Cary flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestCMP_Equal() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x02)

	// Set the Accumulator to a known value
	suite.cpu.A = 0x02

	// Execute CMP instruction
	extraCycle := processor.CMP(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Cary flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.True(suite.T(), extraCycle, "Expected extraCycle to be true")
}

func (suite *InstructionsSuite) TestCPX_LessThan() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x02)

	// Set the X Register to a known value
	suite.cpu.X = 0x03

	// Execute CPX instruction
	extraCycle := processor.CPX(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Cary flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestCPX_GreaterThan() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x03)

	// Set the X Register to a known value
	suite.cpu.X = 0x02

	// Execute CPX instruction
	extraCycle := processor.CPX(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Cary flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestCPX_Equal() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x02)

	// Set the X Register to a known value
	suite.cpu.X = 0x02

	// Execute CPX instruction
	extraCycle := processor.CPX(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Cary flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestCPY_LessThan() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x02)

	// Set the Y Register to a known value
	suite.cpu.Y = 0x03

	// Execute CPY instruction
	extraCycle := processor.CPY(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Cary flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestCPY_GreaterThan() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x03)

	// Set the Y Register to a known value
	suite.cpu.Y = 0x02

	// Execute CPY instruction
	extraCycle := processor.CPY(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.False(suite.T(), suite.cpu.GetFlag(processor.C), "Cary flag should be false")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be false")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be true")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestCPY_Equal() {
	// Write a value to memory at address 0x2000
	suite.bus.Write(0x2000, 0x02)

	// Set the Y Register to a known value
	suite.cpu.Y = 0x02

	// Execute CPY instruction
	extraCycle := processor.CPY(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.True(suite.T(), suite.cpu.GetFlag(processor.C), "Cary flag should be true")
	assert.True(suite.T(), suite.cpu.GetFlag(processor.Z), "Zero flag should be true")
	assert.False(suite.T(), suite.cpu.GetFlag(processor.N), "Negative flag should be false")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

//
// Branch Instructions
//

func (suite *InstructionsSuite) TestBCC_Taken() {
	// Set the Carry flag to false
	suite.cpu.SetFlag(processor.C, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BCC instruction
	extraCycle := processor.BCC(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1050), suite.cpu.PC, "Expected PC to be 0x1050")
	assert.Equal(suite.T(), uint8(1), suite.cpu.Cycles(), "Expected cycles to be 1")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBCC_TakenAndPageCrossed() {
	// Set the Carry flag to false
	suite.cpu.SetFlag(processor.C, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BCC instruction
	extraCycle := processor.BCC(suite.cpu, processor.AddressInfo{Address: 0x2000, PageChanged: true})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(suite.T(), uint8(2), suite.cpu.Cycles(), "Expected cycles to be 2")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBCC_NotTaken() {
	// Set the Carry flag to true
	suite.cpu.SetFlag(processor.C, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BCC instruction
	extraCycle := processor.BCC(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1000), suite.cpu.PC, "Expected PC to be 0x1000")
	assert.Equal(suite.T(), uint8(0), suite.cpu.Cycles(), "Expected cycles to be 0")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBCS_Taken() {
	// Set the Carry flag to true
	suite.cpu.SetFlag(processor.C, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BCS instruction
	extraCycle := processor.BCS(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1050), suite.cpu.PC, "Expected PC to be 0x1050")
	assert.Equal(suite.T(), uint8(1), suite.cpu.Cycles(), "Expected cycles to be 1")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBCS_TakenAndPageCrossed() {
	// Set the Carry flag to true
	suite.cpu.SetFlag(processor.C, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BCS instruction
	extraCycle := processor.BCS(suite.cpu, processor.AddressInfo{Address: 0x2000, PageChanged: true})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(suite.T(), uint8(2), suite.cpu.Cycles(), "Expected cycles to be 2")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBCS_NotTaken() {
	// Set the Carry flag to false
	suite.cpu.SetFlag(processor.C, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BCS instruction
	extraCycle := processor.BCS(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1000), suite.cpu.PC, "Expected PC to be 0x1000")
	assert.Equal(suite.T(), uint8(0), suite.cpu.Cycles(), "Expected cycles to be 0")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBEQ_Taken() {
	// Set the Zero flag to true
	suite.cpu.SetFlag(processor.Z, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BEQ instruction
	extraCycle := processor.BEQ(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1050), suite.cpu.PC, "Expected PC to be 0x1050")
	assert.Equal(suite.T(), uint8(1), suite.cpu.Cycles(), "Expected cycles to be 1")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBEQ_TakenAndPageCrossed() {
	// Set the Zero flag to true
	suite.cpu.SetFlag(processor.Z, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BEQ instruction
	extraCycle := processor.BEQ(suite.cpu, processor.AddressInfo{Address: 0x2000, PageChanged: true})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(suite.T(), uint8(2), suite.cpu.Cycles(), "Expected cycles to be 2")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBEQ_NotTaken() {
	// Set the Zero flag to false
	suite.cpu.SetFlag(processor.Z, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BEQ instruction
	extraCycle := processor.BEQ(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1000), suite.cpu.PC, "Expected PC to be 0x1000")
	assert.Equal(suite.T(), uint8(0), suite.cpu.Cycles(), "Expected cycles to be 0")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBNE_Taken() {
	// Set the Zero flag to false
	suite.cpu.SetFlag(processor.Z, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BNE instruction
	extraCycle := processor.BNE(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1050), suite.cpu.PC, "Expected PC to be 0x1050")
	assert.Equal(suite.T(), uint8(1), suite.cpu.Cycles(), "Expected cycles to be 1")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBNE_TakenAndPageCrossed() {
	// Set the Zero flag to false
	suite.cpu.SetFlag(processor.Z, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BNE instruction
	extraCycle := processor.BNE(suite.cpu, processor.AddressInfo{Address: 0x2000, PageChanged: true})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(suite.T(), uint8(2), suite.cpu.Cycles(), "Expected cycles to be 2")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBNE_NotTaken() {
	// Set the Zero flag to true
	suite.cpu.SetFlag(processor.Z, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BNE instruction
	extraCycle := processor.BNE(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1000), suite.cpu.PC, "Expected PC to be 0x1000")
	assert.Equal(suite.T(), uint8(0), suite.cpu.Cycles(), "Expected cycles to be 0")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBPL_Taken() {
	// Set the Negative flag to false
	suite.cpu.SetFlag(processor.N, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BPL instruction
	extraCycle := processor.BPL(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1050), suite.cpu.PC, "Expected PC to be 0x1050")
	assert.Equal(suite.T(), uint8(1), suite.cpu.Cycles(), "Expected cycles to be 1")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBPL_TakenAndPageCrossed() {
	// Set the Negative flag to false
	suite.cpu.SetFlag(processor.N, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BPL instruction
	extraCycle := processor.BPL(suite.cpu, processor.AddressInfo{Address: 0x2000, PageChanged: true})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(suite.T(), uint8(2), suite.cpu.Cycles(), "Expected cycles to be 2")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBPL_NotTaken() {
	// Set the Negative flag to true
	suite.cpu.SetFlag(processor.N, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BPL instruction
	extraCycle := processor.BPL(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1000), suite.cpu.PC, "Expected PC to be 0x1000")
	assert.Equal(suite.T(), uint8(0), suite.cpu.Cycles(), "Expected cycles to be 0")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBMI_Taken() {
	// Set the Negative flag to true
	suite.cpu.SetFlag(processor.N, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BMI instruction
	extraCycle := processor.BMI(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1050), suite.cpu.PC, "Expected PC to be 0x1050")
	assert.Equal(suite.T(), uint8(1), suite.cpu.Cycles(), "Expected cycles to be 1")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBMI_TakenAndPageCrossed() {
	// Set the Negative flag to true
	suite.cpu.SetFlag(processor.N, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BMI instruction
	extraCycle := processor.BMI(suite.cpu, processor.AddressInfo{Address: 0x2000, PageChanged: true})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(suite.T(), uint8(2), suite.cpu.Cycles(), "Expected cycles to be 2")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBMI_NotTaken() {
	// Set the Negative flag to false
	suite.cpu.SetFlag(processor.N, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BMI instruction
	extraCycle := processor.BMI(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1000), suite.cpu.PC, "Expected PC to be 0x1000")
	assert.Equal(suite.T(), uint8(0), suite.cpu.Cycles(), "Expected cycles to be 0")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBVC_Taken() {
	// Set the Overflow flag to false
	suite.cpu.SetFlag(processor.V, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BVC instruction
	extraCycle := processor.BVC(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1050), suite.cpu.PC, "Expected PC to be 0x1050")
	assert.Equal(suite.T(), uint8(1), suite.cpu.Cycles(), "Expected cycles to be 1")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBVC_TakenAndPageCrossed() {
	// Set the Overflow flag to false
	suite.cpu.SetFlag(processor.V, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BVC instruction
	extraCycle := processor.BVC(suite.cpu, processor.AddressInfo{Address: 0x2000, PageChanged: true})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(suite.T(), uint8(2), suite.cpu.Cycles(), "Expected cycles to be 2")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBVC_NotTaken() {
	// Set the Overflow flag to true
	suite.cpu.SetFlag(processor.V, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BVC instruction
	extraCycle := processor.BVC(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1000), suite.cpu.PC, "Expected PC to be 0x1000")
	assert.Equal(suite.T(), uint8(0), suite.cpu.Cycles(), "Expected cycles to be 0")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBVS_Taken() {
	// Set the Overflow flag to true
	suite.cpu.SetFlag(processor.V, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BVS instruction
	extraCycle := processor.BVS(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1050), suite.cpu.PC, "Expected PC to be 0x1050")
	assert.Equal(suite.T(), uint8(1), suite.cpu.Cycles(), "Expected cycles to be 1")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBVS_TakenAndPageCrossed() {
	// Set the Overflow flag to true
	suite.cpu.SetFlag(processor.V, true)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BVS instruction
	extraCycle := processor.BVS(suite.cpu, processor.AddressInfo{Address: 0x2000, PageChanged: true})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(suite.T(), uint8(2), suite.cpu.Cycles(), "Expected cycles to be 2")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBVS_NotTaken() {
	// Set the Overflow flag to false
	suite.cpu.SetFlag(processor.V, false)

	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute BVS instruction
	extraCycle := processor.BVS(suite.cpu, processor.AddressInfo{Address: 0x1050})

	assert.Equal(suite.T(), uint16(0x1000), suite.cpu.PC, "Expected PC to be 0x1000")
	assert.Equal(suite.T(), uint8(0), suite.cpu.Cycles(), "Expected cycles to be 0")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

//
// Jump Instructions
//

func (suite *InstructionsSuite) TestJMP() {
	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1000

	// Execute JMP instruction
	extraCycle := processor.JMP(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestJSR() {
	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1005

	// Execute JSR instruction
	extraCycle := processor.JSR(suite.cpu, processor.AddressInfo{Address: 0x2000})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(suite.T(), uint8(0xFB), suite.cpu.SP, "Expected stack pointer to be 0xFB")
	assert.Equal(suite.T(), uint16(0x1004), suite.cpu.Pop16(), "Expected stack to contain 0x1004")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestRTS() {
	// Set the Program Counter to a known value
	suite.cpu.PC = 0x2000

	// Push an address onto the stack
	suite.cpu.Push16(0x1004)

	// Execute RTS instruction
	extraCycle := processor.RTS(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint16(0x1005), suite.cpu.PC, "Expected PC to be 0x1005")
	assert.Equal(suite.T(), uint8(0xFD), suite.cpu.SP, "Expected stack pointer to be 0xFD")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestBRK() {
	// Set the Program Counter to a known value
	suite.cpu.PC = 0x2000

	// Write a 16-bit value to memory to 0xFFFE
	suite.cpu.Write16(0xFFFE, 0x1005)

	// Read current status flags
	statusFlags := suite.cpu.Status

	// Execute BRK instruction
	extraCycle := processor.BRK(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint16(0x1005), suite.cpu.PC, "Expected PC to be 0x1005")
	assert.Equal(suite.T(), true, suite.cpu.GetFlag(processor.I), "Disable Interrupt flag should be true")
	assert.Equal(suite.T(), uint8(0xFA), suite.cpu.SP, "Expected stack pointer to be 0xFA (3 bytes pushed)")
	assert.Equal(suite.T(), statusFlags|0x10, suite.cpu.Pop(), "Expected stack to contain status flags with B set")
	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.Pop16(), "Expected stack contain original PC value (0x2000)")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

func (suite *InstructionsSuite) TestRTI() {
	// Set the Program Counter to a known value
	suite.cpu.PC = 0x1005

	// Push an address and status flags onto the stack
	suite.cpu.Push16(0x2000)
	suite.cpu.Push(0xB7)

	// Execute RTI instruction
	extraCycle := processor.RTI(suite.cpu, processor.AddressInfo{})

	assert.Equal(suite.T(), uint16(0x2000), suite.cpu.PC, "Expected PC to be 0x2000")
	assert.Equal(suite.T(), false, suite.cpu.GetFlag(processor.B), "Break flag should be false")
	assert.Equal(suite.T(), true, suite.cpu.GetFlag(processor.U), "Unused flag should be true")
	assert.Equal(suite.T(), uint8(0xA7), suite.cpu.Status, "Expected status flags to be 0xA7")
	assert.Equal(suite.T(), uint8(0xFD), suite.cpu.SP, "Expected stack pointer to be 0xFD")
	assert.False(suite.T(), extraCycle, "Expected extraCycle to be false")
}

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
