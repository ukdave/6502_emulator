package main

import (
	"fmt"

	"github.com/ukdave/6502_emulator/bus"
	"github.com/ukdave/6502_emulator/processor"
)

func main() {
	// Create a new bus
	bus := bus.NewSimpleBus()

	// Set the value of the reset vector to 0x8000. This is where our program will start
	bus.Write(0xFFFC, 0x00)
	bus.Write(0xFFFD, 0x80)

	// Write a very simple program to memory starting at 0x8000.
	// The program consists of 4 instructions that does 5+3 and stores the answer at 0x10.
	bus.Write(0x8000, 0xA9) // LDA #$05
	bus.Write(0x8001, 0x05)
	bus.Write(0x8002, 0x18) // CLC
	bus.Write(0x8003, 0x69) // ADC #$03
	bus.Write(0x8004, 0x03)
	bus.Write(0x8005, 0x85)
	bus.Write(0x8006, 0x10) // STA $10

	// Create a new CPU
	cpu := processor.NewCPU(bus)
	for range 4 {
		printInstruction(cpu)
		execute(cpu)
	}
	printInstruction(cpu)
	fmt.Printf("0x10 = %02X", bus.Read(0x10))
}

func execute(cpu *processor.CPU) {
	for {
		cpu.Clock()
		if cpu.Cycles() == 0 {
			break
		}
	}
}

func printInstruction(cpu *processor.CPU) {
	opcode := cpu.Read(cpu.PC)
	op := cpu.GetOperation(opcode)
	b0 := fmt.Sprintf("%02X", cpu.Read(cpu.PC+0))
	b1 := fmt.Sprintf("%02X", cpu.Read(cpu.PC+1))
	b2 := fmt.Sprintf("%02X", cpu.Read(cpu.PC+2))
	if op.Size < 2 {
		b1 = "  "
	}
	if op.Size < 3 {
		b2 = "  "
	}
	fmt.Printf("%4X  %s %s %s %27s A:%02X X:%02X Y:%02X SP:%02X Flags:%08b\n",
		cpu.PC, b0, b1, b2, "", cpu.A, cpu.X, cpu.Y, cpu.SP, cpu.Status)
}
