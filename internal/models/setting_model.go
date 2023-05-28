package models

type SettingModel struct {
	Notification      bool `json:"notification"`
	EmailNotification bool `json:"email_notification" bson:"email_notification"`
}
