.PHONY: build clean

build:
	@echo "Building..."
	go build -o bin/ 

clean:
	@echo "Cleaning..."
	rm bin/*
