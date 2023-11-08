// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"

type Array struct {
	// MyArray corresponds to the JSON schema field "myArray".
	MyArray []interface{} `json:"myArray,omitempty" yaml:"myArray,omitempty" mapstructure:"myArray,omitempty"`

	// MyBooleanArray corresponds to the JSON schema field "myBooleanArray".
	MyBooleanArray []bool `json:"myBooleanArray,omitempty" yaml:"myBooleanArray,omitempty" mapstructure:"myBooleanArray,omitempty"`

	// MyIntegerArray corresponds to the JSON schema field "myIntegerArray".
	MyIntegerArray []int `json:"myIntegerArray,omitempty" yaml:"myIntegerArray,omitempty" mapstructure:"myIntegerArray,omitempty"`

	// MyNestedNullArray corresponds to the JSON schema field "myNestedNullArray".
	MyNestedNullArray [][]interface{} `json:"myNestedNullArray,omitempty" yaml:"myNestedNullArray,omitempty" mapstructure:"myNestedNullArray,omitempty"`

	// MyNullArray corresponds to the JSON schema field "myNullArray".
	MyNullArray []interface{} `json:"myNullArray,omitempty" yaml:"myNullArray,omitempty" mapstructure:"myNullArray,omitempty"`

	// MyNumberArray corresponds to the JSON schema field "myNumberArray".
	MyNumberArray []float64 `json:"myNumberArray,omitempty" yaml:"myNumberArray,omitempty" mapstructure:"myNumberArray,omitempty"`

	// MyObjectArray corresponds to the JSON schema field "myObjectArray".
	MyObjectArray []ArrayMyObjectArrayElem `json:"myObjectArray,omitempty" yaml:"myObjectArray,omitempty" mapstructure:"myObjectArray,omitempty"`

	// MyStringArray corresponds to the JSON schema field "myStringArray".
	MyStringArray []string `json:"myStringArray,omitempty" yaml:"myStringArray,omitempty" mapstructure:"myStringArray,omitempty"`
}

type ArrayMyObjectArrayElem map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Array) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	type Plain Array
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	for i0 := range plain.MyNestedNullArray {
		for i1 := range plain.MyNestedNullArray[i0] {
			if plain.MyNestedNullArray[i0][i1] != nil {
				return fmt.Errorf("field %s: must be null", fmt.Sprintf("myNestedNullArray[%d][%d]", i0, i1))
			}
		}
	}
	for i0 := range plain.MyNullArray {
		if plain.MyNullArray[i0] != nil {
			return fmt.Errorf("field %s: must be null", fmt.Sprintf("myNullArray[%d]", i0))
		}
	}
	*j = Array(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *Array) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	type Plain Array
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	for i0 := range plain.MyNestedNullArray {
		for i1 := range plain.MyNestedNullArray[i0] {
			if plain.MyNestedNullArray[i0][i1] != nil {
				return fmt.Errorf("field %s: must be null", fmt.Sprintf("myNestedNullArray[%d][%d]", i0, i1))
			}
		}
	}
	for i0 := range plain.MyNullArray {
		if plain.MyNullArray[i0] != nil {
			return fmt.Errorf("field %s: must be null", fmt.Sprintf("myNullArray[%d]", i0))
		}
	}
	*j = Array(plain)
	return nil
}
