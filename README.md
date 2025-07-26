# NER Service Go

[![Version](https://img.shields.io/badge/Version-1.0.1-brightgreen)](https://github.com/your-repo/ner-service-go/releases)
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
  -d '{"text": "Juan Pérez vive en Madrid y trabaja en Google España desde 2020."}'
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
./ner-cli "Juan Pérez vive en Madrid y trabaja en Google España desde 2020."

# 5. Start HTTP server
./ner-server

# 6. Test API endpoint  
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Juan Pérez vive en Madrid y trabaja en Google España desde 2020."}'
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

**Option 1: JSON Input (Recommended)**
```bash
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "María García trabaja en Barcelona para Microsoft España y vive cerca del Parque Güell."}'
```

**Option 2: URL-encoded Form Data**
```bash
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "text=El presidente Pedro Sánchez visitó el Museo del Prado en Madrid."
```

**Option 3: Multipart Form Data**
```bash
curl -X POST http://localhost:8080/ner \
  -F "text=La empresa Telefónica tiene su sede en Madrid, España."
```

**Response Examples:**

*For person and organization text:*
```json
[
  {
    "tag": "PERSON",
    "score": "0.892",
    "label": "María García"
  },
  {
    "tag": "LOCATION", 
    "score": "1.456",
    "label": "Barcelona"
  },
  {
    "tag": "ORGANIZATION",
    "score": "1.234", 
    "label": "Microsoft España"
  },
  {
    "tag": "LOCATION",
    "score": "0.987",
    "label": "Parque Güell"
  }
]
```

*For political and cultural entities:*
```json
[
  {
    "tag": "PERSON",
    "score": "1.567",
    "label": "Pedro Sánchez"
  },
  {
    "tag": "ORGANIZATION",
    "score": "1.123",
    "label": "Museo del Prado"
  },
  {
    "tag": "LOCATION",
    "score": "1.789",
    "label": "Madrid"
  }
]
```

### CLI Interface

**Basic text analysis:**
```bash
./ner-cli "María García trabaja en Barcelona para Microsoft España."
# Output: Found 3 entities:
# 1. María García (PERSON) - Score: 0.892
# 2. Barcelona (LOCATION) - Score: 1.456
# 3. Microsoft España (ORGANIZATION) - Score: 1.234
```

**Analyze text from file:**
```bash
# Create a sample file
echo "El director de Telefónica, José María Álvarez-Pallete, anunció la expansión en Valencia." > sample.txt

./ner-cli --file sample.txt
# Output: Found 4 entities:
# 1. Telefónica (ORGANIZATION) - Score: 1.567
# 2. José María Álvarez-Pallete (PERSON) - Score: 1.234
# 3. Valencia (LOCATION) - Score: 1.123
```

**JSON output for integration:**
```bash
./ner-cli --json "Pedro Sánchez visitó el Congreso en Madrid."
# Output: [{"tag":"PERSON","score":"1.567","label":"Pedro Sánchez"},{"tag":"ORGANIZATION","score":"1.123","label":"Congreso"},{"tag":"LOCATION","score":"1.789","label":"Madrid"}]
```

**Custom model path:**
```bash
./ner-cli --model /custom/path/model.dat "Antonio Banderas nació en Málaga."
```

**Complex entity examples:**
```bash
# Sports entities
./ner-cli "Lionel Messi jugó en el Barcelona y ahora está en el PSG de París."

# Business entities  
./ner-cli "El CEO de Zara, Amancio Ortega, vive en La Coruña, Galicia."

# Cultural entities
./ner-cli "García Lorca escribió Bodas de Sangre en Granada, Andalucía."
```

## Configuration

Environment variables:
- `MITIE_MODEL_PATH`: Path to the MITIE model file (default: `models/ner_model.dat`)
- `PORT`: HTTP server port (default: `8080`)

## Entity Types

The service recognizes the following entity types with improved mapping (v1.0.1):
- **PERSON**: People's names (e.g., "María García", "Pedro Sánchez", "Rafael Nadal")
- **LOCATION**: Geographic locations (e.g., "Madrid", "Barcelona", "Valencia")
- **ORGANIZATION**: Companies, institutions, organizations (e.g., "Microsoft España", "Telefónica", "Banco Santander")
- **MISC**: Miscellaneous entities (dates, events, etc.)
- **PLACE**: Places and venues (e.g., "Parque Güell", "Roland Garros")

### Entity Mapping Improvements (v1.0.1)

Recent updates have improved entity classification accuracy:
- Better distinction between PERSON and ORGANIZATION entities
- Improved recognition of Spanish names and surnames
- Enhanced location detection for Spanish geography
- More accurate organization identification for Spanish companies

## Real-World Use Cases & Examples

**News Article Processing:**
```bash
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "El Real Madrid, dirigido por Carlo Ancelotti, se enfrentará al Barcelona en el Santiago Bernabéu el próximo domingo. Luka Modrić y Pedri serán las estrellas del encuentro."}'
```

**Business Intelligence:**
```bash
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Inditex, la empresa de Amancio Ortega con sede en La Coruña, anunció la apertura de nuevas tiendas Zara en México y Argentina durante el próximo trimestre."}'
```

**Social Media Monitoring:**
```bash
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "@mariaperez menciona que visitará el Museo Reina Sofía en Madrid este fin de semana junto con @carlosgarcia para ver las obras de Dalí y Miró."}'
```

**Academic Research:**
```bash
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "La investigación realizada por la Universidad Complutense de Madrid bajo la dirección de la doctora Ana Martínez ha sido publicada en la revista Nature."}'
```

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

**Quick start with comprehensive examples:**
```bash
# Start the official Docker image
docker run -d --name ner-service -p 8080:8080 drzippie/ner-service:latest

# Test health endpoint
curl http://localhost:8080/health
# Expected: {"status":"healthy","service":"ner-service-go"}

# Test with Spanish person and location
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Rafael Nadal ganó el torneo en Roland Garros, París."}'

# Test with business entities
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "El presidente de Telefónica, José María Álvarez-Pallete, anunció inversiones en Madrid."}'

# Test with cultural entities
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Federico García Lorca nació en Fuente Vaqueros, Granada."}'

# Test form data input
curl -X POST http://localhost:8080/ner \
  -d "text=Antoni Gaudí diseñó la Sagrada Familia en Barcelona."

# Test multipart form input
curl -X POST http://localhost:8080/ner \
  -F "text=Pablo Picasso pintó el Guernica durante la Guerra Civil Española."

# Performance test with longer text
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Durante la reunión en el Palacio de la Moncloa, el presidente Pedro Sánchez se reunió con los directores de Banco Santander, BBVA y CaixaBank para discutir la situación económica de España tras la crisis del COVID-19."}'

# Clean up
docker stop ner-service && docker rm ner-service
```

**Expected response format:**
All API calls return a JSON array with entities and confidence scores:
```json
[
  {
    "tag": "PERSON",
    "score": "1.234",
    "label": "Rafael Nadal"
  },
  {
    "tag": "LOCATION",
    "score": "0.987",
    "label": "Roland Garros"
  },
  {
    "tag": "LOCATION",
    "score": "1.456",
    "label": "París"
  }
]
```

## License

MIT License