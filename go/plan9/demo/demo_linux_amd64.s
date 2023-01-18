#include "textflag.h"

DATA ·Data+0(SB)/8,$·Data+16(SB)
DATA ·Data+8(SB)/8,$11
DATA ·Data+16(SB)/16,$"hello world"
GLOBL ·Data(SB),RODATA,$32

TEXT  ·Print(SB),NOSPLIT,$0-0
    MOVQ $1,AX
    MOVQ $1,DI
    MOVQ ·Data+0(SB),SI
    MOVQ $11,DX
    SYSCALL
    RET
