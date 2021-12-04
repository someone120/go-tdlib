// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// ValidateOrderInfo Validates the order information provided by a user and returns the available shipping options for a flexible invoice
// @param chatID Chat identifier of the Invoice message
// @param messageID Message identifier
// @param orderInfo The order information, provided by the user
// @param allowSave True, if the order information can be saved
func (client *Client) ValidateOrderInfo(chatID int64, messageID int64, orderInfo *tdlib.OrderInfo, allowSave bool) (*tdlib.ValidatedOrderInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "validateOrderInfo",
		"chat_id":    chatID,
		"message_id": messageID,
		"order_info": orderInfo,
		"allow_save": allowSave,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var validatedOrderInfo tdlib.ValidatedOrderInfo
	err = json.Unmarshal(result.Raw, &validatedOrderInfo)
	return &validatedOrderInfo, err

}
