// Code generated by github.com/zrma/go-jsonschema, DO NOT EDIT.

package test

type TypeMultiple struct {
	// All corresponds to the JSON schema field "all".
	All interface{} `json:"all,omitempty" yaml:"all,omitempty" mapstructure:"all,omitempty"`

	// AllPrimitives corresponds to the JSON schema field "allPrimitives".
	AllPrimitives interface{} `json:"allPrimitives,omitempty" yaml:"allPrimitives,omitempty" mapstructure:"allPrimitives,omitempty"`

	// ArrayOfAll corresponds to the JSON schema field "arrayOfAll".
	ArrayOfAll []interface{} `json:"arrayOfAll,omitempty" yaml:"arrayOfAll,omitempty" mapstructure:"arrayOfAll,omitempty"`

	// ArrayOfAllPrimitives corresponds to the JSON schema field
	// "arrayOfAllPrimitives".
	ArrayOfAllPrimitives []interface{} `json:"arrayOfAllPrimitives,omitempty" yaml:"arrayOfAllPrimitives,omitempty" mapstructure:"arrayOfAllPrimitives,omitempty"`
}
