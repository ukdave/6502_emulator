// fibonacci.c
// Minimal freestanding C for 6502 emulator
//
// Computes the 10th Fibonacci number and stores the result at 0x0000
//
// Expected result:
//   F(10) = 55 (0x37)

#include <stdint.h>

void main(void) {
  uint8_t fib_prev = 0;
  uint8_t fib_curr = 1;
  uint8_t i, next;

  for (i = 2; i <= 10; i++) {
      next = fib_prev + fib_curr;
      fib_prev = fib_curr;
      fib_curr = next;
  }

  // Store the final result in zero page 0x0000
  *((volatile uint8_t*)0x0000) = fib_curr;

  // Infinite loop to halt
  while (1) {}
}
