# Copyright © 2021 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: MPL-2.0

default: build

build:
	go build -o bin/terraform-provider-tmc
	mkdir -p ~/.terraform.d/plugins/vmware/tanzu/tmc/0.1.1/darwin_amd64/
	cp bin/terraform-provider-tmc ~/.terraform.d/plugins/vmware/tanzu/tmc/0.1.1/darwin_amd64/

test: | gofmt vet lint
	go mod tidy
	go test ./internal/... -cover

# Run go fmt against code
gofmt:
	@echo Checking code is gofmted
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Linter
lint: gofmt
	golangci-lint run -c ./.golangci.yml ./internal/...
