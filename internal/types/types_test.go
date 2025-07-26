package types

import (
	"encoding/json"
	"strings"
	"testing"
)

// Entity represents a named entity with its metadata
type Entity struct {
	Tag   string `json:"tag"`
	Score string `json:"score"`
	Label string `json:"label"`
}

// ExtractRequest represents the request structure for entity extraction
type ExtractRequest struct {
	Text string `json:"text"`
}

// ExtractResponse represents the response structure for entity extraction
type ExtractResponse struct {
	Entities []Entity `json:"entities"`
}

func TestEntity_JSONSerialization(t *testing.T) {
	entity := Entity{
		Tag:   "PERSON",
		Score: "0.95",
		Label: "María García",
	}

	// Test marshaling
	data, err := json.Marshal(entity)
	if err != nil {
		t.Errorf("Failed to marshal entity: %v", err)
	}

	expected := `{"tag":"PERSON","score":"0.95","label":"María García"}`
	if string(data) != expected {
		t.Errorf("Expected JSON %s, but got %s", expected, string(data))
	}

	// Test unmarshaling
	var unmarshaledEntity Entity
	err = json.Unmarshal(data, &unmarshaledEntity)
	if err != nil {
		t.Errorf("Failed to unmarshal entity: %v", err)
	}

	if unmarshaledEntity != entity {
		t.Errorf("Unmarshaled entity %+v doesn't match original %+v", unmarshaledEntity, entity)
	}
}

func TestExtractRequest_JSONSerialization(t *testing.T) {
	request := ExtractRequest{
		Text: "Pedro Sánchez vive en Madrid",
	}

	// Test marshaling
	data, err := json.Marshal(request)
	if err != nil {
		t.Errorf("Failed to marshal request: %v", err)
	}

	expected := `{"text":"Pedro Sánchez vive en Madrid"}`
	if string(data) != expected {
		t.Errorf("Expected JSON %s, but got %s", expected, string(data))
	}

	// Test unmarshaling
	var unmarshaledRequest ExtractRequest
	err = json.Unmarshal(data, &unmarshaledRequest)
	if err != nil {
		t.Errorf("Failed to unmarshal request: %v", err)
	}

	if unmarshaledRequest != request {
		t.Errorf("Unmarshaled request %+v doesn't match original %+v", unmarshaledRequest, request)
	}
}

func TestExtractResponse_JSONSerialization(t *testing.T) {
	response := ExtractResponse{
		Entities: []Entity{
			{
				Tag:   "PERSON",
				Score: "0.95",
				Label: "Pedro Sánchez",
			},
			{
				Tag:   "LOCATION",
				Score: "0.89",
				Label: "Madrid",
			},
		},
	}

	// Test marshaling
	data, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal response: %v", err)
	}

	// Test unmarshaling
	var unmarshaledResponse ExtractResponse
	err = json.Unmarshal(data, &unmarshaledResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if len(unmarshaledResponse.Entities) != len(response.Entities) {
		t.Errorf("Expected %d entities, but got %d", len(response.Entities), len(unmarshaledResponse.Entities))
	}

	for i, entity := range unmarshaledResponse.Entities {
		if entity != response.Entities[i] {
			t.Errorf("Entity %d: %+v doesn't match original %+v", i, entity, response.Entities[i])
		}
	}
}

func TestEntity_EmptyValues(t *testing.T) {
	entity := Entity{}

	data, err := json.Marshal(entity)
	if err != nil {
		t.Errorf("Failed to marshal empty entity: %v", err)
	}

	expected := `{"tag":"","score":"","label":""}`
	if string(data) != expected {
		t.Errorf("Expected JSON %s, but got %s", expected, string(data))
	}
}

func TestExtractRequest_EmptyText(t *testing.T) {
	request := ExtractRequest{Text: ""}

	data, err := json.Marshal(request)
	if err != nil {
		t.Errorf("Failed to marshal request with empty text: %v", err)
	}

	expected := `{"text":""}`
	if string(data) != expected {
		t.Errorf("Expected JSON %s, but got %s", expected, string(data))
	}
}

func TestExtractResponse_EmptyEntities(t *testing.T) {
	response := ExtractResponse{Entities: []Entity{}}

	data, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal response with empty entities: %v", err)
	}

	expected := `{"entities":[]}`
	if string(data) != expected {
		t.Errorf("Expected JSON %s, but got %s", expected, string(data))
	}
}

func TestEntityValidation(t *testing.T) {
	tests := []struct {
		name    string
		entity  Entity
		isValid bool
	}{
		{
			name: "Valid PERSON entity",
			entity: Entity{
				Tag:   "PERSON",
				Score: "0.95",
				Label: "María García",
			},
			isValid: true,
		},
		{
			name: "Valid LOCATION entity",
			entity: Entity{
				Tag:   "LOCATION",
				Score: "0.89",
				Label: "Madrid",
			},
			isValid: true,
		},
		{
			name: "Valid ORGANIZATION entity",
			entity: Entity{
				Tag:   "ORGANIZATION",
				Score: "0.92",
				Label: "Microsoft España",
			},
			isValid: true,
		},
		{
			name: "Empty tag",
			entity: Entity{
				Tag:   "",
				Score: "0.95",
				Label: "María García",
			},
			isValid: false,
		},
		{
			name: "Empty label",
			entity: Entity{
				Tag:   "PERSON",
				Score: "0.95",
				Label: "",
			},
			isValid: false,
		},
		{
			name: "Empty score",
			entity: Entity{
				Tag:   "PERSON",
				Score: "",
				Label: "María García",
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasRequiredFields := tt.entity.Tag != "" && tt.entity.Label != "" && tt.entity.Score != ""

			if tt.isValid && !hasRequiredFields {
				t.Errorf("Entity should be valid but is missing required fields: %+v", tt.entity)
			}

			if !tt.isValid && hasRequiredFields {
				t.Errorf("Entity should be invalid but has all required fields: %+v", tt.entity)
			}
		})
	}
}

func TestExtractRequest_TextValidation(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		isValid bool
	}{
		{"Valid Spanish text", "María García vive en Madrid", true},
		{"Empty text", "", false},
		{"Whitespace only", "   ", false},
		{"Complex Spanish text", "El presidente Pedro Sánchez visitó Barcelona", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := ExtractRequest{Text: tt.text}

			hasValidText := len(request.Text) > 0 && len(strings.TrimSpace(request.Text)) > 0

			if tt.isValid && !hasValidText {
				t.Errorf("Request should be valid but text is invalid: '%s'", tt.text)
			}

			if !tt.isValid && hasValidText {
				t.Errorf("Request should be invalid but text is valid: '%s'", tt.text)
			}
		})
	}
}
