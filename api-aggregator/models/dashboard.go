package models

type DashboardResponse struct {
	User          User           `json:"user"`
	Orders        []Order        `json:"orders"`
	Notifications []Notification `json:"notifications"`
}
