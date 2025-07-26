package testutil

import (
	"strings"
	"testing"
)

func TestSpanishTestTexts(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{"PersonLocation", SpanishTestTexts.PersonLocation},
		{"Organization", SpanishTestTexts.Organization},
		{"Mixed", SpanishTestTexts.Mixed},
		{"Complex", SpanishTestTexts.Complex},
		{"NoEntities", SpanishTestTexts.NoEntities},
		{"Empty", SpanishTestTexts.Empty},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test that text values are defined (except Empty which should be empty)
			if tt.name == "Empty" {
				if tt.text != "" {
					t.Errorf("Expected empty text for Empty case, but got: %s", tt.text)
				}
			} else {
				if tt.text == "" {
					t.Errorf("Expected non-empty text for %s, but got empty string", tt.name)
				}
				
				// Test that Spanish text contains expected patterns
				if tt.name == "PersonLocation" && !containsSpanishPattern(tt.text) {
					t.Errorf("PersonLocation text should contain Spanish names or places: %s", tt.text)
				}
				
				if tt.name == "Organization" && !containsSpanishPattern(tt.text) {
					t.Errorf("Organization text should contain Spanish organization names: %s", tt.text)
				}
			}
		})
	}
}

func TestExpectedEntityTypes(t *testing.T) {
	expectedTypes := []string{"PERSON", "LOCATION", "ORGANIZATION", "MISC"}
	
	if len(ExpectedEntityTypes) != len(expectedTypes) {
		t.Errorf("Expected %d entity types, but got %d", len(expectedTypes), len(ExpectedEntityTypes))
	}
	
	for i, expectedType := range expectedTypes {
		if i >= len(ExpectedEntityTypes) {
			t.Errorf("Missing expected entity type: %s", expectedType)
			continue
		}
		
		if ExpectedEntityTypes[i] != expectedType {
			t.Errorf("Expected entity type %s at index %d, but got %s", 
				expectedType, i, ExpectedEntityTypes[i])
		}
	}
}

func TestSpanishTestTexts_Content(t *testing.T) {
	// Test that PersonLocation contains expected Spanish elements
	if !containsAny(SpanishTestTexts.PersonLocation, []string{"García", "Madrid", "María"}) {
		t.Errorf("PersonLocation should contain typical Spanish names/places")
	}
	
	// Test that Organization contains business-related terms
	if !containsAny(SpanishTestTexts.Organization, []string{"Microsoft", "Trabajo", "España"}) {
		t.Errorf("Organization should contain business or work-related terms")
	}
	
	// Test that Mixed contains multiple entity indicators
	mixedText := SpanishTestTexts.Mixed
	hasPersonIndicator := containsAny(mixedText, []string{"Pedro", "Sánchez"})
	hasLocationIndicator := containsAny(mixedText, []string{"Barcelona"})
	hasOrgIndicator := containsAny(mixedText, []string{"Telefónica"})
	
	if !hasPersonIndicator || !hasLocationIndicator || !hasOrgIndicator {
		t.Errorf("Mixed text should contain person, location, and organization indicators")
	}
	
	// Test that NoEntities doesn't contain obvious entity patterns
	if containsAny(SpanishTestTexts.NoEntities, []string{"García", "Madrid", "Microsoft", "Sánchez"}) {
		t.Errorf("NoEntities text should not contain obvious entity names")
	}
}

func TestEntityTypeConstants(t *testing.T) {
	// Test that all expected entity types are valid strings
	for i, entityType := range ExpectedEntityTypes {
		if entityType == "" {
			t.Errorf("Entity type at index %d is empty", i)
		}
		
		if len(entityType) < 3 {
			t.Errorf("Entity type '%s' seems too short", entityType)
		}
		
		// Test that entity types are uppercase
		if entityType != strings.ToUpper(entityType) {
			t.Errorf("Entity type '%s' should be uppercase", entityType)
		}
	}
}

func TestSpanishCharacterHandling(t *testing.T) {
	// Test that Spanish texts contain proper Spanish characters
	spanishTexts := []string{
		SpanishTestTexts.PersonLocation,
		SpanishTestTexts.Organization,
		SpanishTestTexts.Mixed,
		SpanishTestTexts.Complex,
	}
	
	for i, text := range spanishTexts {
		if text == "" {
			continue
		}
		
		// Check for Spanish characters or common Spanish words
		hasSpanishElements := containsAny(text, []string{
			"ñ", "á", "é", "í", "ó", "ú", "ü", // Spanish characters
			"en", "el", "la", "de", "del", "y", "para", // Spanish articles/prepositions
		})
		
		if !hasSpanishElements {
			t.Errorf("Spanish text %d should contain Spanish language elements: %s", i, text)
		}
	}
}

// Helper functions
func containsSpanishPattern(text string) bool {
	spanishPatterns := []string{
		"García", "Sánchez", "Madrid", "Barcelona", "España", 
		"Microsoft", "Telefónica", "Pedro", "María",
	}
	return containsAny(text, spanishPatterns)
}

func containsAny(text string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(text, pattern) {
			return true
		}
	}
	return false
}