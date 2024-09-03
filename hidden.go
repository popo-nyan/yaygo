package yaygo

type HiddenApi struct {
	s *Session
}

func newHiddenApi(s *Session) *HiddenApi {
	return &HiddenApi{
		s: s,
	}
}
