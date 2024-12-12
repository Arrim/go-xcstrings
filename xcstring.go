package xcstrings

import "encoding/json"

// XCString represents a single localization entry.
type XCString struct {
	Comment         *string                   `json:"comment,omitempty"`
	Key             *string                   `json:"key,omitempty"`
	Localizations   map[Language]Localization `json:"localizations"`
	ExtractionState ExtractionState           `json:"extractionState,omitempty"`
	ShouldTranslate bool                      `json:"shouldTranslate"`
}

// NewXCString creates a new XCString instance.
func NewXCString(key string, comment *string, extractionState ExtractionState, shouldTranslate bool, localizations map[Language]Localization) *XCString {
	return &XCString{
		Comment:         comment,
		Key:             &key,
		Localizations:   localizations,
		ExtractionState: extractionState,
		ShouldTranslate: shouldTranslate,
	}
}

// GetLocalization retrieves a Localization by language from the XCString structure.
// Returns the Localization and a boolean indicating if the language was found.
func (x *XCString) GetLocalization(language Language) (*Localization, bool) {
	loc, found := x.Localizations[language]
	if !found {
		return nil, false
	}
	return &loc, true
}

// GetComment retrieves the comment for the XCString.
// Returns the comment as a string. If no comment is set, returns an empty string.
func (x *XCString) GetComment() string {
	if x.Comment != nil {
		return *x.Comment
	}
	return ""
}

// ExtractionState represents the state of the string extraction.
type ExtractionState string

// Language represents a language code.
type Language string

// UnmarshalJSON for XCString converts localizations into the map of Language.
func (x *XCString) UnmarshalJSON(data []byte) error {
	type Alias XCString
	aux := &struct {
		Localizations map[string]Localization `json:"localizations"`
		*Alias
	}{
		Alias: (*Alias)(x),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	x.Localizations = make(map[Language]Localization)
	for key, loc := range aux.Localizations {
		x.Localizations[Language(key)] = loc
	}
	return nil
}

// MarshalJSON for XCString converts the map of Language to localizations.
func (x *XCString) MarshalJSON() ([]byte, error) {
	type Alias XCString
	aux := &struct {
		Localizations map[string]Localization `json:"localizations"`
		*Alias
	}{
		Alias: (*Alias)(x),
		Localizations: func() map[string]Localization {
			m := make(map[string]Localization)
			for key, loc := range x.Localizations {
				m[string(key)] = loc
			}
			return m
		}(),
	}
	return json.Marshal(aux)
}
