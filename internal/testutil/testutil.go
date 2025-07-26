package testutil

// SpanishTestTexts contains common Spanish test texts for NER testing
var SpanishTestTexts = struct {
	PersonLocation string
	Organization   string
	Mixed          string
	Complex        string
	NoEntities     string
	Empty          string
}{
	PersonLocation: "María García vive en Madrid",
	Organization:   "Trabajo en Microsoft España",
	Mixed:          "Pedro Sánchez visitó Barcelona para reunirse con representantes de Telefónica",
	Complex:        "El presidente del Real Madrid, Florentino Pérez, se reunió con Karim Benzema en el Santiago Bernabéu",
	NoEntities:     "El día está muy soleado y hace calor",
	Empty:          "",
}

// ExpectedEntityTypes contains common entity types for validation
var ExpectedEntityTypes = []string{
	"PERSON",
	"LOCATION",
	"ORGANIZATION",
	"MISC",
}
