// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// TestCallBytes Returns the received bytes; for testing only. This is an offline method. Can be called before authorization
// @param x Bytes to return
func (client *Client) TestCallBytes(x []byte) (*tdlib.TestBytes, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "testCallBytes",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var testBytes tdlib.TestBytes
	err = json.Unmarshal(result.Raw, &testBytes)
	return &testBytes, err

}
