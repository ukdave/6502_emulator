# 6502 Emulator

This is a standalone MOS 6502 CPU emulator written in Go.

It emulates only the 6502 processor itself — no NES, C64, or other system hardware is implemented.

The emulator provides a flat 64 KB address space, backed entirely by RAM, and has been built as learning exercise rather than a full-featured system emulator.

## Overview

### What this is (and isn’t)

This project is:
- A learning-focused 6502 CPU emulator
- Written to understand the 6502 instruction set and CPU behaviour
- Backed by a simple, flat 64 KB RAM model
- Suitable for running small test programs and experiments

This project is not:
- A full NES or C64 emulator
- Intended for running commercial games or software

If you are looking for a fully working NES or C64 emulator, this is not the right project — there are excellent alternatives elsewhere.

CPU details:
- MOS 6502 compatible instruction set
- Decimal mode not implemented (matching the NES 6502 variant)
- Illegal opcodes not implemented
- No memory-mapped I/O or peripheral devices
- No PPU, APU, timers, or interrupts beyond basic CPU behaviour

### Inspiration

This project is loosely based on the "NES Emulator from Scratch" series by javidx9 (OneLoneCoder), but focussing solely on the CPU rather than a full console:

* https://www.youtube.com/playlist?list=PLrOv9FMX8xJHqMvSGB_9G9nZZ_4IgteYf
* https://github.com/OneLoneCoder/olcNES

### Resources

As well as the above video series, the following resources have also been very useful during the development of this emulator:

* NES Dev wiki:
  * https://www.nesdev.org/wiki/CPU
* CPU datasheets:
  * http://archive.6502.org/datasheets/rockwell_r650x_r651x.pdf
  * https://www.princeton.edu/~mae412/HANDOUTS/Datasheets/6502.pdf
* An online 6502 emulator:
  * https://www.masswerk.at/6502/
* A NES emulator written in Go:
  * https://github.com/fogleman/nes

## Development

Installing and running (on a mac):

```bash
# Install go
brew install asdf
asdf plugin add golang
asdf install

# Lint, test, and build code
make

# Run emulator
./6502_emulator example.bin
# or
go run main.go example.bin
```
