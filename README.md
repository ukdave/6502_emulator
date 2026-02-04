# 6502 Emulator

[![Build Status](https://github.com/ukdave/6502_emulator/actions/workflows/main.yml/badge.svg)](https://github.com/ukdave/6502_emulator/actions/workflows/main.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ukdave/6502_emulator)
[![License](https://img.shields.io/github/license/ukdave/6502_emulator)](https://github.com/ukdave/6502_emulator/blob/main/LICENSE.txt)

This is a standalone MOS 6502 CPU emulator written in Go.

It emulates only the 6502 processor itself — no NES, C64, or other system hardware is implemented.

The emulator provides a flat 64 KB address space, backed entirely by RAM, and has been built as learning exercise rather than a full-featured system emulator.

The terminal interface (TUI) was built using [Bubble Tea](https://github.com/charmbracelet/bubbletea).

![Demo](demo.gif?raw=true "Demo")

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

# Install staticcheck
go install honnef.co/go/tools/cmd/staticcheck@latest
asdf reshim golang

# Lint, test, and build code
make

# Run emulator
./6502_emulator example.bin
# or
go run main.go example.bin
```

## Writing 6502 programs

Programs can be written in assembly or C, built into a binary (.bin) file using the [cc65](https://github.com/cc65/cc65) toolchain, and then loaded into the emulator.

```bash
brew install cc65
```

Some sample programs are provided in the `programs/` directory.

```bash
# build all programs
make programs
```

All of the included sample assembly programs are designed to be loaded at memory address 0x8000.

All of the included sample C programs are designed to be loaded at memory address 0x1000.

### Assembly

Steps to build an assembly program:

```bash
# Assemble the assembly code into machine code (creates a re-locatable object file)
ca65 -o my_program.o my_program.s

# Link our machine code into a binary file with everything laid out in specific memory locations using the linker configuration
ld65 -o my_program.bin -C linker.cfg my_program.o

# Load binary file into the emulator at address 0x8000
go run main.go -s 0x8000 my_program.bin
```

### C

Programs can also be written in C and compiled down into a binary file using cc65 that can be loaded into the 6502 emulator.

When compiling C code with cc65 we need to tell it what the target system is (e.g. a C64 or a NES) so that it knows what memory locations our program can be placed in. As this project is just a bare 6502 emulator (i.e. it doesn't implement a full C64 or NES system) we need to specify the `none` target. This assumes a flat 64k memory and requires the program to be loaded at memory address 0x1000.

Steps to compile a C program for a bare 6502 emulator:

```bash
# Compile C code into assembly code
cc65 -t none -o my_program.s my_program.c

# Assemble the assembly code into machine code (creates a re-locatable object file)
ca65 -o my_program.o my_program.s

# Link our machine code and the C runtime into a binary file with everything laid out in specific memory locations for the target system
ld65 -t none -o my_program.bin my_program.o none.lib

# Load binary file into the emulator at address 0x1000
go run main.go -s 0x1000 my_program.bin
```
