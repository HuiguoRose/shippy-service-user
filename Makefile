build-protoc:
	protoc -I. --micro_out=. --go_out=.  proto/user/user.proto
build:
	docker build -t shippy-service-user .
run:
	docker run --net="host" -e MICRO_BROKER=nats -e MICRO_BROKER_ADDRESS=localhost:4222 --rm=true -e MICRO_REGISTRY=mdns -e MICRO_ADDRESS=":50051" -e DB_NAME=postgres -e DB_HOST=localhost -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=postgres -p 50053:50051 shippy-service-user
	#docker run --rm=true -e MICRO_REGISTRY=mdns -e MICRO_ADDRESS=":50051" -e DB_NAME=postgres -e DB_HOST=localhost -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=postgres -p 50053:50051 shippy-service-user
