package bus

// SimpleBus is a minimal concrete implementation of the Bus interface.
//
// It models a flat 64KB address space backed entirely by RAM, which is sufficient for basic CPU emulation
// and testing. In a more complete emulator, the bus would be responsible for routing reads and writes to
// different devices (RAM, ROM, PPU, I/O registers, etc.) based on the address.
type SimpleBus struct {
	ram [64 * 1024]byte // Fake RAM (64KB)
}

// NewSimpleBus creates a new SimpleBus instance.
//
// The RAM is zero-initialized by default. No memory mapping or device configuration is performed here.
func NewSimpleBus() *SimpleBus {
	return &SimpleBus{}
}

// Write stores a single byte at the given 16-bit address.
//
// In this simple implementation, the address maps directly to RAM. Future bus implementations may intercept
// or redirect writes to memory-mapped devices instead.
func (b *SimpleBus) Write(addr uint16, data byte) {
	b.ram[addr] = data
}

// Read returns the byte stored at the given 16-bit address.
//
// As with Write, this currently performs a direct RAM access. More advanced buses may return data from ROM
// or peripheral devices depending on the address range.
func (b *SimpleBus) Read(addr uint16) byte {
	return b.ram[addr]
}
