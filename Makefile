#!/usr/bin/make -f

export GO111MODULE = on

install:
	@echo "installing"
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/autoCode ./cmd
