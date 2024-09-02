package yaygo

type User struct {
	Biography string `json:"biography"`
	ID        int64  `json:"id"`
	Nickname  string `json:"nickname"`
}

type UserResponse struct {
	User User `json:"user"`
}
