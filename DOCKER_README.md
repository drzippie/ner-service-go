# NER Service Go

A Named Entity Recognition (NER) service for Spanish text using MITIE (MIT Information Extraction library). This Docker image provides a production-ready HTTP API for extracting named entities from Spanish text.

## 🚀 Quick Start

```bash
# Pull and run the container
docker run -p 8080:8080 drzippie/ner-service:latest

# Test the API
curl -X POST http://localhost:8080/ner \
  -H "Content-Type: application/json" \
  -d '{"text": "Juan vive en Madrid y trabaja en Google España."}'
```

## 📋 Features

- **Spanish Text NER**: Specialized for Spanish language entity recognition
- **Multiple Entity Types**: PERSON, LOCATION, ORGANIZATION, MISC, PLACE
- **Flexible Input**: JSON, form data, or multipart form support
- **Production Ready**: Multi-stage build, security hardened, health checks
- **Lightweight**: ~539MB image with models included

## 🔧 API Endpoints

### Health Check
```bash
GET http://localhost:8080/health
```

### Named Entity Recognition
```bash
POST http://localhost:8080/ner
Content-Type: application/json

{
  "text": "María trabaja en Barcelona para Microsoft."
}
```

**Response:**
```json
[
  {
    "tag": "PERSON",
    "score": "0.892341",
    "label": "María"
  },
  {
    "tag": "LOCATION", 
    "score": "1.234567",
    "label": "Barcelona"
  },
  {
    "tag": "ORGANIZATION",
    "score": "0.987654",
    "label": "Microsoft"
  }
]
```

## 🐳 Docker Usage

### Basic Usage
```bash
# Run with default settings
docker run -p 8080:8080 drzippie/ner-service:latest

# Run with custom port
docker run -p 3000:3000 -e PORT=3000 drzippie/ner-service:latest

# Run with docker-compose
curl -O https://raw.githubusercontent.com/your-repo/ner-service-go/main/docker-compose.yml
docker-compose up
```

### Environment Variables
- `PORT`: HTTP server port (default: 8080)
- `MITIE_MODEL_PATH`: Path to MITIE model (default: /app/models/ner_model.dat)
- `GIN_MODE`: Gin framework mode (default: release)

## 🏗️ Architecture

- **Base**: Debian Bullseye Slim
- **Runtime**: Go 1.24+ with CGO enabled
- **NER Engine**: MITIE v0.7 with Spanish models
- **Security**: Non-root user, minimal attack surface
- **Size**: ~539MB (includes Spanish language model)

## 📊 Performance

- **Memory**: ~512MB minimum, 2GB recommended
- **CPU**: 0.5-1.0 cores recommended
- **Startup**: ~10-15 seconds (model loading)
- **Throughput**: Varies by text length and complexity

## 🔒 Security Features

- Non-root container execution
- Minimal runtime dependencies
- Built-in health checks
- Resource limits supported
- Security-hardened base image

## 📝 Source Code

- **GitHub**: [ner-service-go](https://github.com/your-repo/ner-service-go)
- **Language**: Go 1.24+
- **License**: MIT
- **Framework**: Gin (HTTP), Cobra (CLI)

## 🏷️ Available Tags

- `latest` - Latest stable release
- `1.0.0` - Specific version release
- Future releases will follow semantic versioning

## 💡 Use Cases

- **Content Analysis**: Extract entities from Spanish articles, blogs, social media
- **Data Processing**: Batch processing of Spanish text documents  
- **Search Enhancement**: Improve search with entity-based indexing
- **Compliance**: Identify names and organizations in documents
- **Research**: Academic and commercial NLP research projects

Built with ❤️ for the Spanish NLP community.