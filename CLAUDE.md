# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Named Entity Recognition (NER) service for Spanish text using MITIE (MIT Information Extraction library). The application provides both CLI and HTTP API interfaces for extracting named entities from Spanish text.

### Architecture

- **Language**: Go 1.24+
- **NER Engine**: MITIE with Spanish language model
- **Web Framework**: Gin for HTTP server
- **CLI Framework**: Cobra for command-line interface
- **Build System**: Make with CGO support for MITIE

### Core Components

```
├── cmd/
│   ├── server/          # HTTP server entry point (cmd/server/main.go)
│   └── cli/             # CLI tool entry point (cmd/cli/main.go)
├── internal/
│   ├── config/          # Configuration management (internal/config/config.go)
│   └── ner/             # NER service core logic
│       ├── service.go   # Main NER service with MITIE integration
│       └── types.go     # Entity and request/response types
├── models/              # MITIE model files (downloaded separately)
└── Dockerfile          # Multi-stage build for Linux AMD64 deployment
```

## Development Commands

### Essential Make Commands
```bash
make setup          # Full setup: install MITIE + download model + build
make build          # Build both server and CLI binaries
make server         # Build server binary only
make cli            # Build CLI binary only
make deps           # Download Go dependencies
make download-model # Download Spanish MITIE model (~450MB)
make test           # Run tests
make clean          # Clean build artifacts
make run-server     # Run server in development mode
make run-cli        # Run CLI in development mode
```

### CGO Configuration
The project requires CGO for MITIE integration:
- **macOS**: Uses Homebrew paths (`/opt/homebrew/include`, `/opt/homebrew/lib`)
- **Docker**: Builds MITIE from source in multi-stage build
- All builds automatically set appropriate CGO flags

### Testing
```bash
go test ./...       # Run all tests
make test          # Same as above via Makefile
```

### Model Management
The Spanish MITIE model is ~450MB and excluded from git:
- Download: `make download-model`
- Default path: `models/ner_model.dat`
- Override via: `MITIE_MODEL_PATH` environment variable

## HTTP Server Architecture

### Server Structure (cmd/server/main.go)
- Uses Gin framework with minimal middleware
- Two endpoints: `/health` and `/ner`
- Flexible input handling: JSON, form data, multipart
- Graceful error handling with proper HTTP status codes

### Input Formats Supported
1. **JSON**: `{"text": "your text here"}`
2. **URL-encoded form**: `text=your text here`  
3. **Multipart form**: Form field named `text`

### Response Format
Always returns JSON array of entities:
```json
[
  {
    "tag": "PLACE",
    "score": "0.758809", 
    "label": "Madrid"
  }
]
```

## CLI Architecture (cmd/cli/main.go)

### CLI Features
- Text input via command argument or `--file` flag
- JSON output with `--json` flag
- Custom model path with `--model` flag
- Human-readable output by default

### Usage Patterns
```bash
./ner-cli "text to analyze"
./ner-cli --file input.txt
./ner-cli --json "text" | jq
./ner-cli --model /custom/path.dat "text"
```

## NER Service Core (internal/ner/)

### Service Pattern
- `Service` struct wraps MITIE extractor
- `NewService()` initializes with model path
- `ExtractEntities()` processes text and returns standardized entities
- `Close()` properly frees MITIE resources

### Entity Mapping
Maps MITIE tags to standard format:
- PERSON → PERSON
- LOCATION → LOCATION  
- ORGANIZATION → ORGANIZATION
- MISC → MISC
- Default → PLACE

## Configuration (internal/config/)

### Environment Variables
- `MITIE_MODEL_PATH`: Model file path (default: `models/ner_model.dat`)
- `PORT`: HTTP server port (default: `8080`)
- `GIN_MODE`: Gin framework mode for production (`release`)

## Docker Deployment

### Multi-stage Build Strategy
1. **Builder stage**: Installs MITIE from source, downloads Spanish model, builds Go binaries
2. **Runtime stage**: Minimal Debian image with only runtime dependencies
3. **Security**: Non-root user, health checks, resource limits

### Docker Commands
```bash
docker build --platform linux/amd64 -t ner-service-go .
docker-compose up
```

### Image Characteristics
- **Size**: ~539MB (includes Spanish model)
- **Platform**: linux/amd64 optimized
- **Security**: Non-root execution, minimal attack surface
- **Health checks**: Built-in `/health` endpoint monitoring

## Dependencies

### Key External Dependencies
- `github.com/gin-gonic/gin`: HTTP framework
- `github.com/spf13/cobra`: CLI framework  
- `github.com/sbl/ner`: Go bindings for MITIE
- MITIE C++ library (system dependency)

### System Requirements
- **Go**: 1.24+
- **MITIE**: System library (Homebrew on macOS, source build in Docker)
- **CGO**: Required for MITIE bindings
- **Memory**: ~512MB minimum for model loading

## Common Development Patterns

### Error Handling
- Service layer returns wrapped errors with context
- HTTP handlers log errors and return appropriate status codes
- CLI exits with status 1 on errors

### Resource Management
- Always call `nerService.Close()` to free MITIE resources
- Use defer statements for cleanup
- Handle model loading failures gracefully

### Testing Considerations
- Model file must be present for integration tests
- Use `make download-model` before running tests
- Mock the NER service for unit tests to avoid model dependency