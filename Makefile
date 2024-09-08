CURRENTDIR := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
generate: ${CURRENTDIR}protos/password_policy.proto
	PROTOS_PATH=${CURRENTDIR}protos GENERATED_PATH=${CURRENTDIR} go generate agent/password.go

build: client agent

CLIENTSRC := $(shell find client -name '*.go')

client: generate ${CLIENTSRC}
	go build -o bin/client ./client/main.go

AGENTSRC := $(shell find . -name '*.go' ! -path './client/*')

agent: generate ${AGENTSRC}
	go build -o bin/agent

test: build
	go test ${CURRENTDIR}agent

clean:
	rm -fr ${CURRENTDIR}bin
	rm -fr ${CURRENTDIR}generated
