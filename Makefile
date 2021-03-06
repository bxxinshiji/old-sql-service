dev:
	make proto && make build && make run

protoc:
	protoc -I . --go_out=plugins=micro:. proto/goods/good.proto
	protoc -I . --go_out=plugins=micro:. proto/inventory/inventory.proto
    
build:
	# GOOS=linux GOARCH=amd64 go build
	docker build -t old-sql-service .

run:
# docker run -p 8080:8080 -e MICRO_REGISTRY=mdns microhq/micro api --handler=rpc --address=:8080 --namespace=xinshiji
	docker run --net="host" -p 12801 \
                    -e MICRO_REGISTRY=mdns  \
                    -e SERVICE_NAME="xinshiji.old_sql"  \
                    -e DB_DRIVER="odbc"  \
                    -e DB_NAME="stmis1"  \
                    -e DB_HOST="192.168.1.26"  \
                    -e DB_PORT="1433"  \
                    -e DB_USER="sa"  \
                    -e DB_PASSWORD=""  \
                    old-sql-service
