package resume

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/goccy/go-yaml"
)

func TestJsonIsYaml(t *testing.T) {
	yamlFile := filepath.Join("..", "schema.yaml")
	jsonFile := filepath.Join("..", "schema.json")

	yamlBytes, err := os.ReadFile(yamlFile)
	if err != nil {
		t.Fatalf("failed to read YAML schema: %v", err)
	}

	jsonBytes, err := os.ReadFile(jsonFile)
	if err != nil {
		t.Fatalf("failed to read JSON schema: %v", err)
	}

	var fromYaml Resume
	if err := yaml.Unmarshal(yamlBytes, &fromYaml); err != nil {
		t.Fatalf("failed to unmarshal YAML: %v", err)
	}

	var fromJson Resume
	if err := yaml.Unmarshal(jsonBytes, &fromJson); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if !reflect.DeepEqual(fromYaml, fromJson) {
		t.Errorf("Resume structs do not match:\nYAML: %+v\nJSON: %+v", fromYaml, fromJson)
	}
}
