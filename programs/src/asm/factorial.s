; factorial.s
; Factorial calculation on a flat 64 KB 6502 system
;
; Computes 5! (factorial of 5) using repeated addition
;
; Final result is stored at memory location $0003.
; Expected result:
;   5! = 120 (0x78)

        .segment "CODE"

main:
        ; Initialize
        LDA #$05        ; n = 5
        STA $0001       ; $0001 = n (counter for outer loop)
        LDA #$01
        STA $0000       ; $0000 = result (initial factorial = 1)

fact_loop:
        ; Multiply result by n
        LDA $0000       ; load current result
        STA $0002       ; $0002 = temp storage for multiplication

        LDA #$00
        STA $0000       ; clear result accumulator

        LDY $0001       ; Y = current multiplier (n)

mul_inner:
        CLC
        LDA $0000
        ADC $0002       ; result += temp
        STA $0000

        DEY
        BNE mul_inner   ; repeat until Y = 0

        ; Decrement n and check for loop end
        DEC $0001
        BNE fact_loop   ; continue outer loop until n = 0

hang:
        JMP hang
