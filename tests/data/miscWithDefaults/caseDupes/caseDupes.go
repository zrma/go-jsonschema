// Code generated by github.com/zrma/go-jsonschema, DO NOT EDIT.

package test

type CaseDupes struct {
	// SomeField corresponds to the JSON schema field "SomeField".
	SomeField *string `json:"SomeField,omitempty" yaml:"SomeField,omitempty" mapstructure:"SomeField,omitempty"`

	// SomeField_2 corresponds to the JSON schema field "someField".
	SomeField_2 *string `json:"someField,omitempty" yaml:"someField,omitempty" mapstructure:"someField,omitempty"`

	// SomeField_3 corresponds to the JSON schema field "some_Field".
	SomeField_3 *string `json:"some_Field,omitempty" yaml:"some_Field,omitempty" mapstructure:"some_Field,omitempty"`

	// SomeField_4 corresponds to the JSON schema field "some_field".
	SomeField_4 *string `json:"some_field,omitempty" yaml:"some_field,omitempty" mapstructure:"some_field,omitempty"`

	// Somefield corresponds to the JSON schema field "somefield".
	Somefield *string `json:"somefield,omitempty" yaml:"somefield,omitempty" mapstructure:"somefield,omitempty"`
}
