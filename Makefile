build:
	@echo "Compiling..."
	@go build -o bin/aioli .

run:
	@echo "Running..."
	@./bin/aioli

clean:
	@echo "Cleaning all binaries..."
	@rm -rf bin


