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
