package entity

type GetMessages struct {
	Code       string   `json:"code"`
	Data       Messages `json:"data"`
	Success    string   `json:"success"`
	ErrCode    string   `json:"err_code"`
	RequestID  string   `json:"request_id"`
	ErrMessage string   `json:"err_message"`
}

type Message struct {
	FromAccountType string `json:"from_account_type"`
	ProcessMsg      string `json:"process_msg"`
	SessionID       string `json:"session_id"`
	MessageID       string `json:"message_id"`
	Type            string `json:"type"`
	Content         string `json:"content"`
	ToAccountID     string `json:"to_account_id"`
	SendTime        string `json:"send_time"`
	AutoReply       string `json:"auto_reply"`
	ToAccountType   string `json:"to_account_type"`
	SiteID          string `json:"site_id"`
	TemplateID      string `json:"template_id"`
	FromAccountID   string `json:"from_account_id"`
	Status          string `json:"status"`
}

type Messages struct {
	LastMessageID string    `json:"last_message_id"`
	MessageList   []Message `json:"message_list"`
	NextStartTime string    `json:"next_start_time"`
	HasMore       string    `json:"has_more"`
}
