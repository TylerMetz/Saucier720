# Set the Go executable name based on the OS
ifeq ($(OS),Windows_NT)
	EXECUTABLE = bin/mealdealz.exe
else
	EXECUTABLE = bin/mealdealz
endif

build:
	@go build -o $(EXECUTABLE)

run: build
	@$(EXECUTABLE)
