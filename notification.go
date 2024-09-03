package yaygo

type NotificationApi struct {
	s *Session
}

func newNotificationApi(s *Session) *NotificationApi {
	return &NotificationApi{
		s: s,
	}
}
