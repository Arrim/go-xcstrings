package xcstrings

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// XCStrings represents the main structure of the localization file.
type XCStrings struct {
	Version        string              `json:"version"`
	SourceLanguage Language            `json:"sourceLanguage"`
	Strings        map[string]XCString `json:"strings"`
}

// GetXCString retrieves an XCString by key from the XCStrings structure.
// Returns the XCString and a boolean indicating if the key was found.
func (x *XCStrings) GetXCString(key string) (*XCString, bool) {
	entry, found := x.Strings[key]
	if !found {
		return nil, false
	}
	return &entry, true
}

// ReadXCStrings reads the .xcstrings file
func ReadXCStrings(filePath string) (*XCStrings, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var xcstrings XCStrings
	if err := json.Unmarshal(byteValue, &xcstrings); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &xcstrings, nil
}

// WriteXCStrings writes changes to the .xcstrings file
func WriteXCStrings(filePath string, xcstrings *XCStrings) error {
	byteValue, err := json.MarshalIndent(xcstrings, "", "  ")
	if err != nil {
		return fmt.Errorf("error serializing JSON: %v", err)
	}

	if err := os.WriteFile(filePath, byteValue, 0644); err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}
