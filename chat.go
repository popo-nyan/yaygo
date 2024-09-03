package yaygo

type ChatApi struct {
	s *Session
}

func newChatApi(s *Session) *ChatApi {
	return &ChatApi{
		s: s,
	}
}
