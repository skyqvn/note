FROM golang:latest

#设置工作目录
WORKDIR $GOPATH/src/app/src/main

#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOPATH=$HOME/go

#go构建可执行文件
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN go build -o ../../bin/app

#暴露端口
EXPOSE 80

#最终运行docker命令
ENTRYPOINT ["../../bin/app"]
