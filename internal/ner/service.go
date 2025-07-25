package ner

import (
	"fmt"
	"strconv"

	"github.com/sbl/ner"
)

type Service struct {
	extractor *ner.Extractor
}

func NewService(modelPath string) (*Service, error) {
	extractor, err := ner.NewExtractor(modelPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create MITIE extractor: %w", err)
	}

	return &Service{
		extractor: extractor,
	}, nil
}

func (s *Service) Close() {
	if s.extractor != nil {
		s.extractor.Free()
	}
}

func (s *Service) ExtractEntities(text string) ([]Entity, error) {
	tokens := ner.Tokenize(text)
	if len(tokens) == 0 {
		return []Entity{}, nil
	}

	entities, err := s.extractor.Extract(tokens)
	if err != nil {
		return nil, fmt.Errorf("failed to extract entities: %w", err)
	}

	result := make([]Entity, len(entities))
	for i, entity := range entities {
		result[i] = Entity{
			Tag:   mapTagToStandardFormat(entity.TagString),
			Score: strconv.FormatFloat(entity.Score, 'f', 6, 64),
			Label: entity.Name,
		}
	}

	return result, nil
}

func mapTagToStandardFormat(mitieTag string) string {
	switch mitieTag {
	case "PERSON":
		return "PERSON"
	case "LOCATION":
		return "LOCATION"
	case "ORGANIZATION":
		return "ORGANIZATION"
	case "MISC":
		return "MISC"
	default:
		return "PLACE"
	}
}