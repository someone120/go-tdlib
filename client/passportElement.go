// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// GetPassportElement Returns one of the available Telegram Passport elements
// @param typeParam Telegram Passport element type
// @param password Password of the current user
func (client *Client) GetPassportElement(typeParam tdlib.PassportElementType, password string) (tdlib.PassportElement, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "getPassportElement",
		"type":     typeParam,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.PassportElementEnum(result.Data["@type"].(string)) {

	case tdlib.PassportElementPersonalDetailsType:
		var passportElement tdlib.PassportElementPersonalDetails
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementPassportType:
		var passportElement tdlib.PassportElementPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementDriverLicenseType:
		var passportElement tdlib.PassportElementDriverLicense
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementIDentityCardType:
		var passportElement tdlib.PassportElementIDentityCard
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementInternalPassportType:
		var passportElement tdlib.PassportElementInternalPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementAddressType:
		var passportElement tdlib.PassportElementAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementUtilityBillType:
		var passportElement tdlib.PassportElementUtilityBill
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementBankStatementType:
		var passportElement tdlib.PassportElementBankStatement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementRentalAgreementType:
		var passportElement tdlib.PassportElementRentalAgreement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementPassportRegistrationType:
		var passportElement tdlib.PassportElementPassportRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementTemporaryRegistrationType:
		var passportElement tdlib.PassportElementTemporaryRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementPhoneNumberType:
		var passportElement tdlib.PassportElementPhoneNumber
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementEmailAddressType:
		var passportElement tdlib.PassportElementEmailAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SetPassportElement Adds an element to the user's Telegram Passport. May return an error with a message "PHONE_VERIFICATION_NEEDED" or "EMAIL_VERIFICATION_NEEDED" if the chosen phone number or the chosen email address must be verified first
// @param element Input Telegram Passport element
// @param password Password of the current user
func (client *Client) SetPassportElement(element tdlib.InputPassportElement, password string) (tdlib.PassportElement, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "setPassportElement",
		"element":  element,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.PassportElementEnum(result.Data["@type"].(string)) {

	case tdlib.PassportElementPersonalDetailsType:
		var passportElement tdlib.PassportElementPersonalDetails
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementPassportType:
		var passportElement tdlib.PassportElementPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementDriverLicenseType:
		var passportElement tdlib.PassportElementDriverLicense
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementIDentityCardType:
		var passportElement tdlib.PassportElementIDentityCard
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementInternalPassportType:
		var passportElement tdlib.PassportElementInternalPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementAddressType:
		var passportElement tdlib.PassportElementAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementUtilityBillType:
		var passportElement tdlib.PassportElementUtilityBill
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementBankStatementType:
		var passportElement tdlib.PassportElementBankStatement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementRentalAgreementType:
		var passportElement tdlib.PassportElementRentalAgreement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementPassportRegistrationType:
		var passportElement tdlib.PassportElementPassportRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementTemporaryRegistrationType:
		var passportElement tdlib.PassportElementTemporaryRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementPhoneNumberType:
		var passportElement tdlib.PassportElementPhoneNumber
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case tdlib.PassportElementEmailAddressType:
		var passportElement tdlib.PassportElementEmailAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
