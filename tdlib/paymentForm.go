// AUTOGENERATED - DO NOT EDIT

package tdlib

// PaymentForm Contains information about an invoice payment form
type PaymentForm struct {
	tdCommon
	ID                     JSONInt64               `json:"id"`                        // The payment form identifier
	Invoice                *Invoice                `json:"invoice"`                   // Full information of the invoice
	URL                    string                  `json:"url"`                       // Payment form URL
	SellerBotUserID        int64                   `json:"seller_bot_user_id"`        // User identifier of the seller bot
	PaymentsProviderUserID int64                   `json:"payments_provider_user_id"` // User identifier of the payment provider bot
	PaymentsProvider       *PaymentsProviderStripe `json:"payments_provider"`         // Contains information about the payment provider, if available, to support it natively without the need for opening the URL; may be null
	SavedOrderInfo         *OrderInfo              `json:"saved_order_info"`          // Saved server-side order information; may be null
	SavedCredentials       *SavedCredentials       `json:"saved_credentials"`         // Contains information about saved card credentials; may be null
	CanSaveCredentials     bool                    `json:"can_save_credentials"`      // True, if the user can choose to save credentials
	NeedPassword           bool                    `json:"need_password"`             // True, if the user will be able to save credentials protected by a password they set up
}

// MessageType return the string telegram-type of PaymentForm
func (paymentForm *PaymentForm) MessageType() string {
	return "paymentForm"
}

// NewPaymentForm creates a new PaymentForm
//
// @param iD The payment form identifier
// @param invoice Full information of the invoice
// @param uRL Payment form URL
// @param sellerBotUserID User identifier of the seller bot
// @param paymentsProviderUserID User identifier of the payment provider bot
// @param paymentsProvider Contains information about the payment provider, if available, to support it natively without the need for opening the URL; may be null
// @param savedOrderInfo Saved server-side order information; may be null
// @param savedCredentials Contains information about saved card credentials; may be null
// @param canSaveCredentials True, if the user can choose to save credentials
// @param needPassword True, if the user will be able to save credentials protected by a password they set up
func NewPaymentForm(iD JSONInt64, invoice *Invoice, uRL string, sellerBotUserID int64, paymentsProviderUserID int64, paymentsProvider *PaymentsProviderStripe, savedOrderInfo *OrderInfo, savedCredentials *SavedCredentials, canSaveCredentials bool, needPassword bool) *PaymentForm {
	paymentFormTemp := PaymentForm{
		tdCommon:               tdCommon{Type: "paymentForm"},
		ID:                     iD,
		Invoice:                invoice,
		URL:                    uRL,
		SellerBotUserID:        sellerBotUserID,
		PaymentsProviderUserID: paymentsProviderUserID,
		PaymentsProvider:       paymentsProvider,
		SavedOrderInfo:         savedOrderInfo,
		SavedCredentials:       savedCredentials,
		CanSaveCredentials:     canSaveCredentials,
		NeedPassword:           needPassword,
	}

	return &paymentFormTemp
}
