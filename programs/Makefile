.PHONY: all clean

# List of programs (without extension)
PROGRAMS := factorial fibonacci multiply

# Derived file lists
BINS := $(PROGRAMS:%=build/%.bin)
OBJS := $(PROGRAMS:%=build/%.o)

# Prevent Make from automatically deleting .o files
.PRECIOUS: $(OBJS)

all: $(BINS)

build:
	mkdir -p build

# Assemble .s -> .o
build/%.o: src/%.s | build
	ca65 $< -o $@

# Link .o -> .bin
build/%.bin: build/%.o linker.cfg | build
	ld65 $< -C linker.cfg -o $@

clean:
	rm -rf build
