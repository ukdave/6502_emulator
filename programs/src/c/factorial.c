// factorial.c
// Minimal freestanding C for 6502 emulator
//
// Computes 5! (factorial of 5) and stores the result at 0x0000
//
// Expected result:
//   5! = 120 (0x78)

#include <stdint.h>

void main(void) {
  uint8_t n = 5;
  uint8_t result = 1;

  // Compute factorial
  while (n > 1) {
    result *= n;
    n--;
  }

  // Store result at memory address 0x0000
  *((volatile uint8_t*)0x0000) = result;

  // Infinite loop to halt
  while (1) {}
}
