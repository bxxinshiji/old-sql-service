# 暂未将 Golang 集成到 docker 中
FROM golang:1.8.0

RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get install -y \
    unixodbc unixodbc-dev freetds-dev freetds-bin tdsodbc

ADD docker/etc_freetds_freetds.conf /etc/freetds/freetds.conf
ADD docker/etc_odbc.ini /etc/odbc.ini
ADD docker/etc_odbcinst.ini /etc/odbcinst.ini

WORKDIR /go/src/gitee.com/xsjcs/old-sql-service
COPY . .
RUN go get
RUN GOOS=linux GOARCH=amd64 go build

CMD ["./old-sql-service"]