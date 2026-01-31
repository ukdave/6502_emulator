; fibonacci.s
; Iterative Fibonacci calculation on a flat 64 KB 6502 system
;
; Computes F(10) using an iterative algorithm:
;   F(0) = 0
;   F(1) = 1
;   F(n) = F(n-1) + F(n-2)
;
; Final result is stored at memory location $0003.
; Expected result:
;   F(10) = 55 (0x37)

        .segment "CODE"

main:
        ; Initialize first Fibonacci value: a = 0
        LDA #$00
        STA $0000       ; a = 0

        ; Initialize second Fibonacci value: b = 1
        LDA #$01
        STA $0001       ; b = 1

        ; Set loop counter to compute F(10)
        ; Each iteration advances the sequence by one step
        LDX #$0A        ; X = 10

loop:
        ; temp = a + b
        LDA $0000       ; load a
        CLC             ; clear carry before addition
        ADC $0001       ; A = a + b
        STA $0002       ; store temporary result

        ; a = b
        LDA $0001
        STA $0000

        ; b = temp
        LDA $0002
        STA $0001

        ; Decrement loop counter and continue if not zero
        DEX
        BNE loop

        ; Store final Fibonacci result (a)
        ; After the final iteration, 'a' contains F(10)
        LDA $0000
        STA $0003       ; result = 55 (0x37)

hang:
        ; Infinite loop to prevent falling through
        JMP hang
