default: build run

test:
	go test internshipApplicationTemplate/pkg/db/charta internshipApplicationTemplate/pkg/service

build:
	go build -o build/ cmd/*.go

run:
	./build/main $(ARGS)

clean:
	rm -rf build