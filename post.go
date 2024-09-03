package yaygo

type PostApi struct {
	s *Session
}

func newPostApi(s *Session) *PostApi {
	return &PostApi{
		s: s,
	}
}
