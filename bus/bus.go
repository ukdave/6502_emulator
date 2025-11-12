// Package bus defines the interface for a system bus used to read and write data.
// A Bus provides methods for transferring bytes between components in a system,
// such as a CPU, memory, or peripheral devices.
package bus

// Bus interface defines the methods that any bus implementation must provide.
type Bus interface {
	Write(addr uint16, data byte)
	Read(addr uint16) byte
}
