clean:
	rm -rf ./build

prepare:
	mkdir build

build:
	make prepare && go build -o ./build/goliatid

run:
	./build/goliatid

.PHONY:
	build run
