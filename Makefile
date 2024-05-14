.PHONY : default
default: help

.PHONY: test
test:unit bench

.PHONY : unit
unit:
	cd src/list/ && go test -v -cover

.PHONY : bench
bench:
	cd src/list/ && go test -bench=. -benchmem -benchtime=20s

.PHONY : build
build:
	go build -o bin/main src/cmd/main.go
	./bin/main

.PHONY : clean
clean:
	rm -rf bin/main

help:
	@echo " default"
	@echo " test"
	@echo " unit"
	@echo " bench"
	@echo " build"
	@echo " clean"
