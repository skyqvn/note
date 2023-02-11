TEXT ·Factorial(SB),$0-16
    SUBQ $24,SP
    MOVQ BP,16(SP)
    LEAQ 16(SP),BP
    MOVQ $1,ret+40(SP)
    MOVQ a+32(SP),AX
    MOVQ AX,t+8(SP)
    begin:
    CMPQ t+8(SP),$0
    JLE end
    MOVQ ret+40(SP),AX
    IMULQ t+8(SP),AX
    MOVQ AX,ret+40(SP)
    DECQ t+8(SP)
    JMP begin
    end:
    MOVQ 16(SP),BP
    ADDQ $24,SP
    RET
