BINARY=chatroom

# colors compatible setting
CRED:=$(shell tput setaf 1 2>/dev/null)
CGREEN:=$(shell tput setaf 2 2>/dev/null)
CYELLOW:=$(shell tput setaf 3 2>/dev/null)
CEND:=$(shell tput sgr0 2>/dev/null)

.PHONY: all
all: | fmt build

# Code format
.PHONY: fmt
fmt: 
	@echo "$(CGREEN)Run gofmt ...$(CEND)"
	@echo "gofmt -l -s -w ..."
	@ret=0 && for d in $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		gofmt -l -s -w $$d/*.go || ret=$$? ; \
	done ; exit $$ret

# build
.PHONY: build-darwin
build-darwin: fmt
	@echo "$(CGREEN)Building for darwin ...$(CEND)"
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o deploy/${BINARY} main.go
	cp config/config_dev.yaml deploy/config/config.yaml
	@echo "$(CGREEN)Build Success!$(CEND)"

# build
.PHONY: build
build: fmt
	@echo "$(CGREEN)Building for linux ...$(CEND)"
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deploy/${BINARY} main.go
	cp config/config.yaml deploy/config/config.yaml
	@echo "$(CGREEN)Build Success!$(CEND)"

# build
.PHONY: build-dev
build-dev: fmt
	@echo "$(CGREEN)Building dev for linux ...$(CEND)"
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deploy/${BINARY} main.go
	cp config/config_dev.yaml deploy/config/config.yaml
	@echo "$(CGREEN)Build Success!$(CEND)"

# install
install: build
	@echo "$(CGREEN)Install ...$(CEND)"
	go install ./...
	@echo "$(CGREEN)Install Success!$(CEND)"

# clean
.PHONY: clean
clean:
	@echo "$(CGREEN)Cleanup ...$(CEND)"
	go clean
	@rm -f deploy/${BINARY}
	@echo "rm -f deploy/${BINARY}"
	@for GOOS in darwin linux windows; do \
	    for GOARCH in 386 amd64; do \
			rm -f deploy/${BINARY}.$${GOOS}-$${GOARCH} ;\
		done ;\
	done
	rm -f ${BINARY} coverage.* test/tmp/*
	find . -name "*.log" -delete

# protoc build
.PHONY: protoc
protoc:
	protoc --gogo_out=. pkg/protocol/*.proto
