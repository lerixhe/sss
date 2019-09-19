
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/PutUserInfo/PutUserInfo.proto

.PHONY: build
build: proto

	go build -o PutUserInfo-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t PutUserInfo-srv:latest
