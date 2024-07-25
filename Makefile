# Variables
APP_NAME := af-api
MAIN_FILE := main.go
DEBUG_OUTPUT := $(APP_NAME).out
RELEASE_OUTPUT := $(APP_NAME)

# Default target
.PHONY: all
all: debug 
	
# Debug build
.PHONY: debug
debug:
	@echo "Building debug version..."
	@go build -o $(DEBUG_OUTPUT) $(MAIN_FILE)
	@echo "Build complete: $(DEBUG_OUTPUT)"

# Release build
.PHONY: release
release:
	@echo "Building release version..."
	@go build -o $(RELEASE_OUTPUT) $(MAIN_FILE)
	@echo "Build complete: $(RELEASE_OUTPUT)"

# Clean target
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -f $(DEBUG_OUTPUT) $(RELEASE_OUTPUT)
	@echo "Clean complete."
