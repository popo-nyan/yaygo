package yaygo

type GroupApi struct {
	s *Session
}

func newGroupApi(s *Session) *GroupApi {
	return &GroupApi{
		s: s,
	}
}
