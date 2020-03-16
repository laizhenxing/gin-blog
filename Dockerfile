#指定基础镜像
FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
#指定工作目录
WORKDIR $GOPATH/src/gin-blog
#copy指令从构建上下文中的<原路径>的文件/目录复制一份到镜像内的<目标路径>位置
COPY . $GOPATH/src/gin-blog
#执行命令
RUN go build .

# EXPOSE指令声明运行时容器提供一个端口，这只是一个声明，在运行时不会因为这个声明应用就会开启这个端口的服务
EXPOSE 8000
# ENTRYPOINT是指定容器启动程序及参数
ENTRYPOINT ["./gin-blog"]