analyze:
	@echo "Running analysis..."
	golangci-lint run
	@echo "Analysis completed"
format:
	@echo "Running format..."
	golangci-lint run --fix
	@echo "Format completed"