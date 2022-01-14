// AUTOGENERATED - DO NOT EDIT

package tdlib

// TestVectorString A simple object containing a vector of strings; for testing only
type TestVectorString struct {
	tdCommon
	Value []string `json:"value"` // Vector of strings
}

// MessageType return the string telegram-type of TestVectorString
func (testVectorString *TestVectorString) MessageType() string {
	return "testVectorString"
}

// NewTestVectorString creates a new TestVectorString
//
// @param value Vector of strings
func NewTestVectorString(value []string) *TestVectorString {
	testVectorStringTemp := TestVectorString{
		tdCommon: tdCommon{Type: "testVectorString"},
		Value:    value,
	}

	return &testVectorStringTemp
}