.PHONY: all asm c clean

# ----------------------------
# Program lists
# ----------------------------

ASM_PROGRAMS := factorial fibonacci multiply
C_PROGRAMS   := factorial fibonacci multiply


# ----------------------------
# Intermediate and output files
# ----------------------------

ASM_BINS := $(ASM_PROGRAMS:%=build/asm/%.bin)
C_BINS   := $(C_PROGRAMS:%=build/c/%.bin)

ASM_OBJS := $(ASM_PROGRAMS:%=build/asm/%.o)
C_SRCS   := $(C_PROGRAMS:%=build/c/%.s)
C_OBJS   := $(C_PROGRAMS:%=build/c/%.o)

# Preserve intermediates
.PRECIOUS: $(ASM_OBJS) $(C_SRCS) $(C_OBJS)


# ----------------------------
# Default target
# ----------------------------
all: $(ASM_BINS) $(C_BINS)

asm: $(ASM_BINS)
c:   $(C_BINS)


# ----------------------------
# Build directories
# ----------------------------

build/asm build/c:
	mkdir -p $@


# ============================
# ASM pipeline
# ============================

build/asm/%.o: src/asm/%.s | build/asm
	ca65 -o $@ $<

build/asm/%.bin: build/asm/%.o linker.cfg
	ld65 -o $@ -C linker.cfg $<


# ============================
# C pipeline
# ============================

build/c/%.s: src/c/%.c | build/c
	cc65 -t none -O -o $@ $<

build/c/%.o: build/c/%.s
	ca65 -o $@ $<

build/c/%.bin: build/c/%.o
	ld65 -t none -o $@ $< none.lib


# ----------------------------
# Cleanup
# ----------------------------

clean:
	rm -rf build
