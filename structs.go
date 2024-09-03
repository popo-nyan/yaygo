package yaygo

type User struct {
	Biography string `json:"biography"`
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
}

type UserResponse struct {
	User User `json:"user"`
}
