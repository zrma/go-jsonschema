// Code generated by github.com/zrma/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"

type RequiredFieldsMyObject struct {
	// MyNestedObjectString corresponds to the JSON schema field
	// "myNestedObjectString".
	MyNestedObjectString string `json:"myNestedObjectString" yaml:"myNestedObjectString" mapstructure:"myNestedObjectString"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequiredFieldsMyObject) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["myNestedObjectString"]; !ok || v == nil {
		return fmt.Errorf("field myNestedObjectString in RequiredFieldsMyObject: required")
	}
	type Plain RequiredFieldsMyObject
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RequiredFieldsMyObject(plain)
	return nil
}

type RequiredFieldsMyObjectArrayElem struct {
	// MyNestedObjectString corresponds to the JSON schema field
	// "myNestedObjectString".
	MyNestedObjectString string `json:"myNestedObjectString" yaml:"myNestedObjectString" mapstructure:"myNestedObjectString"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequiredFieldsMyObjectArrayElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["myNestedObjectString"]; !ok || v == nil {
		return fmt.Errorf("field myNestedObjectString in RequiredFieldsMyObjectArrayElem: required")
	}
	type Plain RequiredFieldsMyObjectArrayElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RequiredFieldsMyObjectArrayElem(plain)
	return nil
}

type RequiredFields struct {
	// MyBoolean corresponds to the JSON schema field "myBoolean".
	MyBoolean bool `json:"myBoolean" yaml:"myBoolean" mapstructure:"myBoolean"`

	// MyBooleanArray corresponds to the JSON schema field "myBooleanArray".
	MyBooleanArray []bool `json:"myBooleanArray" yaml:"myBooleanArray" mapstructure:"myBooleanArray"`

	// MyInteger corresponds to the JSON schema field "myInteger".
	MyInteger *int `json:"myInteger,omitempty" yaml:"myInteger,omitempty" mapstructure:"myInteger,omitempty"`

	// MyIntegerArray corresponds to the JSON schema field "myIntegerArray".
	MyIntegerArray []int `json:"myIntegerArray,omitempty" yaml:"myIntegerArray,omitempty" mapstructure:"myIntegerArray,omitempty"`

	// MyNull corresponds to the JSON schema field "myNull".
	MyNull interface{} `json:"myNull" yaml:"myNull" mapstructure:"myNull"`

	// MyNullArray corresponds to the JSON schema field "myNullArray".
	MyNullArray []interface{} `json:"myNullArray" yaml:"myNullArray" mapstructure:"myNullArray"`

	// MyNumber corresponds to the JSON schema field "myNumber".
	MyNumber float64 `json:"myNumber" yaml:"myNumber" mapstructure:"myNumber"`

	// MyNumberArray corresponds to the JSON schema field "myNumberArray".
	MyNumberArray []float64 `json:"myNumberArray" yaml:"myNumberArray" mapstructure:"myNumberArray"`

	// MyObject corresponds to the JSON schema field "myObject".
	MyObject RequiredFieldsMyObject `json:"myObject" yaml:"myObject" mapstructure:"myObject"`

	// MyObjectArray corresponds to the JSON schema field "myObjectArray".
	MyObjectArray []RequiredFieldsMyObjectArrayElem `json:"myObjectArray" yaml:"myObjectArray" mapstructure:"myObjectArray"`

	// MyString corresponds to the JSON schema field "myString".
	MyString string `json:"myString" yaml:"myString" mapstructure:"myString"`

	// MyStringArray corresponds to the JSON schema field "myStringArray".
	MyStringArray []string `json:"myStringArray" yaml:"myStringArray" mapstructure:"myStringArray"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequiredFields) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["myBoolean"]; !ok || v == nil {
		return fmt.Errorf("field myBoolean in RequiredFields: required")
	}
	if v, ok := raw["myBooleanArray"]; !ok || v == nil {
		return fmt.Errorf("field myBooleanArray in RequiredFields: required")
	}
	if v, ok := raw["myNull"]; !ok || v == nil {
		return fmt.Errorf("field myNull in RequiredFields: required")
	}
	if v, ok := raw["myNullArray"]; !ok || v == nil {
		return fmt.Errorf("field myNullArray in RequiredFields: required")
	}
	if v, ok := raw["myNumber"]; !ok || v == nil {
		return fmt.Errorf("field myNumber in RequiredFields: required")
	}
	if v, ok := raw["myNumberArray"]; !ok || v == nil {
		return fmt.Errorf("field myNumberArray in RequiredFields: required")
	}
	if v, ok := raw["myObject"]; !ok || v == nil {
		return fmt.Errorf("field myObject in RequiredFields: required")
	}
	if v, ok := raw["myObjectArray"]; !ok || v == nil {
		return fmt.Errorf("field myObjectArray in RequiredFields: required")
	}
	if v, ok := raw["myString"]; !ok || v == nil {
		return fmt.Errorf("field myString in RequiredFields: required")
	}
	if v, ok := raw["myStringArray"]; !ok || v == nil {
		return fmt.Errorf("field myStringArray in RequiredFields: required")
	}
	type Plain RequiredFields
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if plain.MyNull != nil {
		return fmt.Errorf("field %s: must be null", "myNull")
	}
	for i0 := range plain.MyNullArray {
		if plain.MyNullArray[i0] != nil {
			return fmt.Errorf("field %s: must be null", fmt.Sprintf("myNullArray[%d]", i0))
		}
	}
	*j = RequiredFields(plain)
	return nil
}
