// Code generated by github.com/zrma/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"
import "reflect"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GopkgYAMLv3) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	type Plain GopkgYAMLv3
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.MyNull != nil {
		return fmt.Errorf("field %s: must be null", "myNull")
	}
	*j = GopkgYAMLv3(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *GopkgYAMLv3) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	type Plain GopkgYAMLv3
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	if plain.MyNull != nil {
		return fmt.Errorf("field %s: must be null", "myNull")
	}
	*j = GopkgYAMLv3(plain)
	return nil
}

type GopkgYAMLv3 struct {
	// MyBoolean corresponds to the JSON schema field "myBoolean".
	MyBoolean *bool `json:"myBoolean,omitempty" yaml:"myBoolean,omitempty" mapstructure:"myBoolean,omitempty"`

	// MyEnum corresponds to the JSON schema field "myEnum".
	MyEnum *GopkgYAMLv3MyEnum `json:"myEnum,omitempty" yaml:"myEnum,omitempty" mapstructure:"myEnum,omitempty"`

	// MyInteger corresponds to the JSON schema field "myInteger".
	MyInteger *int `json:"myInteger,omitempty" yaml:"myInteger,omitempty" mapstructure:"myInteger,omitempty"`

	// MyNull corresponds to the JSON schema field "myNull".
	MyNull interface{} `json:"myNull,omitempty" yaml:"myNull,omitempty" mapstructure:"myNull,omitempty"`

	// MyNumber corresponds to the JSON schema field "myNumber".
	MyNumber *float64 `json:"myNumber,omitempty" yaml:"myNumber,omitempty" mapstructure:"myNumber,omitempty"`

	// MyString corresponds to the JSON schema field "myString".
	MyString *string `json:"myString,omitempty" yaml:"myString,omitempty" mapstructure:"myString,omitempty"`
}

type GopkgYAMLv3MyEnum string

const GopkgYAMLv3MyEnumX GopkgYAMLv3MyEnum = "x"
const GopkgYAMLv3MyEnumY GopkgYAMLv3MyEnum = "y"

var enumValues_GopkgYAMLv3MyEnum = []interface{}{
	"x",
	"y",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GopkgYAMLv3MyEnum) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GopkgYAMLv3MyEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GopkgYAMLv3MyEnum, v)
	}
	*j = GopkgYAMLv3MyEnum(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *GopkgYAMLv3MyEnum) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GopkgYAMLv3MyEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GopkgYAMLv3MyEnum, v)
	}
	*j = GopkgYAMLv3MyEnum(v)
	return nil
}
