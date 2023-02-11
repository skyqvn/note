#include "textflag.h"

//DEMO1
DATA ·Data+0(SB)/8,$·Data+16(SB)
DATA ·Data+8(SB)/8,$12
DATA ·Data+16(SB)/16,$"hello world\n"
GLOBL ·Data(SB),RODATA,$32

TEXT  ·Print(SB),NOSPLIT,$0-0
    MOVQ $1,AX
    MOVQ $1,DI
    MOVQ ·Data+0(SB),SI
    MOVQ ·Data+8(SB),DX
    SYSCALL
    RET

//DEMO2
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
