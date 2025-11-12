package bus_test

import (
	"testing"

	"github.com/ukdave/6502_emulator/bus"
)

func TestSimpleBus(t *testing.T) {
	// Create a new bus
	bus := bus.NewSimpleBus()

	// Define test data and address
	testCases := []struct {
		addr uint16
		data byte
	}{
		{0x0000, 0x01},
		{0x1000, 0x42},
		{0xFFFF, 0x99},
	}

	// Loop over test cases
	for _, tc := range testCases {
		// Write the data to the bus
		bus.Write(tc.addr, tc.data)

		// Read the data back from the bus
		readData := bus.Read(tc.addr)

		// Check that the data matches what we wrote
		if readData != tc.data {
			t.Errorf("At address 0x%X, expected %v but got %v", tc.addr, tc.data, readData)
		}
	}
}
