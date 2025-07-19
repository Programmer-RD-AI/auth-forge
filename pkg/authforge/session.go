package authforge

func CreateOrExpireToken(userId string) string {
	return *new(string)
}

func SessionValidation(sessionId string) bool { return true }

func RevokeSession(sessionId string) (bool, error) { return true, nil }
