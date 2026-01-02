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
