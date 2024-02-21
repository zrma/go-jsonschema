// Code generated by github.com/zrma/go-jsonschema, DO NOT EDIT.

package test

type Ref struct {
	// MyThing corresponds to the JSON schema field "myThing".
	MyThing *Thing_1 `json:"myThing,omitempty" yaml:"myThing,omitempty" mapstructure:"myThing,omitempty"`

	// MyThing2 corresponds to the JSON schema field "myThing2".
	MyThing2 *Thing_1 `json:"myThing2,omitempty" yaml:"myThing2,omitempty" mapstructure:"myThing2,omitempty"`
}

type RefExternalFileWithDupe struct {
	// MyExternalThing corresponds to the JSON schema field "myExternalThing".
	MyExternalThing *Thing_1 `json:"myExternalThing,omitempty" yaml:"myExternalThing,omitempty" mapstructure:"myExternalThing,omitempty"`

	// MyThing corresponds to the JSON schema field "myThing".
	MyThing *Thing `json:"myThing,omitempty" yaml:"myThing,omitempty" mapstructure:"myThing,omitempty"`
}

type Thing struct {
	// Something corresponds to the JSON schema field "something".
	Something *string `json:"something,omitempty" yaml:"something,omitempty" mapstructure:"something,omitempty"`
}

type Thing_1 struct {
	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
}
