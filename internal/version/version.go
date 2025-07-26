package version

// Version represents the current version of the NER service
const Version = "1.0.1"

// BuildInfo provides build information
type BuildInfo struct {
	Version string `json:"version"`
	Service string `json:"service"`
}

// GetBuildInfo returns build information for the service
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version: Version,
		Service: "ner-service-go",
	}
}
