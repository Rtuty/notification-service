package entities

type (
	Client struct {
		ID          string `json:"client_id"`
		PhoneNumber string `json:"phone_number"`
		PhoneCode   string `json:"phone_code"`
		Tag         string `json:"tag"`
		TimeZone    string `json:"time_zone"`
	}

	Mailing struct {
		ID         string  `json:"mailing_id"`
		StartTime  string  `json:"start_time"`
		FinishTime string  `json:"finish_time"`
		Message    Message `json:"message"`
		Filter     string  `json:"filter"`
	}

	Message struct {
		ID        string `json:"message_id"`
		StartTime string `json:"start_time"`
		Status    string `json:"status"`
		MailingID string `json:"mailing_id"`
		ClientID  string `json:"client_id"`
	}
)
