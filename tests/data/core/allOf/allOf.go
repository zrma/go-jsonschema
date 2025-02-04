// Code generated by github.com/zrma/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AllOfConfigurationsElem) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if v, ok := raw["bar"]; !ok || v == nil {
		return fmt.Errorf("field bar in AllOfConfigurationsElem: required")
	}
	if v, ok := raw["foo"]; !ok || v == nil {
		return fmt.Errorf("field foo in AllOfConfigurationsElem: required")
	}
	type Plain AllOfConfigurationsElem
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = AllOfConfigurationsElem(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *AllOfConfigurationsElem) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	if v, ok := raw["bar"]; !ok || v == nil {
		return fmt.Errorf("field bar in AllOfConfigurationsElem: required")
	}
	if v, ok := raw["foo"]; !ok || v == nil {
		return fmt.Errorf("field foo in AllOfConfigurationsElem: required")
	}
	type Plain AllOfConfigurationsElem
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = AllOfConfigurationsElem(plain)
	return nil
}

type AllOf struct {
	// Configurations corresponds to the JSON schema field "configurations".
	Configurations []AllOfConfigurationsElem `json:"configurations,omitempty" yaml:"configurations,omitempty" mapstructure:"configurations,omitempty"`
}

type AllOfConfigurationsElem struct {
	// Bar corresponds to the JSON schema field "bar".
	Bar float64 `json:"bar" yaml:"bar" mapstructure:"bar"`

	// Foo corresponds to the JSON schema field "foo".
	Foo string `json:"foo" yaml:"foo" mapstructure:"foo"`
}
