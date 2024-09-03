package yaygo

type AuthApi struct {
	s *Session
}

func newAuthApi(s *Session) *AuthApi {
	return &AuthApi{
		s: s,
	}
}
