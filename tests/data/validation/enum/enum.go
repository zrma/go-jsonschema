// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"
import "reflect"

type EnumMyBooleanTypedEnum bool

var enumValues_EnumMyBooleanTypedEnum = []interface{}{
	true,
	false,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyBooleanTypedEnum) UnmarshalJSON(b []byte) error {
	var v bool
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyBooleanTypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyBooleanTypedEnum, v)
	}
	*j = EnumMyBooleanTypedEnum(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyBooleanTypedEnum) UnmarshalYAML(value *yaml.Node) error {
	var v bool
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyBooleanTypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyBooleanTypedEnum, v)
	}
	*j = EnumMyBooleanTypedEnum(v)
	return nil
}

type EnumMyBooleanUntypedEnum bool

var enumValues_EnumMyBooleanUntypedEnum = []interface{}{
	true,
	false,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyBooleanUntypedEnum) UnmarshalJSON(b []byte) error {
	var v bool
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyBooleanUntypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyBooleanUntypedEnum, v)
	}
	*j = EnumMyBooleanUntypedEnum(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyBooleanUntypedEnum) UnmarshalYAML(value *yaml.Node) error {
	var v bool
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyBooleanUntypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyBooleanUntypedEnum, v)
	}
	*j = EnumMyBooleanUntypedEnum(v)
	return nil
}

type EnumMyIntegerTypedEnum int

var enumValues_EnumMyIntegerTypedEnum = []interface{}{
	1,
	2,
	3,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyIntegerTypedEnum) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyIntegerTypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyIntegerTypedEnum, v)
	}
	*j = EnumMyIntegerTypedEnum(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyIntegerTypedEnum) UnmarshalYAML(value *yaml.Node) error {
	var v int
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyIntegerTypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyIntegerTypedEnum, v)
	}
	*j = EnumMyIntegerTypedEnum(v)
	return nil
}

type EnumMyMixedTypeEnum struct {
	Value interface{}
}

var enumValues_EnumMyMixedTypeEnum = []interface{}{
	42.0,
	"smurf",
}

// MarshalJSON implements json.Marshaler.
func (j *EnumMyMixedTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyMixedTypeEnum) UnmarshalJSON(b []byte) error {
	var v struct {
		Value interface{}
	}
	if err := json.Unmarshal(b, &v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyMixedTypeEnum {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyMixedTypeEnum, v.Value)
	}
	*j = EnumMyMixedTypeEnum(v)
	return nil
}

// MarshalYAML implements yaml.Marshal.
func (j *EnumMyMixedTypeEnum) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(j.Value)
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyMixedTypeEnum) UnmarshalYAML(value *yaml.Node) error {
	var v struct {
		Value interface{}
	}
	if err := value.Decode(&v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyMixedTypeEnum {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyMixedTypeEnum, v.Value)
	}
	*j = EnumMyMixedTypeEnum(v)
	return nil
}

type EnumMyMixedUntypedEnum struct {
	Value interface{}
}

var enumValues_EnumMyMixedUntypedEnum = []interface{}{
	"red",
	1.0,
	true,
	nil,
}

// MarshalJSON implements json.Marshaler.
func (j *EnumMyMixedUntypedEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyMixedUntypedEnum) UnmarshalJSON(b []byte) error {
	var v struct {
		Value interface{}
	}
	if err := json.Unmarshal(b, &v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyMixedUntypedEnum {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyMixedUntypedEnum, v.Value)
	}
	*j = EnumMyMixedUntypedEnum(v)
	return nil
}

// MarshalYAML implements yaml.Marshal.
func (j *EnumMyMixedUntypedEnum) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(j.Value)
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyMixedUntypedEnum) UnmarshalYAML(value *yaml.Node) error {
	var v struct {
		Value interface{}
	}
	if err := value.Decode(&v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyMixedUntypedEnum {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyMixedUntypedEnum, v.Value)
	}
	*j = EnumMyMixedUntypedEnum(v)
	return nil
}

type EnumMyNullTypedEnum struct {
	Value interface{}
}

var enumValues_EnumMyNullTypedEnum = []interface{}{
	nil,
}

// MarshalJSON implements json.Marshaler.
func (j *EnumMyNullTypedEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyNullTypedEnum) UnmarshalJSON(b []byte) error {
	var v struct {
		Value interface{}
	}
	if err := json.Unmarshal(b, &v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyNullTypedEnum {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyNullTypedEnum, v.Value)
	}
	*j = EnumMyNullTypedEnum(v)
	return nil
}

// MarshalYAML implements yaml.Marshal.
func (j *EnumMyNullTypedEnum) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(j.Value)
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyNullTypedEnum) UnmarshalYAML(value *yaml.Node) error {
	var v struct {
		Value interface{}
	}
	if err := value.Decode(&v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyNullTypedEnum {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyNullTypedEnum, v.Value)
	}
	*j = EnumMyNullTypedEnum(v)
	return nil
}

type EnumMyNullUntypedEnum struct {
	Value interface{}
}

var enumValues_EnumMyNullUntypedEnum = []interface{}{
	nil,
}

// MarshalJSON implements json.Marshaler.
func (j *EnumMyNullUntypedEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyNullUntypedEnum) UnmarshalJSON(b []byte) error {
	var v struct {
		Value interface{}
	}
	if err := json.Unmarshal(b, &v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyNullUntypedEnum {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyNullUntypedEnum, v.Value)
	}
	*j = EnumMyNullUntypedEnum(v)
	return nil
}

// MarshalYAML implements yaml.Marshal.
func (j *EnumMyNullUntypedEnum) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(j.Value)
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyNullUntypedEnum) UnmarshalYAML(value *yaml.Node) error {
	var v struct {
		Value interface{}
	}
	if err := value.Decode(&v.Value); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyNullUntypedEnum {
		if reflect.DeepEqual(v.Value, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyNullUntypedEnum, v.Value)
	}
	*j = EnumMyNullUntypedEnum(v)
	return nil
}

type EnumMyNumberTypedEnum float64

var enumValues_EnumMyNumberTypedEnum = []interface{}{
	1.0,
	2.0,
	3.0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyNumberTypedEnum) UnmarshalJSON(b []byte) error {
	var v float64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyNumberTypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyNumberTypedEnum, v)
	}
	*j = EnumMyNumberTypedEnum(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyNumberTypedEnum) UnmarshalYAML(value *yaml.Node) error {
	var v float64
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyNumberTypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyNumberTypedEnum, v)
	}
	*j = EnumMyNumberTypedEnum(v)
	return nil
}

type EnumMyNumberUntypedEnum float64

var enumValues_EnumMyNumberUntypedEnum = []interface{}{
	1.0,
	2.0,
	3.0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyNumberUntypedEnum) UnmarshalJSON(b []byte) error {
	var v float64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyNumberUntypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyNumberUntypedEnum, v)
	}
	*j = EnumMyNumberUntypedEnum(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyNumberUntypedEnum) UnmarshalYAML(value *yaml.Node) error {
	var v float64
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyNumberUntypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyNumberUntypedEnum, v)
	}
	*j = EnumMyNumberUntypedEnum(v)
	return nil
}

type EnumMyStringTypedEnum string

var enumValues_EnumMyStringTypedEnum = []interface{}{
	"red",
	"blue",
	"green",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyStringTypedEnum) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyStringTypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyStringTypedEnum, v)
	}
	*j = EnumMyStringTypedEnum(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyStringTypedEnum) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyStringTypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyStringTypedEnum, v)
	}
	*j = EnumMyStringTypedEnum(v)
	return nil
}

const EnumMyStringTypedEnumBlue EnumMyStringTypedEnum = "blue"
const EnumMyStringTypedEnumGreen EnumMyStringTypedEnum = "green"
const EnumMyStringTypedEnumRed EnumMyStringTypedEnum = "red"

type EnumMyStringUntypedEnum string

var enumValues_EnumMyStringUntypedEnum = []interface{}{
	"red",
	"blue",
	"green",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnumMyStringUntypedEnum) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyStringUntypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyStringUntypedEnum, v)
	}
	*j = EnumMyStringUntypedEnum(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EnumMyStringUntypedEnum) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnumMyStringUntypedEnum {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnumMyStringUntypedEnum, v)
	}
	*j = EnumMyStringUntypedEnum(v)
	return nil
}

type Enum struct {
	// MyBooleanTypedEnum corresponds to the JSON schema field "myBooleanTypedEnum".
	MyBooleanTypedEnum *EnumMyBooleanTypedEnum `json:"myBooleanTypedEnum,omitempty" yaml:"myBooleanTypedEnum,omitempty" mapstructure:"myBooleanTypedEnum,omitempty"`

	// MyBooleanUntypedEnum corresponds to the JSON schema field
	// "myBooleanUntypedEnum".
	MyBooleanUntypedEnum *EnumMyBooleanUntypedEnum `json:"myBooleanUntypedEnum,omitempty" yaml:"myBooleanUntypedEnum,omitempty" mapstructure:"myBooleanUntypedEnum,omitempty"`

	// MyIntegerTypedEnum corresponds to the JSON schema field "myIntegerTypedEnum".
	MyIntegerTypedEnum *EnumMyIntegerTypedEnum `json:"myIntegerTypedEnum,omitempty" yaml:"myIntegerTypedEnum,omitempty" mapstructure:"myIntegerTypedEnum,omitempty"`

	// MyMixedTypeEnum corresponds to the JSON schema field "myMixedTypeEnum".
	MyMixedTypeEnum *EnumMyMixedTypeEnum `json:"myMixedTypeEnum,omitempty" yaml:"myMixedTypeEnum,omitempty" mapstructure:"myMixedTypeEnum,omitempty"`

	// MyMixedUntypedEnum corresponds to the JSON schema field "myMixedUntypedEnum".
	MyMixedUntypedEnum *EnumMyMixedUntypedEnum `json:"myMixedUntypedEnum,omitempty" yaml:"myMixedUntypedEnum,omitempty" mapstructure:"myMixedUntypedEnum,omitempty"`

	// MyNullTypedEnum corresponds to the JSON schema field "myNullTypedEnum".
	MyNullTypedEnum *EnumMyNullTypedEnum `json:"myNullTypedEnum,omitempty" yaml:"myNullTypedEnum,omitempty" mapstructure:"myNullTypedEnum,omitempty"`

	// MyNullUntypedEnum corresponds to the JSON schema field "myNullUntypedEnum".
	MyNullUntypedEnum *EnumMyNullUntypedEnum `json:"myNullUntypedEnum,omitempty" yaml:"myNullUntypedEnum,omitempty" mapstructure:"myNullUntypedEnum,omitempty"`

	// MyNumberTypedEnum corresponds to the JSON schema field "myNumberTypedEnum".
	MyNumberTypedEnum *EnumMyNumberTypedEnum `json:"myNumberTypedEnum,omitempty" yaml:"myNumberTypedEnum,omitempty" mapstructure:"myNumberTypedEnum,omitempty"`

	// MyNumberUntypedEnum corresponds to the JSON schema field "myNumberUntypedEnum".
	MyNumberUntypedEnum *EnumMyNumberUntypedEnum `json:"myNumberUntypedEnum,omitempty" yaml:"myNumberUntypedEnum,omitempty" mapstructure:"myNumberUntypedEnum,omitempty"`

	// MyStringTypedEnum corresponds to the JSON schema field "myStringTypedEnum".
	MyStringTypedEnum *EnumMyStringTypedEnum `json:"myStringTypedEnum,omitempty" yaml:"myStringTypedEnum,omitempty" mapstructure:"myStringTypedEnum,omitempty"`

	// MyStringUntypedEnum corresponds to the JSON schema field "myStringUntypedEnum".
	MyStringUntypedEnum *EnumMyStringUntypedEnum `json:"myStringUntypedEnum,omitempty" yaml:"myStringUntypedEnum,omitempty" mapstructure:"myStringUntypedEnum,omitempty"`
}

const EnumMyStringUntypedEnumBlue EnumMyStringUntypedEnum = "blue"
const EnumMyStringUntypedEnumGreen EnumMyStringUntypedEnum = "green"
const EnumMyStringUntypedEnumRed EnumMyStringUntypedEnum = "red"
