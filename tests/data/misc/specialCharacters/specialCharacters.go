// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"
import "reflect"

type License string

const LicenseGPL30 License = "GPL-3.0"
const LicenseMIT License = "MIT"

type License_1 string

const License_1_GPL30 License_1 = "GPL-3.0+"
const License_1_MIT License_1 = "MIT+"

type SpecialCharacters struct {
	// PlainLicenses corresponds to the JSON schema field "plainLicenses".
	PlainLicenses *SpecialCharactersPlainLicenses `json:"plainLicenses,omitempty" yaml:"plainLicenses,omitempty" mapstructure:"plainLicenses,omitempty"`

	// PlainLicensesRef corresponds to the JSON schema field "plainLicensesRef".
	PlainLicensesRef []License `json:"plainLicensesRef,omitempty" yaml:"plainLicensesRef,omitempty" mapstructure:"plainLicensesRef,omitempty"`

	// PlusLicenses corresponds to the JSON schema field "plusLicenses".
	PlusLicenses *SpecialCharactersPlusLicenses `json:"plusLicenses,omitempty" yaml:"plusLicenses,omitempty" mapstructure:"plusLicenses,omitempty"`

	// PlusLicensesRef corresponds to the JSON schema field "plusLicensesRef".
	PlusLicensesRef []License_1 `json:"plusLicensesRef,omitempty" yaml:"plusLicensesRef,omitempty" mapstructure:"plusLicensesRef,omitempty"`
}

type SpecialCharactersPlainLicenses string

const SpecialCharactersPlainLicensesGPL30 SpecialCharactersPlainLicenses = "GPL-3.0"

// UnmarshalJSON implements json.Unmarshaler.
func (j *License) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_License {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_License, v)
	}
	*j = License(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *License_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_License_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_License_1, v)
	}
	*j = License_1(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *License) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_License {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_License, v)
	}
	*j = License(v)
	return nil
}

const SpecialCharactersPlainLicensesMIT SpecialCharactersPlainLicenses = "MIT"

var enumValues_SpecialCharactersPlainLicenses = []interface{}{
	"GPL-3.0",
	"MIT",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SpecialCharactersPlainLicenses) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SpecialCharactersPlainLicenses {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SpecialCharactersPlainLicenses, v)
	}
	*j = SpecialCharactersPlainLicenses(v)
	return nil
}

type SpecialCharactersPlusLicenses string

var enumValues_License_1 = []interface{}{
	"GPL-3.0+",
	"MIT+",
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *License_1) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_License_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_License_1, v)
	}
	*j = License_1(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *SpecialCharactersPlainLicenses) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SpecialCharactersPlainLicenses {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SpecialCharactersPlainLicenses, v)
	}
	*j = SpecialCharactersPlainLicenses(v)
	return nil
}

var enumValues_SpecialCharactersPlusLicenses = []interface{}{
	"GPL-3.0+",
	"MIT+",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SpecialCharactersPlusLicenses) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SpecialCharactersPlusLicenses {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SpecialCharactersPlusLicenses, v)
	}
	*j = SpecialCharactersPlusLicenses(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *SpecialCharactersPlusLicenses) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SpecialCharactersPlusLicenses {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SpecialCharactersPlusLicenses, v)
	}
	*j = SpecialCharactersPlusLicenses(v)
	return nil
}

const SpecialCharactersPlusLicensesGPL30 SpecialCharactersPlusLicenses = "GPL-3.0+"
const SpecialCharactersPlusLicensesMIT SpecialCharactersPlusLicenses = "MIT+"

var enumValues_License = []interface{}{
	"GPL-3.0",
	"MIT",
}
