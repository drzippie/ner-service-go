# NER Service Go

A Named Entity Recognition (NER) service for Spanish text using MITIE (MIT Information Extraction library). This service provides both CLI and HTTP API interfaces for extracting named entities from Spanish text.

## ✅ Status: Fully Working

The application is **fully functional and ready to use**. Both CLI and HTTP server modes are operational with the Spanish language model.

## 🚀 Quick Start

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

- Named Entity Recognition for Spanish text
- Support for entity types: PERSON, LOCATION, ORGANIZATION, MISC, PLACE
- CLI interface for command-line usage
- HTTP API with REST endpoint
- JSON response format with entity scores

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

Request:
```json
{
  "text": "Juan vive en Madrid y trabaja en Google España."
}
```

Response:
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

#### Example with curl:
```bash
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Juan vive en Madrid y trabaja en Google España."}'
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

### Option 1: Image with Models Included (Recommended)
Build Docker image with Spanish models included (~539MB):
```bash
# Build image with models
docker build --platform linux/amd64 -t ner-service-go .

# Run directly
docker run -p 8080:8080 ner-service-go

# Or run with docker-compose
docker-compose up
```

### Option 2: Lightweight Image + External Models
For smaller images, mount models as external volume:
```bash
# Download models locally first
make download-model

# Build without models (modify Dockerfile to remove model download)
docker build --platform linux/amd64 -t ner-service-go:lite .

# Run with volume mount
docker run -p 8080:8080 -v ./models:/app/models:ro ner-service-go:lite
```

### Testing
```bash
# Test API
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Juan vive en Madrid y trabaja en Google España."}'

# Test health
curl http://localhost:8080/health
```

## License

MIT License