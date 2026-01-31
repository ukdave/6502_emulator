; multiply.s
; Minimal 6502 program on a flat 64 KB 6502 system
;
; Computes 3 * 10 by adding 3 to itself 10 times.
;
; Final result is stored at memory location $0000.
; Expected result:
;   3 * 10 = 30 (0x1E)

        .segment "CODE"

main:
        LDA #$00        ; A = 0
        LDX #$0A        ; X = 10

loop:
        CLC             ; clear carry flag before addition
        ADC #$03        ; A += 3
        DEX             ; X--
        BNE loop        ; if X != 0, repeat loop

        STA $0000       ; store result (30) at 0x0000

hang:
        JMP hang        ; infinite loop
