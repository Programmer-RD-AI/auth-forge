package session


type SessionService struct {
	sessionRepository *SessionRepository
}

func NewSessionService(repo *SessionRepository) SessionService{
	return SessionService{
		sessionRepository: repo,
	}
}


func (s *SessionService) CreateSession(){
	
}
