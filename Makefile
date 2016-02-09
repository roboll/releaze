GVT     := $(GOPATH)/bin/gvt
$(GVT): ; GO15VENDOREXPERIMENT=1 go get github.com/FiloSottile/gvt

.PHONY: tools
tools: | $(GVT)

.PHONY: build
build:
	GO15VENDOREXPERIMENT=1 go build ./...

.PHONY: test
test:
	GO15VENDOREXPERIMENT=1 go test ./...
	GO15VENDOREXPERIMENT=1 go vet ./...

.PHONY: ci
ci: $(GVT) build test
