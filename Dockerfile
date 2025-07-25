# Multi-stage build for NER Service Go - Linux AMD64
FROM golang:1.24-bullseye AS builder

# Install system dependencies
RUN apt-get update && apt-get install -y \
    gcc \
    g++ \
    make \
    wget \
    unzip \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

# Install MITIE from source
RUN wget https://github.com/mit-nlp/MITIE/archive/v0.7.tar.gz -O mitie.tar.gz && \
    tar -xzf mitie.tar.gz && \
    cd MITIE-0.7 && \
    make -j$(nproc) && \
    cp mitielib/libmitie.so /usr/lib/ && \
    cp mitielib/include/mitie.h /usr/include/ && \
    ldconfig && \
    cd .. && \
    rm -rf MITIE-0.7 mitie.tar.gz

# Set working directory
WORKDIR /app

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy source code
COPY . .

# Set CGO flags for MITIE
ENV CGO_ENABLED=1
ENV CGO_CFLAGS="-I/usr/include"
ENV CGO_LDFLAGS="-L/usr/lib"
ENV GOOS=linux
ENV GOARCH=amd64

# Build the server binary
RUN go build -ldflags="-s -w" -o ner-server cmd/server/main.go

# Build the CLI binary  
RUN go build -ldflags="-s -w" -o ner-cli cmd/cli/main.go

# Production stage - minimal runtime image
FROM debian:bullseye-slim

# Install runtime dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Copy MITIE library from builder stage
COPY --from=builder /usr/lib/libmitie.so /usr/lib/
RUN ldconfig

# Create non-root user for security
RUN groupadd -g 1001 appgroup && \
    useradd -u 1001 -g appgroup -s /bin/bash -m appuser

# Create app directory and set ownership
WORKDIR /app
RUN chown -R appuser:appgroup /app

# Copy binaries from builder stage
COPY --from=builder --chown=appuser:appgroup /app/ner-server /app/ner-cli ./

# Create models directory
RUN mkdir -p models && chown -R appuser:appgroup models

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Set environment variables
ENV MITIE_MODEL_PATH=/app/models/ner_model.dat
ENV PORT=8080
ENV GIN_MODE=release

# Add healthcheck
HEALTHCHECK --interval=30s --timeout=10s --start-period=40s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider --timeout=5 http://localhost:8080/health || exit 1

# Default command
CMD ["./ner-server"]