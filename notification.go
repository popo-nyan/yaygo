package yaygo

type NotificationAPI struct {
	s *Session
}

func newNotificationAPI(s *Session) *NotificationAPI {
	return &NotificationAPI{
		s: s,
	}
}
