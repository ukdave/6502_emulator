package processor

// The 6502 has a 16-bit address space ranging from 0x0000 to 0xFFFF. The upper byte of an address is commonly
// referred to as the "page", while the lower byte represents the offset within that page. This divides memory
// into 256 pages of 256 bytes each.
//
// Some addressing modes may incur an extra clock cycle when a memory access crosses a page boundary. Certain
// instructions allow this additional cycle to be taken. To model this behaviour, both the addressing mode and
// the instruction report whether an extra cycle is possible. If both indicate true, one additional clock cycle
// is added.

type AddressModeFunc func(*CPU) AddressInfo

type AddressInfo struct {
	Address       uint16
	PageChanged   bool
	IsAccumulator bool
}

// ACC implements "Accumulator" address mode.
// The operation is performed directly on the accumulator register (A) rather than on a memory location.
// Instructions using this addressing mode are 1 byte long.
func ACC(cpu *CPU) AddressInfo {
	return AddressInfo{IsAccumulator: true}
}

// IMM implements "Immediate" address mode.
// The operand is the byte immediately following the opcode and is treated as a literal value rather than a
// memory address.
func IMM(cpu *CPU) AddressInfo {
	return AddressInfo{Address: cpu.PC + 1}
}

// ABS implements "Absolute" address mode.
// A full 16-bit address is read from the instruction operands and used directly.
func ABS(cpu *CPU) AddressInfo {
	return AddressInfo{Address: cpu.Read16(cpu.PC + 1)}
}

// ABX implements "Absolute with X Offset" address mode.
// Similar to absolute addressing where a full 16-bit address is read from the instruction operands, except the
// value of the X Register is then added to form the effective address. Some instructions require an additional
// clock cycle if this addition causes a page boundary to be crossed.
func ABX(cpu *CPU) AddressInfo {
	addr := cpu.Read16(cpu.PC + 1)
	addr += uint16(cpu.X)
	pageChanged := pagesDiffer(addr-uint16(cpu.X), addr)
	return AddressInfo{Address: addr, PageChanged: pageChanged}
}

// ABY implements "Absolute with Y Offset" address mode.
// Same as ABX but using the Y Register instead. A full 16-bit address is read from the instruction operands and
// then the value of the Y Resister is added to form the effective address. Some instructions require an additional
// clock cycle if this addition causes a page boundary to be crossed.
func ABY(cpu *CPU) AddressInfo {
	addr := cpu.Read16(cpu.PC + 1)
	addr += uint16(cpu.Y)
	pageChanged := pagesDiffer(addr-uint16(cpu.Y), addr)
	return AddressInfo{Address: addr, PageChanged: pageChanged}
}

// ZP0 implements "Zero Page" address mode.
// The operand is an 8-bit address that implicitly refers to a location within page zero (0x0000–0x00FF). This
// addressing mode saves program bytes by only requiring one byte instead of two.
func ZP0(cpu *CPU) AddressInfo {
	return AddressInfo{Address: uint16(cpu.Read(cpu.PC + 1))}
}

// ZPX implements "Zero Page with X Offset" address mode.
// Similar to zero page addressing where an 8-bit address is read from the instruction operand, except the
// value of the X Register is then added to form the effective address. Any wrapping of the result occurs within
// page zero so the final address will always be in the range 0x0000–0x00FF.
func ZPX(cpu *CPU) AddressInfo {
	addr := uint16(cpu.Read(cpu.PC+1)+cpu.X) & 0x00FF
	return AddressInfo{Address: addr}
}

// ZPY implements "Zero Page with Y Offset" address mode.
// Same as ZPX but using the Y Register instead. An 8-bit address is read from the instruction operand and then
// the value of the Y Resister is added to form the effective address. Any wrapping of the result occurs within
// page zero so the final address will always be in the range 0x0000–0x00FF.
func ZPY(cpu *CPU) AddressInfo {
	addr := uint16(cpu.Read(cpu.PC+1)+cpu.Y) & 0x00FF
	return AddressInfo{Address: addr}
}

// IMP implements "Implied" address mode.
// The instruction implicitly operates on internal CPU state and or registers and does not reference a memory
// address. Instructions using this addressing mode are 1 byte long.
func IMP(cpu *CPU) AddressInfo {
	return AddressInfo{}
}

// REL implements "Relative" address mode.
// The operand is a signed 8-bit offset relative to the current program counter and is used exclusively by branch
// instructions. The final address will be in the range -128 to +127 of the program counter. This means it is not
// possible to branch to any address in the full address space. If a page boundary is crossed then two additional
// clock cycles will be required, but only if the branch is taken.
func REL(cpu *CPU) AddressInfo {
	offset := cpu.Read(cpu.PC + 1)
	baseAddr := cpu.PC + 2
	addr := baseAddr + uint16(offset)
	if offset >= 0x80 {
		addr -= 0x100
	}
	pageChanged := pagesDiffer(baseAddr, addr)
	return AddressInfo{Address: addr, PageChanged: pageChanged}
}

// Note: The next 3 address modes use indirection (aka Pointers!)

// IND implements "Absolute Indirect" address mode.
// A 16-bit pointer is read from the instruction operands and used to fetch the final address. Due to a hardware
// bug in the original 6502, if the pointer’s low byte is 0xFF, the high byte of the target address is read from
// the beginning of the same page instead of the next page (i.e. the address wraps within the page).
func IND(cpu *CPU) AddressInfo {
	ptr := cpu.Read16(cpu.PC + 1)

	var addr uint16
	if ptr&0x00FF == 0x00FF {
		// Simulate page boundary hardware bug
		lo := uint16(cpu.Read(ptr))
		hi := uint16(cpu.Read(ptr & 0xFF00))
		addr = (hi << 8) | lo
	} else {
		// Behave normally
		lo := uint16(cpu.Read(ptr))
		hi := uint16(cpu.Read(ptr + 1))
		addr = (hi << 8) | lo
	}
	return AddressInfo{Address: addr}
}

// INDX implements "Indexed Indirect X" address mode.
// A zero-page (8-bit) base address is read from the instruction operand, then the X register is added to it with
// zero-page wraparound. The result is used as a pointer to fetch the final 16-bit address.
func INDX(cpu *CPU) AddressInfo {
	ptr := uint16((cpu.Read(cpu.PC+1) + cpu.X) & 0x00FF)
	lo := uint16(cpu.Read(ptr))
	hi := uint16(cpu.Read((ptr + 1) & 0xFF))
	addr := (hi << 8) | lo
	return AddressInfo{Address: addr}
}

// INDY implements "Indirect Indexed Y" address mode.
// A zero-page (8-bit) address is read from the instruction operand and used as a pointer to fetch a 16-bit base
// address. The Y register is added to form the final address. Some instructions require an additional clock cycle
// if this addition crosses a page boundary.
func INDY(cpu *CPU) AddressInfo {
	ptr := uint16(cpu.Read(cpu.PC + 1))
	lo := uint16(cpu.Read(ptr))
	hi := uint16(cpu.Read((ptr + 1) & 0xFF))
	baseAddr := (hi << 8) | lo
	addr := baseAddr + uint16(cpu.Y)
	pageChanged := pagesDiffer(baseAddr, addr)
	return AddressInfo{Address: addr, PageChanged: pageChanged}
}
