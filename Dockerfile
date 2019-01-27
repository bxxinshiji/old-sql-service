# 暂未将 Golang 集成到 docker 中
FROM golang:1.11-alpine3.8 as builder

# 安装 odbc 依赖
RUN apk add --no-cache gcc git musl-dev make zip unixodbc unixodbc-dev freetds

ADD docker/etc_odbcinst.ini /etc/odbcinst.ini
# 设置编译环境代理 一般为本机代理 ip 端口
ENV https_proxy http://192.168.1.68:8118
ENV GO111MODULE=off
# 设置运行环境并且编译
WORKDIR /go/src/github.com/bxxinshiji/old-sql-service
COPY . .
RUN go get
# RUN GOOS=linux GOARCH=amd64 go build
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .

# 运行环境
FROM alpine:3.8

# RUN apk --no-cache add ca-certificates

# 安装 odbc 依赖
RUN apk add --update unixodbc unixodbc-dev freetds

ADD docker/etc_odbcinst.ini /etc/odbcinst.ini

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/bxxinshiji/old-sql-service .

CMD ["./old-sql-service"]

