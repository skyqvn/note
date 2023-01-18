#include "textflag.h"

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
