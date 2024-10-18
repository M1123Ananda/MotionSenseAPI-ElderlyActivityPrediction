build: clean
	go build -o build/motionsense -v

run: build
	./build/motionsense

test:
	go test -v ./...

clean:
	go clean
	rm -rf build

install:
	go get ./...

.PHONY: clean all test clean