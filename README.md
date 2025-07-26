# NER Service Go

[![Version](https://img.shields.io/badge/Version-1.0.0-brightgreen)](https://github.com/your-repo/ner-service-go/releases)
[![Go Version](https://img.shields.io/badge/Go-1.24+-blue?logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Docker Hub](https://img.shields.io/badge/Docker%20Hub-drzippie%2Fner--service-blue?logo=docker)](https://hub.docker.com/r/drzippie/ner-service)
[![Docker Pulls](https://img.shields.io/docker/pulls/drzippie/ner-service)](https://hub.docker.com/r/drzippie/ner-service)
[![Docker Image Size](https://img.shields.io/docker/image-size/drzippie/ner-service/latest)](https://hub.docker.com/r/drzippie/ner-service)
[![MITIE](https://img.shields.io/badge/NER-MITIE%20Spanish-orange)](https://github.com/mit-nlp/MITIE)

A Named Entity Recognition (NER) service for Spanish text using MITIE (MIT Information Extraction library). This service provides both CLI and HTTP API interfaces for extracting named entities from Spanish text.

## 🐳 Official Docker Image

**Ready-to-use Docker image**: `drzippie/ner-service:latest`

- ✅ **Spanish MITIE models included** (~539MB)
- ✅ **Production-ready** with health checks
- ✅ **Security hardened** (non-root, minimal dependencies)
- ✅ **Multi-platform** support (linux/amd64)

## ✅ Status: Fully Working

The application is **fully functional and ready to use**. Both CLI and HTTP server modes are operational with the Spanish language model.

## 🚀 Quick Start

### Using Docker (Recommended)
```bash
# Pull and run the pre-built image
docker run -p 8080:8080 drzippie/ner-service:latest

# Test the API
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Juan vive en Madrid y trabaja en Google España."}'
```

### Local Development
```bash
# 1. Install dependencies
brew install go mitie

# 2. Download Spanish model (~450MB)
make download-model

# 3. Build the application
make build

# 4. Test CLI
./ner-cli "Juan vive en Madrid y trabaja en Google España."

# 5. Start HTTP server
./ner-server

# 6. Test API endpoint
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Juan vive en Madrid y trabaja en Google España."}'
```

## Features

- **Named Entity Recognition** for Spanish text
- **Multiple entity types**: PERSON, LOCATION, ORGANIZATION, MISC, PLACE
- **CLI interface** for command-line usage  
- **HTTP API** with REST endpoint
- **JSON response** format with confidence scores
- **Docker image** available on Docker Hub: [`drzippie/ner-service`](https://hub.docker.com/r/drzippie/ner-service)

## Prerequisites

### System Dependencies

1. **Go 1.19+**
   ```bash
   # macOS
   brew install go
   
   # Ubuntu/Debian
   sudo apt-get install golang-go
   ```

2. **MITIE Library**
   ```bash
   # macOS
   brew install mitie
   
   # Ubuntu/Debian - build from source
   git clone https://github.com/mit-nlp/MITIE.git
   cd MITIE
   make MITIE-models
   sudo make install
   ```

3. **Spanish Language Model**
   The repository includes model download automation:
   ```bash
   make download-model
   ```
   
   This downloads ~450MB from SourceForge. Model files are excluded from git.

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd ner-service-go
   ```

2. Install Go dependencies:
   ```bash
   go mod download
   ```

3. Download the Spanish model:
   ```bash
   make download-model
   ```

4. Build the application:
   ```bash
   make build
   ```

## Usage

### HTTP Server

Start the HTTP server:
```bash
make run-server
# or
./ner-server
```

The server starts on port 8080 by default. You can change the port using the `PORT` environment variable.

#### API Endpoints

**GET /health**
```bash
curl http://localhost:8080/health
# Response: {"status":"healthy","service":"ner-service-go"}
```

**POST /ner**

The endpoint accepts text input in multiple formats:

**Option 1: JSON Input**
```bash
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Juan vive en Madrid y trabaja en Google España."}'
```

**Option 2: URL-encoded Form Data**
```bash
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "text=Juan vive en Madrid y trabaja en Google España."
```

**Option 3: Multipart Form Data**
```bash
curl -X POST http://localhost:8080/ner \
  -F "text=Juan vive en Madrid y trabaja en Google España."
```

**Response** (same for all input methods):
```json
[
  {
    "tag": "PLACE",
    "score": "0.758809",
    "label": "Juan"
  },
  {
    "tag": "PLACE", 
    "score": "1.289719",
    "label": "Madrid"
  },
  {
    "tag": "PLACE",
    "score": "0.733541", 
    "label": "Google España"
  }
]
```

### CLI Interface

Analyze text directly:
```bash
./ner-cli "Juan vive en Madrid y trabaja en Google España."
# Output: Found 3 entities:
# 1. Juan (PLACE) - Score: 0.758809
# 2. Madrid (PLACE) - Score: 1.289719  
# 3. Google España (PLACE) - Score: 0.733541
```

Analyze text from file:
```bash
./ner-cli --file input.txt
```

JSON output:
```bash
./ner-cli --json "Juan vive en Madrid y trabaja en Google España."
# Output: [{"tag":"PLACE","score":"0.758809","label":"Juan"}...]
```

Custom model path:
```bash
./ner-cli --model /path/to/model.dat "Your text here"
```

## Configuration

Environment variables:
- `MITIE_MODEL_PATH`: Path to the MITIE model file (default: `models/ner_model.dat`)
- `PORT`: HTTP server port (default: `8080`)

## Entity Types

The service recognizes the following entity types:
- **PERSON**: People's names
- **LOCATION**: Geographic locations  
- **ORGANIZATION**: Companies, institutions, organizations
- **MISC**: Miscellaneous entities
- **PLACE**: Places and venues

## Development

### Project Structure
```
├── cmd/
│   ├── server/          # HTTP server implementation
│   └── cli/             # CLI implementation  
├── internal/
│   ├── config/          # Configuration management
│   └── ner/             # NER service logic
├── models/              # MITIE model files (downloaded separately)
│   └── README.md        # Model download instructions
├── Makefile            # Build automation
└── .gitignore          # Excludes model files and binaries
```

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
# Build both server and CLI
make build

# Build individually  
make server
make cli
```

### Make Commands
```bash
make build          # Build both server and CLI
make deps           # Download Go dependencies
make download-model # Download Spanish MITIE model
make run-server     # Run server in development
make run-cli        # Run CLI in development  
make clean          # Clean build artifacts
make setup          # Full setup (install + download + build)
```

## Docker Support

### Pre-built Image (Recommended)
Use the official Docker image from [Docker Hub](https://hub.docker.com/r/drzippie/ner-service) with Spanish models included (~539MB):

```bash
# Pull and run from Docker Hub
docker pull drzippie/ner-service:latest
docker run -p 8080:8080 drzippie/ner-service:latest

# Or use specific version
docker run -p 8080:8080 drzippie/ner-service:1.0.0

# With custom port
docker run -p 3000:3000 -e PORT=3000 drzippie/ner-service:latest
```

**Available on Docker Hub**: https://hub.docker.com/r/drzippie/ner-service

### Build Your Own Image
Build Docker image locally with Spanish models included:
```bash
# Build image with models
docker build --platform linux/amd64 -t ner-service-go .

# Run directly
docker run -p 8080:8080 ner-service-go

# Or run with docker-compose
docker-compose up
```

### Custom Build with External Models
For smaller images, mount models as external volume:
```bash
# Download models locally first
make download-model

# Build without models (modify Dockerfile to remove model download)
docker build --platform linux/amd64 -t ner-service-go:lite .

# Run with volume mount
docker run -p 8080:8080 -v ./models:/app/models:ro ner-service-go:lite
```

### Testing the Docker Image
```bash
# Start the official Docker image
docker run -d --name ner-service -p 8080:8080 drzippie/ner-service:latest

# Test API with JSON
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Juan vive en Madrid y trabaja en Google España."}'

# Test API with form data  
curl -X POST http://localhost:8080/ner \
  -d "text=María trabaja en Barcelona para Microsoft."

# Test health endpoint
curl http://localhost:8080/health

# Clean up
docker stop ner-service && docker rm ner-service
```

## License

MIT License