// multiply.c
// Minimal freestanding C for 6502 emulator
//
// Computes 3 * 10 and stores the result at 0x0000
//
// Expected result:
//   3 * 10 = 30 (0x1E)

#include <stdint.h>

void main(void) {
  uint8_t a = 3;
  uint8_t b = 10;
  uint8_t result;

  result = a * b;

  // Store the final result in zero page 0x0000
  *((volatile uint8_t*)0x0000) = result;

  // Infinite loop to halt
  while (1) {}
}
