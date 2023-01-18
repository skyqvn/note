#动态
#go build -buildmode=c-shared -o name.so [name.go]
go build -buildmode=c-shared -o sum.so
g++ main.cpp ./sum.so -o main
#静态
#go build -buildmode=c-archive -o name.a [name.go]
go build -buildmode=c-archive -o sum.a
g++ main.cpp ./sum.a -o main
