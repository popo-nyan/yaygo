package yaygo

type Response struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Response
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
	BanUntil  int    `json:"ban_until"`
	RetryIn   int    `json:"retry_in"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}
