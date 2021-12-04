// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// GetSavedAnimations Returns saved animations
func (client *Client) GetSavedAnimations() (*tdlib.Animations, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getSavedAnimations",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var animations tdlib.Animations
	err = json.Unmarshal(result.Raw, &animations)
	return &animations, err

}
