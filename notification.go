package gonextcloud

import "time"

// Notification is a nextcloud notification (from notification app)
type Notification struct {
	NotificationID        int           `json:"notification_id"`
	App                   string        `json:"app"`
	User                  string        `json:"user"`
	Datetime              time.Time     `json:"datetime"`
	ObjectType            string        `json:"object_type"`
	ObjectID              string        `json:"object_id"`
	Subject               string        `json:"subject"`
	Message               string        `json:"message"`
	Link                  string        `json:"link"`
	SubjectRich           string        `json:"subjectRich"`
	SubjectRichParameters []interface{} `json:"subjectRichParameters"`
	MessageRich           string        `json:"messageRich"`
	MessageRichParameters []interface{} `json:"messageRichParameters"`
	Icon                  string        `json:"icon"`
	Actions               []interface{} `json:"actions"`
}
