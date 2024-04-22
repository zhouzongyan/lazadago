package entity

type GetSeesions struct {
	Code       string   `json:"code"`
	Data       Sessions `json:"data"`
	Success    string   `json:"success"`
	ErrCode    string   `json:"err_code"`
	RequestID  string   `json:"request_id"`
	ErrMessage string   `json:"err_message"`
}
type Session struct {
	Summary         string   `json:"summary"`
	UnreadCount     string   `json:"unread_count"`
	LastMessageID   string   `json:"last_message_id"`
	HeadURL         string   `json:"head_url"`
	SelfPosition    string   `json:"self_position"`
	SiteID          string   `json:"site_id"`
	LastMessageTime string   `json:"last_message_time"`
	SessionID       string   `json:"session_id"`
	BuyerID         string   `json:"buyer_id"`
	Title           string   `json:"title"`
	ToPosition      string   `json:"to_position"`
	Tags            []string `json:"tags"`
}
type Sessions struct {
	SessionList   []Session `json:"session_list"`
	NextStartTime string    `json:"next_start_time"`
	HasMore       string    `json:"has_more"`
	LastSessionID string    `json:"last_session_id"`
}
