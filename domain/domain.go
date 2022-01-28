package domain

type ItemsNWPayload struct {
	Status string `json:"status"`
}

type DenounceItem struct {
	ReportReasonID string `json:"report_reason_id"`
	Comment        string `json:"comment"`
	CallerID       string `json:"caller_id"`
	ItemID         string `json:"item_id"`
	ElementID      string `json:"element_id"`
	Type           string `json:"type"`
	Origin         string `json:"origin"`
}

type ItemsOWPayload struct {
	ItemID      string `json:"item_id"`
	Observation string `json:"observation"`
	ActionType  string `json:"action_type"`
}

type Email struct {
	UserID    int     `json:"user_id"`
	Template  string  `json:"template"`
	Recipient string  `json:"recipient"`
	Context   Context `json:"context"`
}

type Context struct {
	EmailSubject string `json:"email_subject"`
	UsersData    string `json:"users_data"`
}
