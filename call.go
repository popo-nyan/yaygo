package yaygo

type CallApi struct {
	s *Session
}

func newCallApi(s *Session) *CallApi {
	return &CallApi{
		s: s,
	}
}
