package yaygo

type PostAPI struct {
	s *Session
}

func newPostAPI(s *Session) *PostAPI {
	return &PostAPI{
		s: s,
	}
}
