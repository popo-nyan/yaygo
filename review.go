package yaygo

type ReviewApi struct {
	s *Session
}

func newReviewApi(s *Session) *ReviewApi {
	return &ReviewApi{
		s: s,
	}
}
