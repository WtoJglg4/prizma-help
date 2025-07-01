# Define the base paths and options
OPENAPI_SPEC_DIR = ./
TS_OUTPUT_BASE_DIR = openapi-ts
GO_OUTPUT_BASE_DIR = ../api
GO_GENERATOR_CONFIG = config.yaml

# TypeScript generation options
# Note: npmName will be appended with the API name later
TS_GENERATOR_BASE_OPTS = -g typescript --additional-properties=nullSafeAdditionalProps=true,supportsES6=true,npmName=@fp-api/

# --- Targets ---

# Default target: generate all if API is not specified, or the specific API if specified
all: check-generator
ifdef API
	@make generate-go API=$(API)
	@make generate-ts API=$(API)
else
	@make generate-go
	@make generate-ts
endif

# Check if the generators are installed
check-generator: check-go-generator check-ts-generator

check-ts-generator:
	@command -v openapi-generator-cli >/dev/null 2>&1 || make install-ts

check-go-generator:
	@command -v ogen >/dev/null 2>&1 || make install-go-generator

# Install the generators
install-ts:
	@echo "Installing openapi-generator-cli..."
	npm install @openapitools/openapi-generator-cli -g

install-go-generator:
	@echo "Installing ogen..."
	go install github.com/ogen-go/ogen/cmd/ogen@latest
	@echo "Make sure $(shell go env GOPATH)/bin is in your PATH"

# --- Generation Logic ---

# Generate Go server code
generate-go: check-go-generator
ifdef API
	# Generate for a single specified API
	@echo "Generating Go server code for API: $(API)..."
	$(eval SPEC_SUBDIR := $(OPENAPI_SPEC_DIR)/$(API)
	$(eval SPEC_FILE := $(firstword $(wildcard $(SPEC_SUBDIR)/openapi.yaml $(SPEC_SUBDIR)/openapi.yml $(SPEC_SUBDIR)/openapi.json)))
	@if [ -z "$(SPEC_FILE)" ]; then \
		echo "Error: No spec file found for API '$(API)' in $(SPEC_SUBDIR)"; \
		exit 1; \
	fi
	$(eval OUTPUT_DIR := $(GO_OUTPUT_BASE_DIR)/$(API))
	@echo "  Spec: $(SPEC_FILE)"
	@echo "  Output: $(OUTPUT_DIR)"
	rm -rf $(OUTPUT_DIR)
	mkdir -p $(OUTPUT_DIR)
	ogen -target $(OUTPUT_DIR) -package $(API) -config $(GO_GENERATOR_CONFIG) $(SPEC_FILE)
else
	# Generate for all APIs found
	@echo "Generating Go server code for all APIs..."
	@for spec_path in $(wildcard $(OPENAPI_SPEC_DIR)/*/openapi.yaml $(OPENAPI_SPEC_DIR)/*/openapi.yml $(OPENAPI_SPEC_DIR)/*/openapi.json); do \
		dir=$$(dirname $$spec_path); \
		name=$$(basename $$dir); \
		output_dir=$(GO_OUTPUT_BASE_DIR)/$$name; \
		echo "  Processing $$name (spec: $$spec_path)..."; \
		rm -rf $$output_dir; \
		mkdir -p $$output_dir; \
		ogen -target $$output_dir -package $$name -config $(GO_GENERATOR_CONFIG) $$spec_path; \
	done
endif

# Generate TypeScript code
generate-ts: check-ts-generator
ifdef API
	# Generate for a single specified API
	@echo "Generating TypeScript code for API: $(API)..."
	$(eval SPEC_SUBDIR := $(OPENAPI_SPEC_DIR)/$(API))
	$(eval SPEC_FILE := $(firstword $(wildcard $(SPEC_SUBDIR)/openapi.yaml $(SPEC_SUBDIR)/openapi.yml $(SPEC_SUBDIR)/openapi.json)))
	@if [ -z "$(SPEC_FILE)" ]; then \
		echo "Error: No spec file found for API '$(API)' in $(SPEC_SUBDIR)"; \
		exit 1; \
	fi
	$(eval OUTPUT_DIR := $(TS_OUTPUT_BASE_DIR)/$(API))
	$(eval TS_GENERATOR_OPTS := $(TS_GENERATOR_BASE_OPTS)$(API))
	@echo "  Spec: $(SPEC_FILE)"
	@echo "  Output: $(OUTPUT_DIR)"
	rm -rf $(OUTPUT_DIR)
	# Generate code
	openapi-generator-cli generate -i $(SPEC_FILE) $(TS_GENERATOR_OPTS) -o $(OUTPUT_DIR)
	# Copy original spec
	mkdir -p $(OUTPUT_DIR)/
	cp $(SPEC_FILE) $(OUTPUT_DIR)/
	# Install dependencies
	@echo "  Installing npm dependencies in $(OUTPUT_DIR)..."
	cd $(OUTPUT_DIR) && npm i && cd - > /dev/null
else
	# Generate for all APIs found
	@echo "Generating TypeScript code for all APIs..."
	@for spec_path in $(wildcard $(OPENAPI_SPEC_DIR)/*/openapi.yaml $(OPENAPI_SPEC_DIR)/*/openapi.yml $(OPENAPI_SPEC_DIR)/*/openapi.json); do \
		dir=$$(dirname $$spec_path); \
		name=$$(basename $$dir); \
		output_dir=$(TS_OUTPUT_BASE_DIR)/$$name; \
		ts_generator_opts=$(TS_GENERATOR_BASE_OPTS)$$name; \
		echo "  Processing $$name (spec: $$spec_path)..."; \
		rm -rf $$output_dir; \
		openapi-generator-cli generate -i $$spec_path $$ts_generator_opts -o $$output_dir; \
		cp $$spec_path $$output_dir/; \
		echo "    Installing npm dependencies in $$output_dir..."; \
		cd $$output_dir && npm i && cd - > /dev/null; \
	done
endif

# Clean target
clean:
	@echo "Cleaning generated files..."
	# Remove Go dirs only if they exist within the base dir structure
	@for name in $(shell basename -a $(wildcard $(OPENAPI_SPEC_DIR)/*)); do \
		if [ -d "$(GO_OUTPUT_BASE_DIR)/$$name" ]; then \
			echo "  Removing $(GO_OUTPUT_BASE_DIR)/$$name"; \
			rm -rf "$(GO_OUTPUT_BASE_DIR)/$$name"; \
		fi; \
	done
	# Remove TS base dir
	if [ -d "$(TS_OUTPUT_BASE_DIR)" ]; then \
		echo "  Removing $(TS_OUTPUT_BASE_DIR)"; \
		rm -rf $(TS_OUTPUT_BASE_DIR); \
	fi

.PHONY: all generate-go generate-ts clean check-generator check-ts-generator check-go-generator install-ts install-go-generator