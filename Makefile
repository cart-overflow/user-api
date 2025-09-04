# Set up environment
setup:
	@lefthook install
	@echo "✅ Environment ready"

# Generate Go code from Protocol Buffer files
gen-proto:
	protoc --go_out=. --go-grpc_out=. pkg/pb/user.proto

.PHONY: setup gen-proto