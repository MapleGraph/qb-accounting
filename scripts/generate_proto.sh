#!/bin/bash
# Proto generation script for qb-accounting
# Generates Go code from proto files for gRPC services

set -e

PROTO_DIR="./proto"
OUT_DIR="./internal/proto"

echo "Generating proto files..."

# Generate Setup service
echo "Generating Setup service..."
protoc \
    --proto_path=$PROTO_DIR/setup \
    --go_out=$OUT_DIR/setup \
    --go_opt=paths=source_relative \
    --go-grpc_out=$OUT_DIR/setup \
    --go-grpc_opt=paths=source_relative \
    $PROTO_DIR/setup/setup.proto

# Generate Employee service
echo "Generating Employee service..."
protoc \
    --proto_path=$PROTO_DIR/employee \
    --go_out=$OUT_DIR/employee \
    --go_opt=paths=source_relative \
    --go-grpc_out=$OUT_DIR/employee \
    --go-grpc_opt=paths=source_relative \
    $PROTO_DIR/employee/user.proto

# Generate Notification service
echo "Generating Notification service..."
protoc \
    --proto_path=$PROTO_DIR/notification \
    --go_out=$OUT_DIR/notification \
    --go_opt=paths=source_relative \
    --go-grpc_out=$OUT_DIR/notification \
    --go-grpc_opt=paths=source_relative \
    $PROTO_DIR/notification/notification.proto

# Generate Catalogue (Product) service
echo "Generating Catalogue service..."
protoc \
    --proto_path=$PROTO_DIR/catalogue \
    --go_out=$OUT_DIR/catalogue \
    --go_opt=paths=source_relative \
    --go-grpc_out=$OUT_DIR/catalogue \
    --go-grpc_opt=paths=source_relative \
    $PROTO_DIR/catalogue/product.proto

echo "Proto generation complete!"
