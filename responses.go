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

type LoginUserResponse struct {
	Response
	AccessToken  string  `json:"access_token"`
	ExpiresIn    int     `json:"expires_in"`
	IsNew        bool    `json:"is_new"`
	RefreshToken string  `json:"refresh_token"`
	SNSInfo      SNSInfo `json:"sns_info"`
	UserID       int     `json:"user_id"`
	Username     string  `json:"username"`
}

type UserTimestampResponse struct {
	Response
	Country   string `json:"country"`
	IPAddress string `json:"ip_address"`
	Time      int    `json:"time"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}
