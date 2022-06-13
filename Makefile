.PHONY : test build clean format

ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

test:
	$(foreach pkg, $(ALL_PACKAGES),\
	go test -race -short $(pkg);)

build:
	@echo "building binary"
	@go build -o gorengan  github.com/musobarlab/gorengan/cmd/gorengan

clean:
	@echo "cleaning unused file"
	rm -rf gorengan \
	&& rm -rf *.txt

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w