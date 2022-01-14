// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
	"fmt"
)

// LanguagePackStringValue Represents the value of a string in a language pack
type LanguagePackStringValue interface {
	GetLanguagePackStringValueEnum() LanguagePackStringValueEnum
}

// LanguagePackStringValueEnum Alias for abstract LanguagePackStringValue 'Sub-Classes', used as constant-enum here
type LanguagePackStringValueEnum string

// LanguagePackStringValue enums
const (
	LanguagePackStringValueOrdinaryType   LanguagePackStringValueEnum = "languagePackStringValueOrdinary"
	LanguagePackStringValuePluralizedType LanguagePackStringValueEnum = "languagePackStringValuePluralized"
	LanguagePackStringValueDeletedType    LanguagePackStringValueEnum = "languagePackStringValueDeleted"
)

func unmarshalLanguagePackStringValue(rawMsg *json.RawMessage) (LanguagePackStringValue, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch LanguagePackStringValueEnum(objMap["@type"].(string)) {
	case LanguagePackStringValueOrdinaryType:
		var languagePackStringValueOrdinary LanguagePackStringValueOrdinary
		err := json.Unmarshal(*rawMsg, &languagePackStringValueOrdinary)
		return &languagePackStringValueOrdinary, err

	case LanguagePackStringValuePluralizedType:
		var languagePackStringValuePluralized LanguagePackStringValuePluralized
		err := json.Unmarshal(*rawMsg, &languagePackStringValuePluralized)
		return &languagePackStringValuePluralized, err

	case LanguagePackStringValueDeletedType:
		var languagePackStringValueDeleted LanguagePackStringValueDeleted
		err := json.Unmarshal(*rawMsg, &languagePackStringValueDeleted)
		return &languagePackStringValueDeleted, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// LanguagePackStringValueOrdinary An ordinary language pack string
type LanguagePackStringValueOrdinary struct {
	tdCommon
	Value string `json:"value"` // String value
}

// MessageType return the string telegram-type of LanguagePackStringValueOrdinary
func (languagePackStringValueOrdinary *LanguagePackStringValueOrdinary) MessageType() string {
	return "languagePackStringValueOrdinary"
}

// NewLanguagePackStringValueOrdinary creates a new LanguagePackStringValueOrdinary
//
// @param value String value
func NewLanguagePackStringValueOrdinary(value string) *LanguagePackStringValueOrdinary {
	languagePackStringValueOrdinaryTemp := LanguagePackStringValueOrdinary{
		tdCommon: tdCommon{Type: "languagePackStringValueOrdinary"},
		Value:    value,
	}

	return &languagePackStringValueOrdinaryTemp
}

// GetLanguagePackStringValueEnum return the enum type of this object
func (languagePackStringValueOrdinary *LanguagePackStringValueOrdinary) GetLanguagePackStringValueEnum() LanguagePackStringValueEnum {
	return LanguagePackStringValueOrdinaryType
}

// LanguagePackStringValuePluralized A language pack string which has different forms based on the number of some object it mentions. See https://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html for more info
type LanguagePackStringValuePluralized struct {
	tdCommon
	ZeroValue  string `json:"zero_value"`  // Value for zero objects
	OneValue   string `json:"one_value"`   // Value for one object
	TwoValue   string `json:"two_value"`   // Value for two objects
	FewValue   string `json:"few_value"`   // Value for few objects
	ManyValue  string `json:"many_value"`  // Value for many objects
	OtherValue string `json:"other_value"` // Default value
}

// MessageType return the string telegram-type of LanguagePackStringValuePluralized
func (languagePackStringValuePluralized *LanguagePackStringValuePluralized) MessageType() string {
	return "languagePackStringValuePluralized"
}

// NewLanguagePackStringValuePluralized creates a new LanguagePackStringValuePluralized
//
// @param zeroValue Value for zero objects
// @param oneValue Value for one object
// @param twoValue Value for two objects
// @param fewValue Value for few objects
// @param manyValue Value for many objects
// @param otherValue Default value
func NewLanguagePackStringValuePluralized(zeroValue string, oneValue string, twoValue string, fewValue string, manyValue string, otherValue string) *LanguagePackStringValuePluralized {
	languagePackStringValuePluralizedTemp := LanguagePackStringValuePluralized{
		tdCommon:   tdCommon{Type: "languagePackStringValuePluralized"},
		ZeroValue:  zeroValue,
		OneValue:   oneValue,
		TwoValue:   twoValue,
		FewValue:   fewValue,
		ManyValue:  manyValue,
		OtherValue: otherValue,
	}

	return &languagePackStringValuePluralizedTemp
}

// GetLanguagePackStringValueEnum return the enum type of this object
func (languagePackStringValuePluralized *LanguagePackStringValuePluralized) GetLanguagePackStringValueEnum() LanguagePackStringValueEnum {
	return LanguagePackStringValuePluralizedType
}

// LanguagePackStringValueDeleted A deleted language pack string, the value should be taken from the built-in english language pack
type LanguagePackStringValueDeleted struct {
	tdCommon
}

// MessageType return the string telegram-type of LanguagePackStringValueDeleted
func (languagePackStringValueDeleted *LanguagePackStringValueDeleted) MessageType() string {
	return "languagePackStringValueDeleted"
}

// NewLanguagePackStringValueDeleted creates a new LanguagePackStringValueDeleted
//
func NewLanguagePackStringValueDeleted() *LanguagePackStringValueDeleted {
	languagePackStringValueDeletedTemp := LanguagePackStringValueDeleted{
		tdCommon: tdCommon{Type: "languagePackStringValueDeleted"},
	}

	return &languagePackStringValueDeletedTemp
}

// GetLanguagePackStringValueEnum return the enum type of this object
func (languagePackStringValueDeleted *LanguagePackStringValueDeleted) GetLanguagePackStringValueEnum() LanguagePackStringValueEnum {
	return LanguagePackStringValueDeletedType
}