package yaygo

type ThreadApi struct {
	s *Session
}

func newThreadApi(s *Session) *ThreadApi {
	return &ThreadApi{
		s: s,
	}
}
