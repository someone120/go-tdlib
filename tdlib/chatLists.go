// AUTOGENERATED - DO NOT EDIT

package tdlib

// ChatLists Contains a list of chat lists
type ChatLists struct {
	tdCommon
	ChatLists []ChatList `json:"chat_lists"` // List of chat lists
}

// MessageType return the string telegram-type of ChatLists
func (chatLists *ChatLists) MessageType() string {
	return "chatLists"
}

// NewChatLists creates a new ChatLists
//
// @param chatLists List of chat lists
func NewChatLists(chatLists []ChatList) *ChatLists {
	chatListsTemp := ChatLists{
		tdCommon:  tdCommon{Type: "chatLists"},
		ChatLists: chatLists,
	}

	return &chatListsTemp
}