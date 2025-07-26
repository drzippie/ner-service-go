# Testing Guide

This document describes the simplified testing approach for the NER Service Go project.

## Test Structure

The project includes focused unit tests that don't require CGO dependencies or MITIE model downloads:

### Unit Tests

**Location**: `internal/*/` directories  
**Files**: `*_test.go`

- **Configuration Tests** (`internal/config/config_test.go`)
  - Environment variable handling
  - Default value validation
  - Configuration loading

- **Types Tests** (`internal/types/types_test.go`)
  - JSON serialization/deserialization
  - Data structure validation
  - Entity validation logic

- **Test Utilities Tests** (`internal/testutil/testutil_test.go`)
  - Spanish test text validation
  - Entity type constants verification
  - Spanish character handling

## Test Utilities

**Location**: `internal/testutil/testutil.go`

Provides common utilities for testing:

- **Spanish Test Texts**: Predefined Spanish test cases for different entity types
- **Expected Entity Types**: Standard entity type constants

## Running Tests

### Prerequisites

**Install Dependencies**:
```bash
make deps
```

No MITIE model download required for unit tests!

### Test Commands

#### Run All Tests
```bash
make test
```

#### Unit Tests Only (same as above)
```bash
make test-unit
```

#### Test Coverage Report
```bash
make test-coverage
```
Generates `coverage.html` with detailed coverage report.

#### Verbose Testing
```bash
make test-verbose
```

#### Direct Go Commands
```bash
# All tests
go test -v ./internal/config ./internal/testutil ./internal/types

# Specific package
go test -v ./internal/config

# With coverage
go test -v -coverprofile=coverage.out ./internal/config ./internal/testutil ./internal/types
```

## Test Categories by Function

### Configuration Tests

Validate configuration management:

- **Environment variables**: `MITIE_MODEL_PATH`, `PORT`
- **Default values**: Fallback configuration
- **Partial configuration**: Mixed env vars and defaults

### Data Structure Tests

Validate JSON serialization and entity structures:

- **Entity serialization**: JSON marshaling/unmarshaling
- **Request/Response structures**: API data types
- **Validation logic**: Required field checks
- **Edge cases**: Empty values, invalid inputs

### Test Utility Tests

Validate testing infrastructure:

- **Spanish test texts**: Predefined test cases
- **Entity type constants**: Standard type definitions
- **Character handling**: Spanish language support

## Test Data

### Spanish Test Texts

The test suite includes carefully selected Spanish texts:

```go
// Person and Location
"María García vive en Madrid"

// Organization
"Trabajo en Microsoft España"

// Complex Mixed
"Pedro Sánchez visitó Barcelona para reunirse con representantes de Telefónica"

// No Entities
"El día está muy soleado y hace calor"
```

### Entity Type Validation

Tests verify standard entity types:

- `"PERSON"` - People's names
- `"LOCATION"` - Geographic locations
- `"ORGANIZATION"` - Companies, institutions
- `"MISC"` - Miscellaneous entities

## GitHub Actions Integration

### Automated Testing

The project includes GitHub Actions workflows:

1. **Main Tests** (`.github/workflows/test.yml`): Runs on push/PR
   - Unit tests with coverage
   - Go linting and formatting checks
   - Security scanning

2. **Pull Request Checks** (`.github/workflows/pr.yml`): Quick validation
   - Fast unit tests
   - Code formatting verification
   - Docker linting

3. **Release** (`.github/workflows/release.yml`): On version tags
   - Full test suite before release
   - Automated GitHub release creation

### Example CI Configuration

```yaml
- name: Run unit tests
  run: |
    go test -v ./internal/config ./internal/testutil ./internal/types

- name: Run tests with coverage
  run: |
    go test -v -coverprofile=coverage.out ./internal/config ./internal/testutil ./internal/types
    go tool cover -func=coverage.out
```

## Troubleshooting

### Common Issues

1. **Go Version**: Ensure Go 1.24+ is installed
2. **Dependencies**: Run `go mod download` before testing
3. **Formatting**: Use `go fmt ./...` to fix formatting issues

### Debug Mode

For detailed test output:

```bash
go test -v -count=1 ./internal/config ./internal/testutil ./internal/types
```

## Contributing

When adding new tests:

1. **Follow naming conventions**: `Test*` for tests
2. **Use clear test names**: Describe what you're testing
3. **Test edge cases**: Empty values, invalid inputs
4. **Keep tests simple**: No external dependencies
5. **Validate structure**: Check all required fields