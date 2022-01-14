// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// TestCallString Returns the received string; for testing only. This is an offline method. Can be called before authorization
// @param x String to return
func (client *Client) TestCallString(x string) (*tdlib.TestString, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "testCallString",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var testString tdlib.TestString
	err = json.Unmarshal(result.Raw, &testString)
	return &testString, err

}