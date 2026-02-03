package processor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ukdave/6502_emulator/bus"
	"github.com/ukdave/6502_emulator/processor"
)

// TestIntegration is a simple integration test that loads a basic program into memory and executes it.
// Only a few instructions and address modes are tested here, but combined with the unit tests it's good
// enough for now. For a full test we'd probably want to run something like nestest.
// https://www.nesdev.org/wiki/Emulator_tests
func TestIntegration(t *testing.T) {
	// Create a new bus
	bus := bus.NewSimpleBus()

	// Set the value of the reset vector to 0x8000. This is where our program will start
	bus.Write(0xFFFC, 0x00)
	bus.Write(0xFFFD, 0x80)

	// This program will multiply 10 (0x0A) by 3 (0x03) using repeated addition and store the result at 0x0002
	bytes := []byte{
		0xA2, 0x0A, //         LDX #$0A {IMM}
		0x8E, 0x00, 0x00, //   STX $0000 {ABS}
		0xA2, 0x03, //         LDX #$03 {IMM}
		0x8E, 0x01, 0x00, //   STX $0001 {ABS}
		0xAC, 0x00, 0x00, //   LDY $0000 {ABS}
		0xA9, 0x00, //         LDA #$00 {IMM}
		0x18,             //   CLC {IMP}
		0x6D, 0x01, 0x00, //   ADC $0001 {ABS}
		0x88,       //         DEY {IMP}
		0xD0, 0xFA, //         BNE $FA [$8010] {REL}
		0x8D, 0x02, 0x00, //   STA $0002 {ABS}
		0xEA, //               NOP {IMP}
		0xEA, //               NOP {IMP}
		0xEA, //               NOP {IMP}
	}
	for i, b := range bytes {
		bus.Write(0x8000+uint16(i), b)
	}

	// Create new CPU
	cpu := processor.NewCPU(bus)

	// Clock the CPU until the Program Counter equals 0x0000 indicating that our program has finished.
	// This works because the memory is zeroed-out when it is initialised. This means that the next
	// instruction after the end of our program will be interpreted as a BRK (interrupt). This will cause
	// the Program Counter to be set to the memory address stored in the IRQ Vector (0xFFFE) which will
	// be 0x0000. In case of a bug in the emulator we will also stop if we clock the CPU more than 1,000 times.
	totalClockCycles := 0
	clockCycleLimit := 1000
	for {
		// Step through one instruction
		for {
			cpu.Clock()
			totalClockCycles++
			if cpu.Cycles() == 0 || totalClockCycles > clockCycleLimit {
				break
			}
		}
		// Break if Program Counter == 0x0000
		if cpu.PC == 0x0000 || totalClockCycles > clockCycleLimit {
			break
		}
	}

	assert.Equal(t, uint16(0x0000), cpu.PC, "Expected Program Counter to be 0x0000")
	assert.Equal(t, 126, totalClockCycles, "Expected totalClockCycles to be 126")
	assert.Equal(t, uint8(0x1E), bus.Read(0x0002), "Expected program result (at 0x0002) to be 0x1E")
}
