// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// GetActiveSessions Returns all active sessions of the current user
func (client *Client) GetActiveSessions() (*tdlib.Sessions, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getActiveSessions",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var sessions tdlib.Sessions
	err = json.Unmarshal(result.Raw, &sessions)
	return &sessions, err

}
