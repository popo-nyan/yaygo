package yaygo

type MiscApi struct {
	s *Session
}

func newMiscApi(s *Session) *MiscApi {
	return &MiscApi{
		s: s,
	}
}
