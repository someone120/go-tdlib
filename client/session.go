// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// ConfirmQrCodeAuthentication Confirms QR code authentication on another device. Returns created session on success
// @param link A link from a QR code. The link must be scanned by the in-app camera
func (client *Client) ConfirmQrCodeAuthentication(link string) (*tdlib.Session, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "confirmQrCodeAuthentication",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var session tdlib.Session
	err = json.Unmarshal(result.Raw, &session)
	return &session, err

}