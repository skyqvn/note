;https://blog.csdn.net/icandoit_2014/article/details/87385583
;https://zhuanlan.zhihu.com/p/333461542
;https://blog.csdn.net/chuck_huang/article/details/79922595
.section .data
msg :
    .ascii "hello world\n"
    len = . - msg

.section .text
.globl _start
_start:
    mov $1,%rax
    mov $1,%rdi
    mov $msg,%rsi
    mov $len,%rdx
    syscall
    mov $60,%rax
    mov $0,%rdi
    syscall
