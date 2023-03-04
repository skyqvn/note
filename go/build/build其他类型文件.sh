#动态链接库
#go build -buildmode=c-shared -o name.so [name.go]
go build -buildmode=c-shared -o name.so
g++ main.cpp ./name.so -o main

#静态链接库
#go build -buildmode=c-archive -o name.a [name.go]
go build -buildmode=c-archive -o name.a
g++ main.cpp ./name.a -o main

#.o文件
go tool compile name.go

#反编译为汇编代码
go tool objdump name.o

#反编译某函数为汇编代码
go tool objdump -s nameFunc name.o

#编译plan9汇编
go tool compile -S name.go

#使用静态链接库
go tool compile -I ./name name.go
go tool link -o name -L ./name name.o

#不启动命令行
go build -ldflags="-H windowsgui"
