package domain

type ItemsNWPayload struct {
	Status string `json:"status"`
}

type ItemsOWPayload struct {
	ItemID string `json:"item_id"`
	Observation string `json:"observation"`
	ActionType string `json:"action_type"`
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
