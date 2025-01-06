# Makefile

# Define the output directory
BUILD_DIR = build

# Define the commands
HELLO_CMD = ./cmd/hello/hello.go
MATH_CMD = ./cmd/math/math.go

# Define the targets
HELLO_TARGET = $(BUILD_DIR)/hello
MATH_TARGET = $(BUILD_DIR)/math

# Default target
all: $(HELLO_TARGET) $(MATH_TARGET)

# Build hello executable
$(HELLO_TARGET): $(HELLO_CMD)
	go build -o $@ $<

# Build math executable
$(MATH_TARGET): $(MATH_CMD)
	go build -o $@ $<

# Clean up build artifacts
clean:
	rm -f $(HELLO_TARGET) $(MATH_TARGET)
