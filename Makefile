.PHONY: full build clean

full: clean build copy

build:
	@mkdir -p build
	go build -trimpath -ldflags '-w -s' -o ./build/server ./cmd/

clean:
	rm -rf build/

copy:
	@mkdir -p build
	cp -r ./template ./build/
