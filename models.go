package yaygo

type User struct {
	Biography string `json:"biography"`
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
}

type SNSInfo struct {
	Biography    string `json:"biography"`
	Gender       string `json:"gender"`
	Nickname     string `json:"nickname"`
	ProfileImage string `json:"profile_image"`
	Type         string `json:"type"`
	UID          string `json:"uid"`
}
