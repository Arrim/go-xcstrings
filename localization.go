package xcstrings

import "encoding/json"

// Localization represents the localization data.
type Localization struct {
	StringUnit *StringUnit `json:"stringUnit,omitempty"`
	Variations *Variation  `json:"variations,omitempty"`
}

// IsPlural checks if the Localization is a plural variation.
func (l *Localization) IsPlural() bool {
	return l.Variations != nil && l.Variations.Plural != nil
}

// IsDevice checks if the Localization is a device variation.
func (l *Localization) IsDevice() bool {
	return l.Variations != nil && l.Variations.Device != nil
}

// AddPlural adds or updates a plural variation for the specified PluralType.
func (l *Localization) AddPlural(pluralType PluralType, loc Localization) {
	if l.Variations == nil {
		l.Variations = &Variation{
			Plural: make(map[PluralType]Localization),
		}
	}
	if l.Variations.Plural == nil {
		l.Variations.Plural = make(map[PluralType]Localization)
	}
	l.Variations.Plural[pluralType] = loc
}

// VariationType represents the type of variation (plural or device).
type VariationType string

// Variation represents either plural or device-specific localizations.
type Variation struct {
	Plural map[PluralType]Localization `json:"plural,omitempty"`
	Device map[DeviceType]Localization `json:"device,omitempty"`
}

// StringUnit represents the state and value of a localized string.
type StringUnit struct {
	State StringState `json:"state"`
	Value string      `json:"value"`
}

// StringState represents the translation state of a string.
type StringState string

// PluralType represents pluralization types.
type PluralType string

// DeviceType represents the device type for variations.
type DeviceType string

// UnmarshalJSON for Variation handles plural and device-specific localizations.
func (v *Variation) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if pluralData, ok := raw["plural"]; ok {
		var plural map[PluralType]Localization
		if err := json.Unmarshal(pluralData, &plural); err != nil {
			return err
		}
		v.Plural = plural
	}

	if deviceData, ok := raw["device"]; ok {
		var device map[DeviceType]Localization
		if err := json.Unmarshal(deviceData, &device); err != nil {
			return err
		}
		v.Device = device
	}

	return nil
}

// MarshalJSON for Variation serializes plural and device-specific localizations.
func (v *Variation) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	if v.Plural != nil {
		data["plural"] = v.Plural
	}
	if v.Device != nil {
		data["device"] = v.Device
	}
	return json.Marshal(data)
}

// UnmarshalJSON for Localization handles stringUnit and variations.
func (l *Localization) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if stringUnitData, ok := raw["stringUnit"]; ok {
		var stringUnit StringUnit
		if err := json.Unmarshal(stringUnitData, &stringUnit); err != nil {
			return err
		}
		l.StringUnit = &stringUnit
	}

	if variationsData, ok := raw["variations"]; ok {
		var variations Variation
		if err := json.Unmarshal(variationsData, &variations); err != nil {
			return err
		}
		l.Variations = &variations
	}

	return nil
}

// MarshalJSON for Localization serializes stringUnit and variations.
func (l *Localization) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	if l.StringUnit != nil {
		data["stringUnit"] = l.StringUnit
	}
	if l.Variations != nil {
		data["variations"] = l.Variations
	}
	return json.Marshal(data)
}
