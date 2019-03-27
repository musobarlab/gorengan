.PHONY : test

ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

test:
	$(foreach pkg, $(ALL_PACKAGES),\
	go test -race -short $(pkg);)