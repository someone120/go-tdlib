// AUTOGENERATED - DO NOT EDIT

package tdlib

// TestInt A simple object containing a number; for testing only
type TestInt struct {
	tdCommon
	Value int32 `json:"value"` // Number
}

// MessageType return the string telegram-type of TestInt
func (testInt *TestInt) MessageType() string {
	return "testInt"
}

// NewTestInt creates a new TestInt
//
// @param value Number
func NewTestInt(value int32) *TestInt {
	testIntTemp := TestInt{
		tdCommon: tdCommon{Type: "testInt"},
		Value:    value,
	}

	return &testIntTemp
}