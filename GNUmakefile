HOSTNAME=peishuli.com
NAMESPACE=dev
NAME=snowflakeprivatelink
BINARY=terraform-provider-${NAME}
VERSION=0.1
OS_ARCH=linux_amd64

# default: testacc
default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

test:
	go build main.go && rm main

